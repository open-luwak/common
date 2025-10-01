package metadata

const (
	CheckpointType = "CP"
	CheckpointDir  = "checkpoint"

	DataPreprocessType = "DPP"
	DataPreprocessDir  = "data_preprocessor"

	MainType = "MAIN"
	MainDir  = "main"

	DatabaseUpdaterType = "DU"
	DatabaseUpdaterDir  = "database_updater"

	DatabaseInserterType = "DI"
	DatabaseInserterDir  = "database_inserter"

	EventBroadcasterType = "EB"
	EventBroadcasterDir  = "event_broadcaster"

	ResponseEnricherType = "RE"
	ResponseEnricherDir  = "response_enricher"

	ConditionalRoleType = "CR"
	ConditionalRoleDir  = "conditional_role"
)

const (
	BeforeHookType = "H_BEFORE"
	BeforeHookDir  = "before"

	AfterSuccessHookType = "H_AFTER_SUCCESS"
	AfterSuccessHookDir  = "after_success"

	AfterErrorHookType = "H_AFTER_ERROR"
	AfterErrorHookDir  = "after_error"

	FinallyHookType = "H_FINALLY"
	FinallyHookDir  = "finally"
)
