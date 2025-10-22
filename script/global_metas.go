package script

import (
	"github.com/spf13/cast"
)

const (
	appKeyKey = "appKey"
	tokenKey  = "token"
)

type Metas map[string]any

func (m Server) AppKey() string {
	return cast.ToString(m[appKeyKey])
}

func (m Server) SetAppKey(appKey string) {
	m[appKeyKey] = appKey
}

func (m Server) Token() string {
	return cast.ToString(m[tokenKey])
}

func (m Server) SetToken(accessToken string) {
	m[tokenKey] = accessToken
}
