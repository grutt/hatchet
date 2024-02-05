# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dispatcher.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x64ispatcher.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"N\n\x15WorkerRegisterRequest\x12\x12\n\nworkerName\x18\x01 \x01(\t\x12\x0f\n\x07\x61\x63tions\x18\x02 \x03(\t\x12\x10\n\x08services\x18\x03 \x03(\t\"P\n\x16WorkerRegisterResponse\x12\x10\n\x08tenantId\x18\x01 \x01(\t\x12\x10\n\x08workerId\x18\x02 \x01(\t\x12\x12\n\nworkerName\x18\x03 \x01(\t\"\xf2\x01\n\x0e\x41ssignedAction\x12\x10\n\x08tenantId\x18\x01 \x01(\t\x12\x15\n\rworkflowRunId\x18\x02 \x01(\t\x12\x18\n\x10getGroupKeyRunId\x18\x03 \x01(\t\x12\r\n\x05jobId\x18\x04 \x01(\t\x12\x0f\n\x07jobName\x18\x05 \x01(\t\x12\x10\n\x08jobRunId\x18\x06 \x01(\t\x12\x0e\n\x06stepId\x18\x07 \x01(\t\x12\x11\n\tstepRunId\x18\x08 \x01(\t\x12\x10\n\x08\x61\x63tionId\x18\t \x01(\t\x12\x1f\n\nactionType\x18\n \x01(\x0e\x32\x0b.ActionType\x12\x15\n\ractionPayload\x18\x0b \x01(\t\"\'\n\x13WorkerListenRequest\x12\x10\n\x08workerId\x18\x01 \x01(\t\",\n\x18WorkerUnsubscribeRequest\x12\x10\n\x08workerId\x18\x01 \x01(\t\"?\n\x19WorkerUnsubscribeResponse\x12\x10\n\x08tenantId\x18\x01 \x01(\t\x12\x10\n\x08workerId\x18\x02 \x01(\t\"\xe1\x01\n\x13GroupKeyActionEvent\x12\x10\n\x08workerId\x18\x01 \x01(\t\x12\x15\n\rworkflowRunId\x18\x02 \x01(\t\x12\x18\n\x10getGroupKeyRunId\x18\x03 \x01(\t\x12\x10\n\x08\x61\x63tionId\x18\x04 \x01(\t\x12\x32\n\x0e\x65ventTimestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12+\n\teventType\x18\x06 \x01(\x0e\x32\x18.GroupKeyActionEventType\x12\x14\n\x0c\x65ventPayload\x18\x07 \x01(\t\"\xec\x01\n\x0fStepActionEvent\x12\x10\n\x08workerId\x18\x01 \x01(\t\x12\r\n\x05jobId\x18\x02 \x01(\t\x12\x10\n\x08jobRunId\x18\x03 \x01(\t\x12\x0e\n\x06stepId\x18\x04 \x01(\t\x12\x11\n\tstepRunId\x18\x05 \x01(\t\x12\x10\n\x08\x61\x63tionId\x18\x06 \x01(\t\x12\x32\n\x0e\x65ventTimestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\teventType\x18\x08 \x01(\x0e\x32\x14.StepActionEventType\x12\x14\n\x0c\x65ventPayload\x18\t \x01(\t\"9\n\x13\x41\x63tionEventResponse\x12\x10\n\x08tenantId\x18\x01 \x01(\t\x12\x10\n\x08workerId\x18\x02 \x01(\t\"9\n SubscribeToWorkflowEventsRequest\x12\x15\n\rworkflowRunId\x18\x01 \x01(\t\"\xd0\x01\n\rWorkflowEvent\x12\x15\n\rworkflowRunId\x18\x01 \x01(\t\x12#\n\x0cresourceType\x18\x02 \x01(\x0e\x32\r.ResourceType\x12%\n\teventType\x18\x03 \x01(\x0e\x32\x12.ResourceEventType\x12\x12\n\nresourceId\x18\x04 \x01(\t\x12\x32\n\x0e\x65ventTimestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x14\n\x0c\x65ventPayload\x18\x06 \x01(\t*N\n\nActionType\x12\x12\n\x0eSTART_STEP_RUN\x10\x00\x12\x13\n\x0f\x43\x41NCEL_STEP_RUN\x10\x01\x12\x17\n\x13START_GET_GROUP_KEY\x10\x02*\xa2\x01\n\x17GroupKeyActionEventType\x12 \n\x1cGROUP_KEY_EVENT_TYPE_UNKNOWN\x10\x00\x12 \n\x1cGROUP_KEY_EVENT_TYPE_STARTED\x10\x01\x12\"\n\x1eGROUP_KEY_EVENT_TYPE_COMPLETED\x10\x02\x12\x1f\n\x1bGROUP_KEY_EVENT_TYPE_FAILED\x10\x03*\x8a\x01\n\x13StepActionEventType\x12\x1b\n\x17STEP_EVENT_TYPE_UNKNOWN\x10\x00\x12\x1b\n\x17STEP_EVENT_TYPE_STARTED\x10\x01\x12\x1d\n\x19STEP_EVENT_TYPE_COMPLETED\x10\x02\x12\x1a\n\x16STEP_EVENT_TYPE_FAILED\x10\x03*e\n\x0cResourceType\x12\x19\n\x15RESOURCE_TYPE_UNKNOWN\x10\x00\x12\x1a\n\x16RESOURCE_TYPE_STEP_RUN\x10\x01\x12\x1e\n\x1aRESOURCE_TYPE_WORKFLOW_RUN\x10\x02*\xde\x01\n\x11ResourceEventType\x12\x1f\n\x1bRESOURCE_EVENT_TYPE_UNKNOWN\x10\x00\x12\x1f\n\x1bRESOURCE_EVENT_TYPE_STARTED\x10\x01\x12!\n\x1dRESOURCE_EVENT_TYPE_COMPLETED\x10\x02\x12\x1e\n\x1aRESOURCE_EVENT_TYPE_FAILED\x10\x03\x12!\n\x1dRESOURCE_EVENT_TYPE_CANCELLED\x10\x04\x12!\n\x1dRESOURCE_EVENT_TYPE_TIMED_OUT\x10\x05\x32\xa6\x03\n\nDispatcher\x12=\n\x08Register\x12\x16.WorkerRegisterRequest\x1a\x17.WorkerRegisterResponse\"\x00\x12\x33\n\x06Listen\x12\x14.WorkerListenRequest\x1a\x0f.AssignedAction\"\x00\x30\x01\x12R\n\x19SubscribeToWorkflowEvents\x12!.SubscribeToWorkflowEventsRequest\x1a\x0e.WorkflowEvent\"\x00\x30\x01\x12?\n\x13SendStepActionEvent\x12\x10.StepActionEvent\x1a\x14.ActionEventResponse\"\x00\x12G\n\x17SendGroupKeyActionEvent\x12\x14.GroupKeyActionEvent\x1a\x14.ActionEventResponse\"\x00\x12\x46\n\x0bUnsubscribe\x12\x19.WorkerUnsubscribeRequest\x1a\x1a.WorkerUnsubscribeResponse\"\x00\x42GZEgithub.com/hatchet-dev/hatchet/internal/services/dispatcher/contractsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'dispatcher_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZEgithub.com/hatchet-dev/hatchet/internal/services/dispatcher/contracts'
  _globals['_ACTIONTYPE']._serialized_start=1408
  _globals['_ACTIONTYPE']._serialized_end=1486
  _globals['_GROUPKEYACTIONEVENTTYPE']._serialized_start=1489
  _globals['_GROUPKEYACTIONEVENTTYPE']._serialized_end=1651
  _globals['_STEPACTIONEVENTTYPE']._serialized_start=1654
  _globals['_STEPACTIONEVENTTYPE']._serialized_end=1792
  _globals['_RESOURCETYPE']._serialized_start=1794
  _globals['_RESOURCETYPE']._serialized_end=1895
  _globals['_RESOURCEEVENTTYPE']._serialized_start=1898
  _globals['_RESOURCEEVENTTYPE']._serialized_end=2120
  _globals['_WORKERREGISTERREQUEST']._serialized_start=53
  _globals['_WORKERREGISTERREQUEST']._serialized_end=131
  _globals['_WORKERREGISTERRESPONSE']._serialized_start=133
  _globals['_WORKERREGISTERRESPONSE']._serialized_end=213
  _globals['_ASSIGNEDACTION']._serialized_start=216
  _globals['_ASSIGNEDACTION']._serialized_end=458
  _globals['_WORKERLISTENREQUEST']._serialized_start=460
  _globals['_WORKERLISTENREQUEST']._serialized_end=499
  _globals['_WORKERUNSUBSCRIBEREQUEST']._serialized_start=501
  _globals['_WORKERUNSUBSCRIBEREQUEST']._serialized_end=545
  _globals['_WORKERUNSUBSCRIBERESPONSE']._serialized_start=547
  _globals['_WORKERUNSUBSCRIBERESPONSE']._serialized_end=610
  _globals['_GROUPKEYACTIONEVENT']._serialized_start=613
  _globals['_GROUPKEYACTIONEVENT']._serialized_end=838
  _globals['_STEPACTIONEVENT']._serialized_start=841
  _globals['_STEPACTIONEVENT']._serialized_end=1077
  _globals['_ACTIONEVENTRESPONSE']._serialized_start=1079
  _globals['_ACTIONEVENTRESPONSE']._serialized_end=1136
  _globals['_SUBSCRIBETOWORKFLOWEVENTSREQUEST']._serialized_start=1138
  _globals['_SUBSCRIBETOWORKFLOWEVENTSREQUEST']._serialized_end=1195
  _globals['_WORKFLOWEVENT']._serialized_start=1198
  _globals['_WORKFLOWEVENT']._serialized_end=1406
  _globals['_DISPATCHER']._serialized_start=2123
  _globals['_DISPATCHER']._serialized_end=2545
# @@protoc_insertion_point(module_scope)
