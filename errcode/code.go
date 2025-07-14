package errcode

// end user = EU
// app developer = AD
// biz error = BIZ
// operations = OPS

// gateway
const (
	EndpointEnvExpress    = "BIZ.Endpoint.EnvExpression"
	EndpointFilterExpress = "BIZ.Endpoint.FilterExpression"
	EndpointValueExpress  = "BIZ.Endpoint.ValueExpression"

	APIScheduledTaskNotAllowed = "AD.API.ScheduledTaskNotAllowed"
	APIInvalidMethodName       = "AD.API.InvalidMethodName"
	APINotFound                = "AD.API.NotFound"
	APIStateDraft              = "AD.API.State.Draft"
	APIStateOffline            = "AD.API.State.Offline"
	APIStateRemoved            = "AD.API.State.Removed"
	APIRequestIdUsed           = "AD.API.RequestIdUsed"
	APIInvalidTimestamp        = "AD.API.InvalidTimestamp"
	APIInvalidSignature        = "AD.API.InvalidSignature"
	APIInvalidAppKey           = "AD.API.InvalidAppKey"
	APIUnauthorized            = "EU.API.Unauthorized"
)

// firewall
const (
	FwMissingMetasParameter = "AD.CheckSign.MissingMetasParameter"
	FwInvalidAppKey         = "AD.CheckSign.InvalidAppKey"
	FwInvalidSignature      = "AD.CheckSign.InvalidSignature"
	FwAclRoleNotAllowed     = "BIZ.ACL.RoleNotAllowed"
)

// core
const (
	ScriptParse                = "BIZ.Script.Parse"
	ScriptExecution            = "BIZ.Script.Execution"
	ScriptTimeout              = "BIZ.Script.Timeout"
	ScriptNotSupported         = "AD.Script.NotSupported"
	ScriptExecutorNotAvailable = "AD.Script.ExecutorNotAvailable"

	SqlQueryFailed = "AD.SQL.QueryFailed"
)

// infrastructure
const (
	DbConnect    = "OPS.DbConnect"
	RedisConnect = "OPS.RedisConnect"
	MqConnect    = "OPS.MqConnect"
)

// license
const (
	LicenseExpired    = "OPS.LicenseExpired"
	LicenseInvalid    = "OPS.LicenseInvalid"
	LicenseRestricted = "OPS.LicenseRestricted"
)
