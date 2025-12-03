package errcode

const (
	APIUnauthorized             = "EU.API.Unauthorized"                // 401 Unauthorized
	SoftwarePrivilegeNotGranted = "AD.Software.Permission.NotAssigned" // 403 Forbidden
	APIPermissionDeniedAccess   = "EU.API.Permission.Denied"           // 403
	APIRateLimitExceeded        = "SP.API.RateLimit.Exceeded"          // 429 Too Many Requests
)
