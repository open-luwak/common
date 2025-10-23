package script

import (
	"github.com/spf13/cast"
)

const (
	apiNameKey     = "apiName"
	remoteAddrKey  = "remoteAddr"
	remoteHostKey  = "remoteHost"
	requestTimeKey = "reqTime"
	requestIdKey   = "reqId"
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
