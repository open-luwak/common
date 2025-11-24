package script

const GlobalKey = "globalThis"

const (
	inputKey   = "input"
	envKey     = "env"
	metaKey    = "meta"
	serverKey  = "server"
	sessionKey = "session"
	outputKey  = "output"
	ctxKey     = "ctx"
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

func (m GlobalThis) Meta() map[string]any {
	if v, ok := m[metaKey]; ok {
		if metas, ok := v.(map[string]any); ok {
			return metas
		}
	}

	metas := make(map[string]any)
	m[metaKey] = metas
	return metas
}

func (m GlobalThis) SetMeta(metas map[string]any) {
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

func (m GlobalThis) CtxData() map[string]any {
	if v, ok := m[ctxKey]; ok {
		if ctx, ok := v.(map[string]any); ok {
			return ctx
		}
	}

	ctx := make(map[string]any)
	m[ctxKey] = ctx
	return ctx
}

func (m GlobalThis) SetCtxData(key string, value any) {
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

func (m GlobalThis) Copy(global map[string]any) {
	if value, ok := global[metaKey]; ok {
		if meta, ok := value.(map[string]any); ok {
			newMeta := make(map[string]any)
			for k, v := range meta {
				newMeta[k] = v
			}
			m.SetMeta(newMeta)
			return
		}
	}
}
