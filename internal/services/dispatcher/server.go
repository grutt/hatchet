package dispatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/steebchen/prisma-client-go/runtime/types"

	"github.com/hatchet-dev/hatchet/internal/datautils"
	"github.com/hatchet-dev/hatchet/internal/repository"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/internal/services/dispatcher/contracts"
	"github.com/hatchet-dev/hatchet/internal/services/shared/tasktypes"
	"github.com/hatchet-dev/hatchet/internal/taskqueue"
	"github.com/hatchet-dev/hatchet/internal/telemetry"
)

func (d *DispatcherImpl) GetWorker(workerId string) (*subscribedWorker, error) {
	workerInt, ok := d.workers.Load(workerId)
	if !ok {
		return nil, fmt.Errorf("worker with id %s not found", workerId)
	}

	worker, ok := workerInt.(subscribedWorker)

	if !ok {
		return nil, fmt.Errorf("failed to cast worker with id %s to subscribedWorker", workerId)
	}

	return &worker, nil
}

type subscribedWorker struct {
	// stream is the server side of the RPC stream
	stream contracts.Dispatcher_ListenServer

	// finished is used to signal closure of a client subscribing goroutine
	finished chan<- bool
}

func (worker *subscribedWorker) StartStepRun(
	ctx context.Context,
	tenantId string,
	stepRun *db.StepRunModel,
) error {
	ctx, span := telemetry.NewSpan(ctx, "start-step-run")
	defer span.End()

	inputBytes := []byte{}
	inputType, ok := stepRun.Input()

	if ok {
		var err error
		inputBytes, err = inputType.MarshalJSON()

		if err != nil {
			return err
		}
	}

	return worker.stream.Send(&contracts.AssignedAction{
		TenantId:      tenantId,
		JobId:         stepRun.Step().JobID,
		JobName:       stepRun.Step().Job().Name,
		JobRunId:      stepRun.JobRunID,
		StepId:        stepRun.StepID,
		StepRunId:     stepRun.ID,
		ActionType:    contracts.ActionType_START_STEP_RUN,
		ActionId:      stepRun.Step().ActionID,
		ActionPayload: string(inputBytes),
	})
}

func (worker *subscribedWorker) StartGroupKeyAction(
	ctx context.Context,
	tenantId string,
	workflowRun *db.WorkflowRunModel,
) error {
	ctx, span := telemetry.NewSpan(ctx, "start-group-key-action")
	defer span.End()

	inputBytes := []byte{}

	concurrency, ok := workflowRun.WorkflowVersion().Concurrency()

	if !ok {
		return fmt.Errorf("could not get concurrency for workflow version %s", workflowRun.WorkflowVersionID)
	}

	concurrencyFn, ok := concurrency.GetConcurrencyGroup()

	if !ok {
		return fmt.Errorf("could not get concurrency group for workflow version %s", workflowRun.WorkflowVersionID)
	}

	// get the input from the workflow run
	triggeredBy, ok := workflowRun.TriggeredBy()

	if !ok {
		return fmt.Errorf("could not get triggered by from workflow run %s", workflowRun.ID)
	}

	var inputData types.JSON

	if event, ok := triggeredBy.Event(); ok {
		inputData, _ = event.Data()
	} else if schedule, ok := triggeredBy.Scheduled(); ok {
		inputData, _ = schedule.Input()
	} else if cron, ok := triggeredBy.Cron(); ok {
		inputData, _ = cron.Input()
	}

	inputBytes, err := inputData.MarshalJSON()

	if err != nil {
		return err
	}

	getGroupKeyRun, ok := workflowRun.GetGroupKeyRun()

	if !ok {
		return fmt.Errorf("could not get get group key run for workflow run %s", workflowRun.ID)
	}

	return worker.stream.Send(&contracts.AssignedAction{
		TenantId:         tenantId,
		WorkflowRunId:    workflowRun.ID,
		GetGroupKeyRunId: getGroupKeyRun.ID,
		ActionType:       contracts.ActionType_START_GET_GROUP_KEY,
		ActionId:         concurrencyFn.ActionID,
		ActionPayload:    string(inputBytes),
	})
}

