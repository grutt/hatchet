// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbsqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ConcurrencyLimitStrategy string

const (
	ConcurrencyLimitStrategyCANCELINPROGRESS ConcurrencyLimitStrategy = "CANCEL_IN_PROGRESS"
	ConcurrencyLimitStrategyDROPNEWEST       ConcurrencyLimitStrategy = "DROP_NEWEST"
	ConcurrencyLimitStrategyQUEUENEWEST      ConcurrencyLimitStrategy = "QUEUE_NEWEST"
	ConcurrencyLimitStrategyGROUPROUNDROBIN  ConcurrencyLimitStrategy = "GROUP_ROUND_ROBIN"
)

func (e *ConcurrencyLimitStrategy) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ConcurrencyLimitStrategy(s)
	case string:
		*e = ConcurrencyLimitStrategy(s)
	default:
		return fmt.Errorf("unsupported scan type for ConcurrencyLimitStrategy: %T", src)
	}
	return nil
}

type NullConcurrencyLimitStrategy struct {
	ConcurrencyLimitStrategy ConcurrencyLimitStrategy `json:"ConcurrencyLimitStrategy"`
	Valid                    bool                     `json:"valid"` // Valid is true if ConcurrencyLimitStrategy is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullConcurrencyLimitStrategy) Scan(value interface{}) error {
	if value == nil {
		ns.ConcurrencyLimitStrategy, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ConcurrencyLimitStrategy.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullConcurrencyLimitStrategy) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ConcurrencyLimitStrategy), nil
}

type InviteLinkStatus string

const (
	InviteLinkStatusPENDING  InviteLinkStatus = "PENDING"
	InviteLinkStatusACCEPTED InviteLinkStatus = "ACCEPTED"
	InviteLinkStatusREJECTED InviteLinkStatus = "REJECTED"
)

func (e *InviteLinkStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = InviteLinkStatus(s)
	case string:
		*e = InviteLinkStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for InviteLinkStatus: %T", src)
	}
	return nil
}

type NullInviteLinkStatus struct {
	InviteLinkStatus InviteLinkStatus `json:"InviteLinkStatus"`
	Valid            bool             `json:"valid"` // Valid is true if InviteLinkStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullInviteLinkStatus) Scan(value interface{}) error {
	if value == nil {
		ns.InviteLinkStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.InviteLinkStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullInviteLinkStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.InviteLinkStatus), nil
}

type JobKind string

const (
	JobKindDEFAULT   JobKind = "DEFAULT"
	JobKindONFAILURE JobKind = "ON_FAILURE"
)

func (e *JobKind) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobKind(s)
	case string:
		*e = JobKind(s)
	default:
		return fmt.Errorf("unsupported scan type for JobKind: %T", src)
	}
	return nil
}

type NullJobKind struct {
	JobKind JobKind `json:"JobKind"`
	Valid   bool    `json:"valid"` // Valid is true if JobKind is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobKind) Scan(value interface{}) error {
	if value == nil {
		ns.JobKind, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobKind.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobKind) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobKind), nil
}

type JobRunStatus string

const (
	JobRunStatusPENDING   JobRunStatus = "PENDING"
	JobRunStatusRUNNING   JobRunStatus = "RUNNING"
	JobRunStatusSUCCEEDED JobRunStatus = "SUCCEEDED"
	JobRunStatusFAILED    JobRunStatus = "FAILED"
	JobRunStatusCANCELLED JobRunStatus = "CANCELLED"
)

func (e *JobRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobRunStatus(s)
	case string:
		*e = JobRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for JobRunStatus: %T", src)
	}
	return nil
}

type NullJobRunStatus struct {
	JobRunStatus JobRunStatus `json:"JobRunStatus"`
	Valid        bool         `json:"valid"` // Valid is true if JobRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.JobRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobRunStatus), nil
}

type LimitResource string

const (
	LimitResourceWORKFLOWRUN LimitResource = "WORKFLOW_RUN"
	LimitResourceEVENT       LimitResource = "EVENT"
	LimitResourceWORKER      LimitResource = "WORKER"
	LimitResourceCRON        LimitResource = "CRON"
	LimitResourceSCHEDULE    LimitResource = "SCHEDULE"
)

