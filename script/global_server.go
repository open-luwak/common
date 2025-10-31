package script

import (
	"github.com/spf13/cast"
)

const (
	remoteAddrKey     = "remoteAddr"
	remoteHostKey     = "remoteHost"
	apiNameKey        = "reqApi"
	requestIdKey      = "reqId"
	requestTimeKey    = "reqTime"
	requestHeaderKey  = "reqHeader"
	responseHeaderKey = "respHeader"
)

type Server map[string]any

func (m Server) ApiName() string {
	return cast.ToString(m[apiNameKey])
}

func (m Server) SetApiName(apiName string) {
	m[apiNameKey] = apiName
}

func (m Server) RemoteAddr() string {
	return cast.ToString(m[remoteAddrKey])
}

func (m Server) SetRemoteAddr(remoteAddr string) {
	m[remoteAddrKey] = remoteAddr
}

func (m Server) RemoteHost() string {
	return cast.ToString(m[remoteHostKey])
}

func (m Server) SetRemoteHost(remoteHost string) {
	m[remoteHostKey] = remoteHost
}

func (m Server) RequestTime() int64 {
	return cast.ToInt64(m[requestTimeKey])
}

func (m Server) SetRequestTime(timestamp int64) {
	m[requestTimeKey] = timestamp
}

func (m Server) RequestId() string {
	return cast.ToString(m[requestIdKey])
}

func (m Server) SetRequestId(remoteHost string) {
	m[requestIdKey] = remoteHost
}

func (m Server) ReqHeader() map[string]string {
	if value, ok := m[requestHeaderKey]; ok {
		if v, ok := value.(map[string]string); ok {
			return v
		}
	}

	return make(map[string]string)
}

func (m Server) SetReqHeader(header map[string]string) {
	m[requestHeaderKey] = header
}

func (m Server) RespHeader() map[string]string {
	if value, ok := m[responseHeaderKey]; ok {
		if v, ok := value.(map[string]string); ok {
			return v
		}
	}

	return make(map[string]string)
}

func (m Server) AddRespHeader(key, value string) {
	if respHeader, ok := m[responseHeaderKey]; ok {
		if v, ok := respHeader.(map[string]string); ok {
			v[key] = value
		}
	}
	respHeader := make(map[string]string)
	respHeader[key] = value
	m[responseHeaderKey] = respHeader
}