func (worker *subscribedWorker) CancelStepRun(
	ctx context.Context,
	tenantId string,
	stepRun *db.StepRunModel,
) error {
	ctx, span := telemetry.NewSpan(ctx, "cancel-step-run")
	defer span.End()

	return worker.stream.Send(&contracts.AssignedAction{
		TenantId:   tenantId,
		JobId:      stepRun.Step().JobID,
		JobName:    stepRun.Step().Job().Name,
		JobRunId:   stepRun.JobRunID,
		StepId:     stepRun.StepID,
		StepRunId:  stepRun.ID,
		ActionType: contracts.ActionType_CANCEL_STEP_RUN,
	})
}

func (s *DispatcherImpl) Register(ctx context.Context, request *contracts.WorkerRegisterRequest) (*contracts.WorkerRegisterResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received register request from ID %s with actions %v", request.WorkerName, request.Actions)

	svcs := request.Services

	if len(svcs) == 0 {
		svcs = []string{"default"}
	}

	// create a worker in the database
	worker, err := s.repo.Worker().CreateNewWorker(tenant.ID, &repository.CreateWorkerOpts{
		DispatcherId: s.dispatcherId,
		Name:         request.WorkerName,
		Actions:      request.Actions,
		Services:     svcs,
	})

	if err != nil {
		s.l.Error().Err(err).Msgf("could not create worker for tenant %s", tenant.ID)
		return nil, err
	}

	s.l.Debug().Msgf("Registered worker with ID: %s", worker.ID)

	// return the worker id to the worker
	return &contracts.WorkerRegisterResponse{
		TenantId:   worker.TenantID,
		WorkerId:   worker.ID,
		WorkerName: worker.Name,
	}, nil
}

// Subscribe handles a subscribe request from a client
func (s *DispatcherImpl) Listen(request *contracts.WorkerListenRequest, stream contracts.Dispatcher_ListenServer) error {
	tenant := stream.Context().Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received subscribe request from ID: %s", request.WorkerId)

	worker, err := s.repo.Worker().GetWorkerById(request.WorkerId)

	if err != nil {
		s.l.Error().Err(err).Msgf("could not get worker %s", request.WorkerId)
		return err
	}

	// check the worker's dispatcher against the current dispatcher. if they don't match, then update the worker
	if worker.DispatcherID != s.dispatcherId {
		_, err = s.repo.Worker().UpdateWorker(tenant.ID, request.WorkerId, &repository.UpdateWorkerOpts{
			DispatcherId: &s.dispatcherId,
		})

		if err != nil {
			s.l.Error().Err(err).Msgf("could not update worker %s dispatcher", request.WorkerId)
			return err
		}
	}

	fin := make(chan bool)

	s.workers.Store(request.WorkerId, subscribedWorker{stream: stream, finished: fin})

	defer func() {
		s.workers.Delete(request.WorkerId)

		inactive := db.WorkerStatusInactive

		_, err := s.repo.Worker().UpdateWorker(tenant.ID, request.WorkerId, &repository.UpdateWorkerOpts{
			Status: &inactive,
		})

		if err != nil {
			s.l.Error().Err(err).Msgf("could not update worker %s status to inactive", request.WorkerId)
		}
	}()

	ctx := stream.Context()

	// update the worker with a last heartbeat time every 5 seconds as long as the worker is connected
	go func() {
		timer := time.NewTicker(100 * time.Millisecond)

		// set the last heartbeat to 6 seconds ago so the first heartbeat is sent immediately
		lastHeartbeat := time.Now().UTC().Add(-6 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				s.l.Debug().Msgf("worker id %s has disconnected", request.WorkerId)
				return
			case <-timer.C:
				if now := time.Now().UTC(); lastHeartbeat.Add(5 * time.Second).Before(now) {
					s.l.Debug().Msgf("updating worker %s heartbeat", request.WorkerId)

					_, err := s.repo.Worker().UpdateWorker(tenant.ID, request.WorkerId, &repository.UpdateWorkerOpts{
						LastHeartbeatAt: &now,
					})

					if err != nil {
						s.l.Error().Err(err).Msgf("could not update worker %s heartbeat", request.WorkerId)
						return
					}

					lastHeartbeat = time.Now().UTC()
				}
			}
		}
	}()

	// Keep the connection alive for sending messages
	for {
		select {
		case <-fin:
			s.l.Debug().Msgf("closing stream for worker id: %s", request.WorkerId)
			return nil
		case <-ctx.Done():
			s.l.Debug().Msgf("worker id %s has disconnected", request.WorkerId)
			return nil
		}
	}
}