func (e *LimitResource) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LimitResource(s)
	case string:
		*e = LimitResource(s)
	default:
		return fmt.Errorf("unsupported scan type for LimitResource: %T", src)
	}
	return nil
}

type NullLimitResource struct {
	LimitResource LimitResource `json:"LimitResource"`
	Valid         bool          `json:"valid"` // Valid is true if LimitResource is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLimitResource) Scan(value interface{}) error {
	if value == nil {
		ns.LimitResource, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LimitResource.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLimitResource) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.LimitResource), nil
}

type LogLineLevel string

const (
	LogLineLevelDEBUG LogLineLevel = "DEBUG"
	LogLineLevelINFO  LogLineLevel = "INFO"
	LogLineLevelWARN  LogLineLevel = "WARN"
	LogLineLevelERROR LogLineLevel = "ERROR"
)

func (e *LogLineLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LogLineLevel(s)
	case string:
		*e = LogLineLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for LogLineLevel: %T", src)
	}
	return nil
}

type NullLogLineLevel struct {
	LogLineLevel LogLineLevel `json:"LogLineLevel"`
	Valid        bool         `json:"valid"` // Valid is true if LogLineLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLogLineLevel) Scan(value interface{}) error {
	if value == nil {
		ns.LogLineLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LogLineLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLogLineLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.LogLineLevel), nil
}

type StepRunEventReason string

const (
	StepRunEventReasonREQUEUEDNOWORKER   StepRunEventReason = "REQUEUED_NO_WORKER"
	StepRunEventReasonREQUEUEDRATELIMIT  StepRunEventReason = "REQUEUED_RATE_LIMIT"
	StepRunEventReasonSCHEDULINGTIMEDOUT StepRunEventReason = "SCHEDULING_TIMED_OUT"
	StepRunEventReasonASSIGNED           StepRunEventReason = "ASSIGNED"
	StepRunEventReasonSTARTED            StepRunEventReason = "STARTED"
	StepRunEventReasonFINISHED           StepRunEventReason = "FINISHED"
	StepRunEventReasonFAILED             StepRunEventReason = "FAILED"
	StepRunEventReasonRETRYING           StepRunEventReason = "RETRYING"
	StepRunEventReasonCANCELLED          StepRunEventReason = "CANCELLED"
	StepRunEventReasonTIMEDOUT           StepRunEventReason = "TIMED_OUT"
	StepRunEventReasonREASSIGNED         StepRunEventReason = "REASSIGNED"
	StepRunEventReasonSLOTRELEASED       StepRunEventReason = "SLOT_RELEASED"
	StepRunEventReasonTIMEOUTREFRESHED   StepRunEventReason = "TIMEOUT_REFRESHED"
	StepRunEventReasonRETRIEDBYUSER      StepRunEventReason = "RETRIED_BY_USER"
)

func (e *StepRunEventReason) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StepRunEventReason(s)
	case string:
		*e = StepRunEventReason(s)
	default:
		return fmt.Errorf("unsupported scan type for StepRunEventReason: %T", src)
	}
	return nil
}

type NullStepRunEventReason struct {
	StepRunEventReason StepRunEventReason `json:"StepRunEventReason"`
	Valid              bool               `json:"valid"` // Valid is true if StepRunEventReason is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStepRunEventReason) Scan(value interface{}) error {
	if value == nil {
		ns.StepRunEventReason, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StepRunEventReason.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStepRunEventReason) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StepRunEventReason), nil
}

type StepRunEventSeverity string

const (
	StepRunEventSeverityINFO     StepRunEventSeverity = "INFO"
	StepRunEventSeverityWARNING  StepRunEventSeverity = "WARNING"
	StepRunEventSeverityCRITICAL StepRunEventSeverity = "CRITICAL"
)

