package script

import (
	"github.com/spf13/cast"
)

const (
	appKeyKey = "appKey"
	tokenKey  = "token"
)

type Metas map[string]any

func (m Metas) AppKey() string {
	return cast.ToString(m[appKeyKey])
}

func (m Metas) SetAppKey(appKey string) {
	m[appKeyKey] = appKey
}

func (m Metas) Token() string {
	return cast.ToString(m[tokenKey])
}

func (m Metas) SetToken(accessToken string) {
	m[tokenKey] = accessToken
}