// SubscribeToWorkflowEvents registers workflow events with the dispatcher
func (s *DispatcherImpl) SubscribeToWorkflowEvents(request *contracts.SubscribeToWorkflowEventsRequest, stream contracts.Dispatcher_SubscribeToWorkflowEventsServer) error {
	tenant := stream.Context().Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received subscribe request for workflow: %s", request.WorkflowRunId)

	q, err := taskqueue.TenantEventConsumerQueue(tenant.ID)

	if err != nil {
		return err
	}

	ctx := stream.Context()

	// subscribe to the task queue for the tenant
	taskChan, err := s.tq.Subscribe(ctx, q)

	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			// drain the existing connections
			return nil
		case task := <-taskChan:
			go func(task *taskqueue.Task) {
				e, err := s.tenantTaskToWorkflowEvent(task, tenant.ID, request.WorkflowRunId)

				if err != nil {
					s.l.Error().Err(err).Msgf("could not convert task to workflow event")
					return
				} else if e == nil {
					return
				}

				fmt.Println("<><><> SENDING TASK!!", e)

				// send the task to the client
				err = stream.Send(e)

				if err != nil {
					s.l.Error().Err(err).Msgf("could not send workflow event to client")
				}
			}(task)
		}
	}
}

func (s *DispatcherImpl) SendStepActionEvent(ctx context.Context, request *contracts.StepActionEvent) (*contracts.ActionEventResponse, error) {
	switch request.EventType {
	case contracts.StepActionEventType_STEP_EVENT_TYPE_STARTED:
		return s.handleStepRunStarted(ctx, request)
	case contracts.StepActionEventType_STEP_EVENT_TYPE_COMPLETED:
		return s.handleStepRunCompleted(ctx, request)
	case contracts.StepActionEventType_STEP_EVENT_TYPE_FAILED:
		return s.handleStepRunFailed(ctx, request)
	}

	return nil, fmt.Errorf("unknown event type %s", request.EventType)
}

func (s *DispatcherImpl) SendGroupKeyActionEvent(ctx context.Context, request *contracts.GroupKeyActionEvent) (*contracts.ActionEventResponse, error) {
	switch request.EventType {
	case contracts.GroupKeyActionEventType_GROUP_KEY_EVENT_TYPE_STARTED:
		return s.handleGetGroupKeyRunStarted(ctx, request)
	case contracts.GroupKeyActionEventType_GROUP_KEY_EVENT_TYPE_COMPLETED:
		return s.handleGetGroupKeyRunCompleted(ctx, request)
	case contracts.GroupKeyActionEventType_GROUP_KEY_EVENT_TYPE_FAILED:
		return s.handleGetGroupKeyRunFailed(ctx, request)
	}

	return nil, fmt.Errorf("unknown event type %s", request.EventType)
}