func (e *StepRunEventSeverity) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StepRunEventSeverity(s)
	case string:
		*e = StepRunEventSeverity(s)
	default:
		return fmt.Errorf("unsupported scan type for StepRunEventSeverity: %T", src)
	}
	return nil
}

type NullStepRunEventSeverity struct {
	StepRunEventSeverity StepRunEventSeverity `json:"StepRunEventSeverity"`
	Valid                bool                 `json:"valid"` // Valid is true if StepRunEventSeverity is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStepRunEventSeverity) Scan(value interface{}) error {
	if value == nil {
		ns.StepRunEventSeverity, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StepRunEventSeverity.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStepRunEventSeverity) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StepRunEventSeverity), nil
}

type StepRunStatus string

const (
	StepRunStatusPENDING           StepRunStatus = "PENDING"
	StepRunStatusPENDINGASSIGNMENT StepRunStatus = "PENDING_ASSIGNMENT"
	StepRunStatusASSIGNED          StepRunStatus = "ASSIGNED"
	StepRunStatusRUNNING           StepRunStatus = "RUNNING"
	StepRunStatusSUCCEEDED         StepRunStatus = "SUCCEEDED"
	StepRunStatusFAILED            StepRunStatus = "FAILED"
	StepRunStatusCANCELLED         StepRunStatus = "CANCELLED"
)

func (e *StepRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StepRunStatus(s)
	case string:
		*e = StepRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for StepRunStatus: %T", src)
	}
	return nil
}

type NullStepRunStatus struct {
	StepRunStatus StepRunStatus `json:"StepRunStatus"`
	Valid         bool          `json:"valid"` // Valid is true if StepRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStepRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.StepRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StepRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStepRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StepRunStatus), nil
}

type TenantMemberRole string

const (
	TenantMemberRoleOWNER  TenantMemberRole = "OWNER"
	TenantMemberRoleADMIN  TenantMemberRole = "ADMIN"
	TenantMemberRoleMEMBER TenantMemberRole = "MEMBER"
)

func (e *TenantMemberRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TenantMemberRole(s)
	case string:
		*e = TenantMemberRole(s)
	default:
		return fmt.Errorf("unsupported scan type for TenantMemberRole: %T", src)
	}
	return nil
}

type NullTenantMemberRole struct {
	TenantMemberRole TenantMemberRole `json:"TenantMemberRole"`
	Valid            bool             `json:"valid"` // Valid is true if TenantMemberRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTenantMemberRole) Scan(value interface{}) error {
	if value == nil {
		ns.TenantMemberRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TenantMemberRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTenantMemberRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TenantMemberRole), nil
}

type TenantResourceLimitAlertType string

const (
	TenantResourceLimitAlertTypeAlarm     TenantResourceLimitAlertType = "Alarm"
	TenantResourceLimitAlertTypeExhausted TenantResourceLimitAlertType = "Exhausted"
)

func (e *TenantResourceLimitAlertType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TenantResourceLimitAlertType(s)
	case string:
		*e = TenantResourceLimitAlertType(s)
	default:
		return fmt.Errorf("unsupported scan type for TenantResourceLimitAlertType: %T", src)
	}
	return nil
}

type NullTenantResourceLimitAlertType struct {
	TenantResourceLimitAlertType TenantResourceLimitAlertType `json:"TenantResourceLimitAlertType"`
	Valid                        bool                         `json:"valid"` // Valid is true if TenantResourceLimitAlertType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTenantResourceLimitAlertType) Scan(value interface{}) error {
	if value == nil {
		ns.TenantResourceLimitAlertType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TenantResourceLimitAlertType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTenantResourceLimitAlertType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TenantResourceLimitAlertType), nil
}

type VcsProvider string

const (
	VcsProviderGITHUB VcsProvider = "GITHUB"
)

func (e *VcsProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VcsProvider(s)
	case string:
		*e = VcsProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for VcsProvider: %T", src)
	}
	return nil
}

type NullVcsProvider struct {
	VcsProvider VcsProvider `json:"VcsProvider"`
	Valid       bool        `json:"valid"` // Valid is true if VcsProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVcsProvider) Scan(value interface{}) error {
	if value == nil {
		ns.VcsProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VcsProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVcsProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VcsProvider), nil
}

