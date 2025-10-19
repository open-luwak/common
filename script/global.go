package script

const (
	inputKey          = "input"
	envKey            = "env"
	metaKey           = "metas"
	serverKey         = "server"
	sessionKey        = "session"
	requestHeaderKey  = "requestHeaders"
	responseHeaderKey = "responseHeaders"
	outputKey         = "output"
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