func (s *DispatcherImpl) Unsubscribe(ctx context.Context, request *contracts.WorkerUnsubscribeRequest) (*contracts.WorkerUnsubscribeResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	// no matter what, remove the worker from the connection pool
	defer s.workers.Delete(request.WorkerId)

	err := s.repo.Worker().DeleteWorker(tenant.ID, request.WorkerId)

	if err != nil {
		return nil, err
	}

	return &contracts.WorkerUnsubscribeResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleStepRunStarted(ctx context.Context, request *contracts.StepActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step started event for step run %s", request.StepRunId)

	startedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.StepRunStartedTaskPayload{
		StepRunId: request.StepRunId,
		StartedAt: startedAt.Format(time.RFC3339),
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.StepRunStartedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.JOB_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "step-run-started",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleStepRunCompleted(ctx context.Context, request *contracts.StepActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step completed event for step run %s", request.StepRunId)

	finishedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.StepRunFinishedTaskPayload{
		StepRunId:      request.StepRunId,
		FinishedAt:     finishedAt.Format(time.RFC3339),
		StepOutputData: request.EventPayload,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.StepRunFinishedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.JOB_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "step-run-finished",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleStepRunFailed(ctx context.Context, request *contracts.StepActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step failed event for step run %s", request.StepRunId)

	failedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.StepRunFailedTaskPayload{
		StepRunId: request.StepRunId,
		FailedAt:  failedAt.Format(time.RFC3339),
		Error:     request.EventPayload,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.StepRunFailedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.JOB_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "step-run-failed",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleGetGroupKeyRunStarted(ctx context.Context, request *contracts.GroupKeyActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step started event for step run %s", request.GetGroupKeyRunId)

	startedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunStartedTaskPayload{
		GetGroupKeyRunId: request.GetGroupKeyRunId,
		StartedAt:        startedAt.Format(time.RFC3339),
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunStartedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.WORKFLOW_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "get-group-key-run-started",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleGetGroupKeyRunCompleted(ctx context.Context, request *contracts.GroupKeyActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step completed event for step run %s", request.GetGroupKeyRunId)

	finishedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunFinishedTaskPayload{
		GetGroupKeyRunId: request.GetGroupKeyRunId,
		FinishedAt:       finishedAt.Format(time.RFC3339),
		GroupKey:         request.EventPayload,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunFinishedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.WORKFLOW_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "get-group-key-run-finished",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) handleGetGroupKeyRunFailed(ctx context.Context, request *contracts.GroupKeyActionEvent) (*contracts.ActionEventResponse, error) {
	tenant := ctx.Value("tenant").(*db.TenantModel)

	s.l.Debug().Msgf("Received step failed event for step run %s", request.GetGroupKeyRunId)

	failedAt := request.EventTimestamp.AsTime()

	payload, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunFailedTaskPayload{
		GetGroupKeyRunId: request.GetGroupKeyRunId,
		FailedAt:         failedAt.Format(time.RFC3339),
		Error:            request.EventPayload,
	})

	metadata, _ := datautils.ToJSONMap(tasktypes.GetGroupKeyRunFailedTaskMetadata{
		TenantId: tenant.ID,
	})

	// send the event to the jobs queue
	err := s.tq.AddTask(ctx, taskqueue.WORKFLOW_PROCESSING_QUEUE, &taskqueue.Task{
		ID:       "get-group-key-run-failed",
		Payload:  payload,
		Metadata: metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.ActionEventResponse{
		TenantId: tenant.ID,
		WorkerId: request.WorkerId,
	}, nil
}

func (s *DispatcherImpl) tenantTaskToWorkflowEvent(task *taskqueue.Task, tenantId, workflowRunId string) (*contracts.WorkflowEvent, error) {
	// TODO: eventually process workflows as well, this is just steps
	workflowEvent := &contracts.WorkflowEvent{
		ResourceType: contracts.ResourceType_RESOURCE_TYPE_STEP_RUN,
	}

	var stepRunId string

	switch task.ID {
	case "step-run-started":
		stepRunId = task.Payload["step_run_id"].(string)
		workflowEvent.ResourceId = stepRunId
		workflowEvent.EventType = contracts.ResourceEventType_RESOURCE_EVENT_TYPE_STARTED
	case "step-run-finished":
		stepRunId = task.Payload["step_run_id"].(string)
		workflowEvent.ResourceId = stepRunId
		workflowEvent.EventType = contracts.ResourceEventType_RESOURCE_EVENT_TYPE_COMPLETED
		workflowEvent.EventPayload = task.Payload["step_output_data"].(string)
	case "step-run-failed":
		stepRunId = task.Payload["step_run_id"].(string)
		workflowEvent.ResourceId = stepRunId
		workflowEvent.EventType = contracts.ResourceEventType_RESOURCE_EVENT_TYPE_FAILED
		workflowEvent.EventPayload = task.Payload["error"].(string)
	case "step-run-cancelled":
		stepRunId = task.Payload["step_run_id"].(string)
		workflowEvent.ResourceId = stepRunId
		workflowEvent.EventType = contracts.ResourceEventType_RESOURCE_EVENT_TYPE_CANCELLED
	case "step-run-timed-out":
		stepRunId = task.Payload["step_run_id"].(string)
		workflowEvent.ResourceId = stepRunId
		workflowEvent.EventType = contracts.ResourceEventType_RESOURCE_EVENT_TYPE_TIMED_OUT
	}

	if stepRunId == "" {
		// expected because not all tasks have step run ids
		return nil, nil
	}

	// determine if this step run matches the workflow run id
	stepRun, err := s.repo.StepRun().GetStepRunById(tenantId, stepRunId)

	if err != nil {
		return nil, err
	}

	if stepRun.JobRun().WorkflowRunID != workflowRunId {
		// this is an expected error, so we don't return it
		return nil, nil
	}

	return workflowEvent, nil
}