type WorkflowRunStatus string

const (
	WorkflowRunStatusPENDING   WorkflowRunStatus = "PENDING"
	WorkflowRunStatusRUNNING   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusSUCCEEDED WorkflowRunStatus = "SUCCEEDED"
	WorkflowRunStatusFAILED    WorkflowRunStatus = "FAILED"
	WorkflowRunStatusQUEUED    WorkflowRunStatus = "QUEUED"
)

func (e *WorkflowRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkflowRunStatus(s)
	case string:
		*e = WorkflowRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkflowRunStatus: %T", src)
	}
	return nil
}

type NullWorkflowRunStatus struct {
	WorkflowRunStatus WorkflowRunStatus `json:"WorkflowRunStatus"`
	Valid             bool              `json:"valid"` // Valid is true if WorkflowRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkflowRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.WorkflowRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkflowRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkflowRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkflowRunStatus), nil
}

type APIToken struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	ExpiresAt   pgtype.Timestamp `json:"expiresAt"`
	Revoked     bool             `json:"revoked"`
	Name        pgtype.Text      `json:"name"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	NextAlertAt pgtype.Timestamp `json:"nextAlertAt"`
}

type Action struct {
	Description pgtype.Text `json:"description"`
	TenantId    pgtype.UUID `json:"tenantId"`
	ActionId    string      `json:"actionId"`
	ID          pgtype.UUID `json:"id"`
}

type ActionToWorker struct {
	B pgtype.UUID `json:"B"`
	A pgtype.UUID `json:"A"`
}

type ControllerPartition struct {
	ID            string           `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp `json:"updatedAt"`
	LastHeartbeat pgtype.Timestamp `json:"lastHeartbeat"`
}

type Dispatcher struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	LastHeartbeatAt pgtype.Timestamp `json:"lastHeartbeatAt"`
	IsActive        bool             `json:"isActive"`
}

type Event struct {
	ID                 pgtype.UUID      `json:"id"`
	CreatedAt          pgtype.Timestamp `json:"createdAt"`
	UpdatedAt          pgtype.Timestamp `json:"updatedAt"`
	DeletedAt          pgtype.Timestamp `json:"deletedAt"`
	Key                string           `json:"key"`
	TenantId           pgtype.UUID      `json:"tenantId"`
	ReplayedFromId     pgtype.UUID      `json:"replayedFromId"`
	Data               []byte           `json:"data"`
	AdditionalMetadata []byte           `json:"additionalMetadata"`
}

type GetGroupKeyRun struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	WorkerId          pgtype.UUID      `json:"workerId"`
	TickerId          pgtype.UUID      `json:"tickerId"`
	Status            StepRunStatus    `json:"status"`
	Input             []byte           `json:"input"`
	Output            pgtype.Text      `json:"output"`
	RequeueAfter      pgtype.Timestamp `json:"requeueAfter"`
	Error             pgtype.Text      `json:"error"`
	StartedAt         pgtype.Timestamp `json:"startedAt"`
	FinishedAt        pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt         pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt       pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason   pgtype.Text      `json:"cancelledReason"`
	CancelledError    pgtype.Text      `json:"cancelledError"`
	WorkflowRunId     pgtype.UUID      `json:"workflowRunId"`
	ScheduleTimeoutAt pgtype.Timestamp `json:"scheduleTimeoutAt"`
}

type Job struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	WorkflowVersionId pgtype.UUID      `json:"workflowVersionId"`
	Name              string           `json:"name"`
	Description       pgtype.Text      `json:"description"`
	Timeout           pgtype.Text      `json:"timeout"`
	Kind              JobKind          `json:"kind"`
}

type JobRun struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	JobId           pgtype.UUID      `json:"jobId"`
	TickerId        pgtype.UUID      `json:"tickerId"`
	Status          JobRunStatus     `json:"status"`
	Result          []byte           `json:"result"`
	StartedAt       pgtype.Timestamp `json:"startedAt"`
	FinishedAt      pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt       pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt     pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason pgtype.Text      `json:"cancelledReason"`
	CancelledError  pgtype.Text      `json:"cancelledError"`
	WorkflowRunId   pgtype.UUID      `json:"workflowRunId"`
}

