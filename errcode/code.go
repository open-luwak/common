package errcode

// gateway
const (
	JsonRpcInvalidRequest = "AD.JsonRpc.InvalidRequest"
	JsonRpcMethodNotFound = "AD.JsonRpc.MethodNotFound"

	EndpointEnvExpress    = "SP.Endpoint.EnvExpression"
	EndpointFilterExpress = "SP.Endpoint.FilterExpression"
	EndpointValueExpress  = "SP.Endpoint.ValueExpression"

	APINotFound         = "DEV.API.NotFound"
	APIStateDraft       = "DEV.API.State.Draft"
	APIStateOffline     = "DEV.API.State.Offline"
	APIStateRemoved     = "DEV.API.State.Removed"
	APIRequestIdUsed    = "AD.API.RequestIdUsed"
	APIInvalidTimestamp = "AD.API.InvalidTimestamp"
	APIInvalidSignature = "AD.API.InvalidSignature"
	APIInvalidAppKey    = "AD.API.InvalidAppKey"
	APIUnauthorized     = "EU.API.Unauthorized"
)

// firewall
const (
	FwMissingMetasParameter = "EU.CheckSign.MissingMetasParameter"
	FwInvalidAppKey         = "EU.CheckSign.InvalidAppKey"
	FwInvalidSignature      = "EU.CheckSign.InvalidSignature"
	FwAclRoleNotAllowed     = "SP.ACL.RoleNotAllowed"
)

// core
const (
	ScriptParse     = "BE.Script.Parse"
	ScriptExecution = "BE.Script.Execution"
	ScriptTimeout   = "BE.Script.Timeout"
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
