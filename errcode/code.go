package errcode

// end user = EU
// app developer = AD
// biz error = BIZ
// operations = OPS

const (
	EndpointEnvExpress         = "AD.API.Endpoint.EnvExpression"
	EndpointFilterExpress      = "AD.API.Endpoint.FilterExpression"
	EndpointValueExpress       = "AD.API.Endpoint.ValueExpression"
	APIScheduledTaskNotAllowed = "AD.API.ScheduledTaskNotAllowed"
	APIInvalidMethodName       = "AD.API.InvalidMethodName"
	APINotFound                = "AD.API.NotFound"
	APIStateDraft              = "AD.API.State.Draft"
	APIStateOffline            = "AD.API.State.Offline"
	APIStateRemoved            = "AD.API.State.Removed"
	APIRequestIdUsed           = "AD.API.RequestIdUsed"
	APIMissingMetasParameter   = "AD.API.MissingMetasParameter"
	APIInvalidTimestamp        = "AD.API.InvalidTimestamp"
	APIInvalidSignature        = "AD.API.InvalidSignature"
	APIUnauthorized            = "EU.API.Unauthorized"

	APPInvalidAppKey       = "AD.APP.Invalid.AppKey"
	APPInvalidSoftwareName = "AD.APP.Invalid.softwareName"
	APPEmptyAppKey         = "AD.APP.Empty.AppKey"
	APPEmptySoftwareName   = "AD.APP.Empty.SoftwareName"

	SoftwarePrivilegeNotGranted = "AD.Software.PrivilegeNotGranted"

	AclRoleNotAllowed = "AD.ACL.RoleNotAllowed"

	ScriptParse                = "AD.Script.Parse"
	ScriptExecution            = "AD.Script.Execution"
	ScriptTimeout              = "AD.Script.Timeout"
	ScriptNotSupported         = "AD.Script.NotSupported"
	ScriptExecutorNotAvailable = "AD.Script.ExecutorNotAvailable"

	DatabaseQueryFailed = "AD.Database.QueryFailed"

	DbConnect         = "OPS.DbConnect"
	RedisConnect      = "OPS.RedisConnect"
	MqConnect         = "OPS.MqConnect"
	LicenseExpired    = "OPS.LicenseExpired"
	LicenseInvalid    = "OPS.LicenseInvalid"
	LicenseRestricted = "OPS.LicenseRestricted"
)