type JobRunLookupData struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
	JobRunId  pgtype.UUID      `json:"jobRunId"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	Data      []byte           `json:"data"`
}

type LogLine struct {
	ID        int64            `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	StepRunId pgtype.UUID      `json:"stepRunId"`
	Message   string           `json:"message"`
	Level     LogLineLevel     `json:"level"`
	Metadata  []byte           `json:"metadata"`
}

type RateLimit struct {
	TenantId   pgtype.UUID      `json:"tenantId"`
	Key        string           `json:"key"`
	LimitValue int32            `json:"limitValue"`
	Value      int32            `json:"value"`
	Window     string           `json:"window"`
	LastRefill pgtype.Timestamp `json:"lastRefill"`
}

type SNSIntegration struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	TopicArn  string           `json:"topicArn"`
}

type SecurityCheckIdent struct {
	ID pgtype.UUID `json:"id"`
}

type Service struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	TenantId    pgtype.UUID      `json:"tenantId"`
}

type ServiceToWorker struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type SlackAppWebhook struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	TeamId      string           `json:"teamId"`
	TeamName    string           `json:"teamName"`
	ChannelId   string           `json:"channelId"`
	ChannelName string           `json:"channelName"`
	WebhookURL  []byte           `json:"webhookURL"`
}

type Step struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	ReadableId      pgtype.Text      `json:"readableId"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	JobId           pgtype.UUID      `json:"jobId"`
	ActionId        string           `json:"actionId"`
	Timeout         pgtype.Text      `json:"timeout"`
	CustomUserData  []byte           `json:"customUserData"`
	Retries         int32            `json:"retries"`
	ScheduleTimeout string           `json:"scheduleTimeout"`
}

type StepOrder struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type StepRateLimit struct {
	Units        int32       `json:"units"`
	StepId       pgtype.UUID `json:"stepId"`
	RateLimitKey string      `json:"rateLimitKey"`
	TenantId     pgtype.UUID `json:"tenantId"`
}

type StepRun struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	JobRunId          pgtype.UUID      `json:"jobRunId"`
	StepId            pgtype.UUID      `json:"stepId"`
	Order             int64            `json:"order"`
	WorkerId          pgtype.UUID      `json:"workerId"`
	TickerId          pgtype.UUID      `json:"tickerId"`
	Status            StepRunStatus    `json:"status"`
	Input             []byte           `json:"input"`
	Output            []byte           `json:"output"`
	RequeueAfter      pgtype.Timestamp `json:"requeueAfter"`
	ScheduleTimeoutAt pgtype.Timestamp `json:"scheduleTimeoutAt"`
	Error             pgtype.Text      `json:"error"`
	StartedAt         pgtype.Timestamp `json:"startedAt"`
	FinishedAt        pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt         pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt       pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason   pgtype.Text      `json:"cancelledReason"`
	CancelledError    pgtype.Text      `json:"cancelledError"`
	InputSchema       []byte           `json:"inputSchema"`
	CallerFiles       []byte           `json:"callerFiles"`
	GitRepoBranch     pgtype.Text      `json:"gitRepoBranch"`
	RetryCount        int32            `json:"retryCount"`
	SemaphoreReleased bool             `json:"semaphoreReleased"`
}

type StepRunEvent struct {
	ID            int64                `json:"id"`
	TimeFirstSeen pgtype.Timestamp     `json:"timeFirstSeen"`
	TimeLastSeen  pgtype.Timestamp     `json:"timeLastSeen"`
	StepRunId     pgtype.UUID          `json:"stepRunId"`
	Reason        StepRunEventReason   `json:"reason"`
	Severity      StepRunEventSeverity `json:"severity"`
	Message       string               `json:"message"`
	Count         int32                `json:"count"`
	Data          []byte               `json:"data"`
}

type StepRunOrder struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type StepRunResultArchive struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	StepRunId       pgtype.UUID      `json:"stepRunId"`
	Order           int64            `json:"order"`
	Input           []byte           `json:"input"`
	Output          []byte           `json:"output"`
	Error           pgtype.Text      `json:"error"`
	StartedAt       pgtype.Timestamp `json:"startedAt"`
	FinishedAt      pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt       pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt     pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason pgtype.Text      `json:"cancelledReason"`
	CancelledError  pgtype.Text      `json:"cancelledError"`
}

type StreamEvent struct {
	ID        int64            `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	StepRunId pgtype.UUID      `json:"stepRunId"`
	Message   []byte           `json:"message"`
	Metadata  []byte           `json:"metadata"`
}

