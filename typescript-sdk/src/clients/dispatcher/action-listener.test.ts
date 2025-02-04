import { ActionType, AssignedAction } from '@hatchet/protoc/dispatcher';
import sleep from '@util/sleep';
import { ServerError, Status } from 'nice-grpc-common';
import { DispatcherClient } from './dispatcher-client';
import { ActionListener } from './action-listener';
import { mockChannel, mockFactory } from '../hatchet-client/hatchet-client.test';

let dispatcher: DispatcherClient;

type AssignActionMock = AssignedAction | Error;

// Mock data for AssignedAction
const mockAssignedActions: AssignActionMock[] = [
  {
    tenantId: 'tenant1',
    jobId: 'job1',
    jobName: 'Job One',
    jobRunId: 'run1',
    stepId: 'step1',
    stepRunId: 'runStep1',
    actionId: 'action1',
    actionType: ActionType.START_STEP_RUN,
    actionPayload: 'payload1',
    workflowRunId: 'workflowRun1',
    getGroupKeyRunId: 'groupKeyRun1',
  },
  // ... Add more mock AssignedAction objects as needed
];

// Mock implementation of the listener
export const mockListener = (fixture: AssignActionMock[]) =>
  (async function* gen() {
    for (const action of fixture) {
      // Simulate asynchronous behavior
      await sleep(100);

      if (action instanceof Error) {
        throw action;
      }

      yield action;
    }
  })();

describe('ActionListener', () => {
  beforeEach(() => {
    dispatcher = new DispatcherClient(
      {
        token: 'TOKEN',

        host_port: 'HOST_PORT',
        log_level: 'OFF',
        tls_config: {
          cert_file: 'TLS_CERT_FILE',
          key_file: 'TLS_KEY_FILE',
          ca_file: 'TLS_ROOT_CA_FILE',
          server_name: 'TLS_SERVER_NAME',
        },
      },
      mockChannel,
      mockFactory
    );
  });

  it('should create a client', async () => {
    const listener = new ActionListener(dispatcher, mockListener(mockAssignedActions), 'WORKER_ID');
    expect(listener).toBeDefined();
    expect(listener.workerId).toEqual('WORKER_ID');
  });

  describe('actions', () => {
    it('it should "yield" actions', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener([...mockAssignedActions, new ServerError(Status.CANCELLED, 'CANCELLED')]),
        'WORKER_ID'
      );

      const actions = listener.actions();
      const res = [];
      for await (const action of actions) {
        res.push(action);
      }

      expect(res[0]).toEqual({
        tenantId: 'tenant1',
        jobId: 'job1',
        jobName: 'Job One',
        jobRunId: 'run1',
        stepId: 'step1',
        stepRunId: 'runStep1',
        actionId: 'action1',
        actionType: ActionType.START_STEP_RUN,
        actionPayload: 'payload1',
        workflowRunId: 'workflowRun1',
        getGroupKeyRunId: 'groupKeyRun1',
      });
    });

    it('it should break on grpc CANCELLED', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener([...mockAssignedActions, new ServerError(Status.CANCELLED, 'CANCELLED')]),
        'WORKER_ID'
      );

      const actions = listener.actions();
      const retrySpy = jest.spyOn(listener, 'retrySubscribe').mockResolvedValue(undefined);

      const res = [];
      for await (const action of actions) {
        res.push(action);
      }

      expect(res.length).toEqual(1);
      expect(retrySpy).not.toHaveBeenCalled();
    });

    it('it should break on unknown error', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener([...mockAssignedActions, new Error('Simulated error')]),
        'WORKER_ID'
      );

      const actions = listener.actions();
      const retrySpy = jest.spyOn(listener, 'retrySubscribe').mockResolvedValue(undefined);

      const res = [];
      for await (const action of actions) {
        res.push(action);
      }

      expect(res.length).toEqual(1);
      expect(retrySpy).not.toHaveBeenCalled();
    });

    it('it should attempt to re-establish connection on grpc UNAVAILABLE', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener([...mockAssignedActions, new ServerError(Status.UNAVAILABLE, 'UNAVAILABLE')]),
        'WORKER_ID'
      );

      const retrySpy = jest.spyOn(listener, 'retrySubscribe').mockResolvedValue(undefined);

      const actions = listener.actions();

      const res = [];
      for await (const action of actions) {
        res.push(action);
      }

      expect(res.length).toEqual(1);
      expect(retrySpy).toHaveBeenCalled();
    });
  });

  describe('retry_subscribe', () => {
    it('should exit after successful connection', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener(mockAssignedActions),
        'WORKER_ID'
      );

      // Mock the listener to throw an error on the first call
      const listenSpy = jest
        .spyOn(listener.client, 'listen')
        .mockReturnValue(mockListener(mockAssignedActions));

      await listener.retrySubscribe();

      expect(listenSpy).toHaveBeenCalledTimes(1);
    });

    it('should retry until success', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener(mockAssignedActions),
        'WORKER_ID'
      );

      // Mock the listener to throw an error on the first call
      const listenSpy = jest
        .spyOn(listener.client, 'listen')
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => mockListener(mockAssignedActions));

      await expect(async () => {
        await listener.retrySubscribe();
      }).not.toThrow();
    });

    it('should not throw an error if successful', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener(mockAssignedActions),
        'WORKER_ID'
      );

      // Mock the listener to throw an error on the first call
      const listenSpy = jest
        .spyOn(listener.client, 'listen')
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => mockListener(mockAssignedActions));

      await listener.retrySubscribe();
      expect(listenSpy).toHaveBeenCalledTimes(2);
    });

    it('should retry at most COUNT times and throw an error', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener(mockAssignedActions),
        'WORKER_ID'
      );

      // Mock the listener to throw an error on the first call
      const listenSpy = jest
        .spyOn(listener.client, 'listen')
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => {
          throw new Error('Simulated error');
        })
        .mockImplementationOnce(() => mockListener(mockAssignedActions));

      try {
        await listener.retrySubscribe();
        expect(listenSpy).toHaveBeenCalledTimes(5);
      } catch (e: any) {
        expect(e.message).toEqual(`Could not subscribe to the worker after 5 retries`);
      }
    });
  });

  describe('unregister', () => {
    it('should unsubscribe itself', async () => {
      const listener = new ActionListener(
        dispatcher,
        mockListener(mockAssignedActions),
        'WORKER_ID'
      );

      const unsubscribeSpy = jest.spyOn(listener.client, 'unsubscribe').mockResolvedValue({
        tenantId: 'TENANT_ID',
        workerId: 'WORKER_ID',
      });

      const res = await listener.unregister();

      expect(unsubscribeSpy).toHaveBeenCalled();
      expect(res.workerId).toEqual('WORKER_ID');
    });
  });
});
