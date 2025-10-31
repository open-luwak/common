package script

import (
	"github.com/spf13/cast"
)

const (
	tokenKey     = "token"
	sessionIdKey = "sessionId"
	appKeyKey    = "appKey"
	softwareKey  = "software"
	rolesKey     = "roles"
)

type Meta map[string]any

func (m Meta) AppKey() string {
	return cast.ToString(m[appKeyKey])
}

func (m Meta) SetAppKey(appKey string) {
	m[appKeyKey] = appKey
}

func (m Meta) Token() string {
	return cast.ToString(m[tokenKey])
}

func (m Meta) SetToken(accessToken string) {
	m[tokenKey] = accessToken
}

func (m Meta) SessionId() string {
	return cast.ToString(m[sessionIdKey])
}

func (m Meta) SetSessionId(sessionId string) {
	m[sessionIdKey] = sessionId
}

func (m Meta) Software() string {
	return cast.ToString(m[softwareKey])
}

func (m Meta) SetSoftware(software string) {
	m[softwareKey] = software
}

func (m Meta) Roles() []string {
	return cast.ToStringSlice(m[rolesKey])
}

func (m Meta) AddRoles(roles []string) {
	if v, ok := m[rolesKey]; ok {
		if vv, ok := v.([]string); ok {
			vv = append(vv, roles...)
		}
	}
	m[rolesKey] = roles
}