type Tenant struct {
	ID                    pgtype.UUID      `json:"id"`
	CreatedAt             pgtype.Timestamp `json:"createdAt"`
	UpdatedAt             pgtype.Timestamp `json:"updatedAt"`
	DeletedAt             pgtype.Timestamp `json:"deletedAt"`
	Name                  string           `json:"name"`
	Slug                  string           `json:"slug"`
	AnalyticsOptOut       bool             `json:"analyticsOptOut"`
	AlertMemberEmails     bool             `json:"alertMemberEmails"`
	ControllerPartitionId pgtype.Text      `json:"controllerPartitionId"`
	WorkerPartitionId     pgtype.Text      `json:"workerPartitionId"`
	DataRetentionPeriod   string           `json:"dataRetentionPeriod"`
}

type TenantAlertEmailGroup struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	Emails    string           `json:"emails"`
}

type TenantAlertingSettings struct {
	ID                              pgtype.UUID      `json:"id"`
	CreatedAt                       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt                       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt                       pgtype.Timestamp `json:"deletedAt"`
	TenantId                        pgtype.UUID      `json:"tenantId"`
	MaxFrequency                    string           `json:"maxFrequency"`
	LastAlertedAt                   pgtype.Timestamp `json:"lastAlertedAt"`
	TickerId                        pgtype.UUID      `json:"tickerId"`
	EnableExpiringTokenAlerts       bool             `json:"enableExpiringTokenAlerts"`
	EnableWorkflowRunFailureAlerts  bool             `json:"enableWorkflowRunFailureAlerts"`
	EnableTenantResourceLimitAlerts bool             `json:"enableTenantResourceLimitAlerts"`
}

type TenantInviteLink struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	TenantId     pgtype.UUID      `json:"tenantId"`
	InviterEmail string           `json:"inviterEmail"`
	InviteeEmail string           `json:"inviteeEmail"`
	Expires      pgtype.Timestamp `json:"expires"`
	Status       InviteLinkStatus `json:"status"`
	Role         TenantMemberRole `json:"role"`
}

type TenantMember struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	UserId    pgtype.UUID      `json:"userId"`
	Role      TenantMemberRole `json:"role"`
}

type TenantResourceLimit struct {
	ID               pgtype.UUID      `json:"id"`
	CreatedAt        pgtype.Timestamp `json:"createdAt"`
	UpdatedAt        pgtype.Timestamp `json:"updatedAt"`
	Resource         LimitResource    `json:"resource"`
	TenantId         pgtype.UUID      `json:"tenantId"`
	LimitValue       int32            `json:"limitValue"`
	AlarmValue       pgtype.Int4      `json:"alarmValue"`
	Value            int32            `json:"value"`
	Window           pgtype.Text      `json:"window"`
	LastRefill       pgtype.Timestamp `json:"lastRefill"`
	CustomValueMeter bool             `json:"customValueMeter"`
}

