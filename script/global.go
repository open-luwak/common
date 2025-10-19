package script

const (
	inputKey   = "input"
	envKey     = "env"
	metaKey    = "metas"
	serverKey  = "server"
	sessionKey = "session"
	outputKey  = "output"
	ctxKey     = "ctx"

	requestHeaderKey  = "requestHeaders"
	responseHeaderKey = "responseHeaders"
	rolesKey          = "roles"
)

type GlobalThis map[string]any

func (m GlobalThis) Input() any {
	return m[inputKey]
}

func (m GlobalThis) SetInput(input any) {
	m[inputKey] = input
}

func (m GlobalThis) Output() any {
	return m[outputKey]
}

func (m GlobalThis) SetOutput(output any) {
	m[outputKey] = output
}

func (m GlobalThis) Metas() map[string]any {
	if v, ok := m[metaKey]; ok {
		if metas, ok := v.(map[string]any); ok {
			return metas
		}
	}

	metas := make(map[string]any)
	m[metaKey] = metas
	return metas
}

func (m GlobalThis) SetMetas(metas map[string]any) {
	m[metaKey] = metas
}

func (m GlobalThis) Server() map[string]any {
	if v, ok := m[serverKey]; ok {
		if server, ok := v.(map[string]any); ok {
			return server
		}
	}

	server := make(map[string]any)
	m[serverKey] = server
	return server
}

func (m GlobalThis) SetServer(server map[string]any) {
	m[serverKey] = server
}

func (m GlobalThis) Session() map[string]any {
	if v, ok := m[sessionKey]; ok {
		if sess, ok := v.(map[string]any); ok {
			return sess
		}
	}

	sess := make(map[string]any)
	m[sessionKey] = sess
	return sess
}

func (m GlobalThis) SetSession(sess map[string]any) {
	m[sessionKey] = sess
}

func (m GlobalThis) Env() map[string]any {
	if v, ok := m[envKey]; ok {
		if env, ok := v.(map[string]any); ok {
			return env
		}
	}

	env := make(map[string]any)
	m[envKey] = env
	return env
}

func (m GlobalThis) SetEnv(env map[string]any) {
	m[envKey] = env
}

func (m GlobalThis) SetCtx(key string, value any) {
	if v, ok := m[ctxKey]; ok {
		if ctx, ok := v.(map[string]any); ok {
			ctx[key] = value
			return
		}
	}

	ctx := make(map[string]any)
	ctx[key] = value
	m[ctxKey] = ctx
}
