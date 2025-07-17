package errcode

// end user = EU
// app developer = AD
// biz error = BIZ
// operations = OPS

const (
	EndpointEnvExpress         = "AD.API.EndpointEnvExpression"
	EndpointFilterExpress      = "AD.API.EndpointFilterExpression"
	EndpointValueExpress       = "AD.API.EndpointValueExpression"
	APIScheduledTaskNotAllowed = "AD.API.ScheduledTaskNotAllowed"
	APIInvalidMethodName       = "AD.API.InvalidMethodName"
	APINotFound                = "AD.API.NotFound"
	APIStateDraft              = "AD.API.StateDraft"
	APIStateOffline            = "AD.API.StateOffline"
	APIStateRemoved            = "AD.API.StateRemoved"
	APIRequestIdUsed           = "AD.API.RequestIdUsed"
	APIMissingMetasParameter   = "AD.API.MissingMetasParameter"
	APIInvalidTimestamp        = "AD.API.InvalidTimestamp"
	APIInvalidSignature        = "AD.API.InvalidSignature"
	APIUnauthorized            = "EU.API.Unauthorized"

	APPInvalidAppKey     = "AD.APP.InvalidAppKey"
	APPEmptyAppKey       = "AD.APP.EmptyAppKey"
	APPEmptySoftwareName = "AD.APP.EmptySoftwareName"

	SoftwareInvalidSoftwareName = "AD.Software.InvalidSoftwareName"
	SoftwarePrivilegeNotGranted = "AD.Software.PrivilegeNotGranted"

	AclRoleNotAllowed = "AD.ACL.RoleNotAllowed"

	ScriptParseFailed          = "AD.Script.ParseFailed"
	ScriptExecutionFailed      = "AD.Script.ExecutionFailed"
	ScriptExecutionTimeout     = "AD.Script.ExecutionTimeout"
	ScriptLangNotSupported     = "AD.Script.LangNotSupported"
	ScriptExecutorNotAvailable = "AD.Script.ExecutorNotAvailable"

	DatabaseQueryFailed = "AD.Database.QueryFailed"

	DbConnectFailed    = "OPS.Database.ConnectFailed"
	RedisConnectFailed = "OPS.Redis.ConnectFailed"
	MqConnectFailed    = "OPS.Mq.ConnectFailed"
	LicenseExpired     = "OPS.License.Expired"
	LicenseInvalid     = "OPS.License.Invalid"
	LicenseRestricted  = "OPS.License.Restricted"
)