type TenantResourceLimitAlert struct {
	ID              pgtype.UUID                  `json:"id"`
	CreatedAt       pgtype.Timestamp             `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp             `json:"updatedAt"`
	ResourceLimitId pgtype.UUID                  `json:"resourceLimitId"`
	TenantId        pgtype.UUID                  `json:"tenantId"`
	Resource        LimitResource                `json:"resource"`
	AlertType       TenantResourceLimitAlertType `json:"alertType"`
	Value           int32                        `json:"value"`
	Limit           int32                        `json:"limit"`
}

type TenantVcsProvider struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	VcsProvider VcsProvider      `json:"vcsProvider"`
	Config      []byte           `json:"config"`
}

type TenantWorkerPartition struct {
	ID            string           `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp `json:"updatedAt"`
	LastHeartbeat pgtype.Timestamp `json:"lastHeartbeat"`
}

type Ticker struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	LastHeartbeatAt pgtype.Timestamp `json:"lastHeartbeatAt"`
	IsActive        bool             `json:"isActive"`
}

type User struct {
	ID            pgtype.UUID      `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp `json:"updatedAt"`
	DeletedAt     pgtype.Timestamp `json:"deletedAt"`
	Email         string           `json:"email"`
	EmailVerified bool             `json:"emailVerified"`
	Name          pgtype.Text      `json:"name"`
}

type UserOAuth struct {
	ID             pgtype.UUID      `json:"id"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
	UpdatedAt      pgtype.Timestamp `json:"updatedAt"`
	UserId         pgtype.UUID      `json:"userId"`
	Provider       string           `json:"provider"`
	ProviderUserId string           `json:"providerUserId"`
	ExpiresAt      pgtype.Timestamp `json:"expiresAt"`
	AccessToken    []byte           `json:"accessToken"`
	RefreshToken   []byte           `json:"refreshToken"`
}

type UserPassword struct {
	Hash   string      `json:"hash"`
	UserId pgtype.UUID `json:"userId"`
}

type UserSession struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	UserId    pgtype.UUID      `json:"userId"`
	Data      []byte           `json:"data"`
	ExpiresAt pgtype.Timestamp `json:"expiresAt"`
}

