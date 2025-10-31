package api

import (
	"github.com/spf13/cast"
)

const (
	tokenKey     = "token"
	signKey      = "signature"
	appKeyKey    = "appKey"
	timestampKey = "timestamp"
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

func (m Metas) Signature() string {
	return cast.ToString(m[signKey])
}

func (m Metas) SetSignature(sign string) {
	m[signKey] = sign
}

func (m Metas) Timestamp() string {
	return cast.ToString(m[timestampKey])
}

func (m Metas) SetTimestamp(timestamp string) {
	m[timestampKey] = timestamp
}