type WebhookWorker struct {
	ID         pgtype.UUID      `json:"id"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	UpdatedAt  pgtype.Timestamp `json:"updatedAt"`
	Name       string           `json:"name"`
	Secret     string           `json:"secret"`
	Url        string           `json:"url"`
	TokenValue pgtype.Text      `json:"tokenValue"`
	Deleted    bool             `json:"deleted"`
	TokenId    pgtype.UUID      `json:"tokenId"`
	TenantId   pgtype.UUID      `json:"tenantId"`
}

type WebhookWorkerWorkflow struct {
	ID              pgtype.UUID `json:"id"`
	WebhookWorkerId pgtype.UUID `json:"webhookWorkerId"`
	WorkflowId      pgtype.UUID `json:"workflowId"`
}

type Worker struct {
	ID                      pgtype.UUID      `json:"id"`
	CreatedAt               pgtype.Timestamp `json:"createdAt"`
	UpdatedAt               pgtype.Timestamp `json:"updatedAt"`
	DeletedAt               pgtype.Timestamp `json:"deletedAt"`
	TenantId                pgtype.UUID      `json:"tenantId"`
	LastHeartbeatAt         pgtype.Timestamp `json:"lastHeartbeatAt"`
	Name                    string           `json:"name"`
	DispatcherId            pgtype.UUID      `json:"dispatcherId"`
	MaxRuns                 int32            `json:"maxRuns"`
	IsActive                bool             `json:"isActive"`
	LastListenerEstablished pgtype.Timestamp `json:"lastListenerEstablished"`
	IsPaused                bool             `json:"isPaused"`
}

type WorkerSemaphore struct {
	WorkerId pgtype.UUID `json:"workerId"`
	Slots    int32       `json:"slots"`
}

type WorkerSemaphoreSlot struct {
	ID        pgtype.UUID `json:"id"`
	WorkerId  pgtype.UUID `json:"workerId"`
	StepRunId pgtype.UUID `json:"stepRunId"`
}

type Workflow struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
}

type WorkflowConcurrency struct {
	ID                    pgtype.UUID              `json:"id"`
	CreatedAt             pgtype.Timestamp         `json:"createdAt"`
	UpdatedAt             pgtype.Timestamp         `json:"updatedAt"`
	WorkflowVersionId     pgtype.UUID              `json:"workflowVersionId"`
	GetConcurrencyGroupId pgtype.UUID              `json:"getConcurrencyGroupId"`
	MaxRuns               int32                    `json:"maxRuns"`
	LimitStrategy         ConcurrencyLimitStrategy `json:"limitStrategy"`
}

type WorkflowRun struct {
	CreatedAt          pgtype.Timestamp  `json:"createdAt"`
	UpdatedAt          pgtype.Timestamp  `json:"updatedAt"`
	DeletedAt          pgtype.Timestamp  `json:"deletedAt"`
	TenantId           pgtype.UUID       `json:"tenantId"`
	WorkflowVersionId  pgtype.UUID       `json:"workflowVersionId"`
	Status             WorkflowRunStatus `json:"status"`
	Error              pgtype.Text       `json:"error"`
	StartedAt          pgtype.Timestamp  `json:"startedAt"`
	FinishedAt         pgtype.Timestamp  `json:"finishedAt"`
	ConcurrencyGroupId pgtype.Text       `json:"concurrencyGroupId"`
	DisplayName        pgtype.Text       `json:"displayName"`
	ID                 pgtype.UUID       `json:"id"`
	ChildIndex         pgtype.Int4       `json:"childIndex"`
	ChildKey           pgtype.Text       `json:"childKey"`
	ParentId           pgtype.UUID       `json:"parentId"`
	ParentStepRunId    pgtype.UUID       `json:"parentStepRunId"`
	AdditionalMetadata []byte            `json:"additionalMetadata"`
	Duration           pgtype.Int4       `json:"duration"`
}

type WorkflowRunTriggeredBy struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	DeletedAt    pgtype.Timestamp `json:"deletedAt"`
	TenantId     pgtype.UUID      `json:"tenantId"`
	EventId      pgtype.UUID      `json:"eventId"`
	CronParentId pgtype.UUID      `json:"cronParentId"`
	CronSchedule pgtype.Text      `json:"cronSchedule"`
	ScheduledId  pgtype.UUID      `json:"scheduledId"`
	Input        []byte           `json:"input"`
	ParentId     pgtype.UUID      `json:"parentId"`
}

type WorkflowTag struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	Name      string           `json:"name"`
	Color     string           `json:"color"`
}

type WorkflowToWorkflowTag struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type WorkflowTriggerCronRef struct {
	ParentId pgtype.UUID `json:"parentId"`
	Cron     string      `json:"cron"`
	TickerId pgtype.UUID `json:"tickerId"`
	Input    []byte      `json:"input"`
	Enabled  bool        `json:"enabled"`
}

type WorkflowTriggerEventRef struct {
	ParentId pgtype.UUID `json:"parentId"`
	EventKey string      `json:"eventKey"`
}

type WorkflowTriggerScheduledRef struct {
	ID                  pgtype.UUID      `json:"id"`
	ParentId            pgtype.UUID      `json:"parentId"`
	TriggerAt           pgtype.Timestamp `json:"triggerAt"`
	TickerId            pgtype.UUID      `json:"tickerId"`
	Input               []byte           `json:"input"`
	ChildIndex          pgtype.Int4      `json:"childIndex"`
	ChildKey            pgtype.Text      `json:"childKey"`
	ParentStepRunId     pgtype.UUID      `json:"parentStepRunId"`
	ParentWorkflowRunId pgtype.UUID      `json:"parentWorkflowRunId"`
}

type WorkflowTriggers struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	WorkflowVersionId pgtype.UUID      `json:"workflowVersionId"`
	TenantId          pgtype.UUID      `json:"tenantId"`
}

type WorkflowVersion struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	Version         pgtype.Text      `json:"version"`
	Order           int64            `json:"order"`
	WorkflowId      pgtype.UUID      `json:"workflowId"`
	Checksum        string           `json:"checksum"`
	ScheduleTimeout string           `json:"scheduleTimeout"`
	OnFailureJobId  pgtype.UUID      `json:"onFailureJobId"`
}
