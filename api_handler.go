package common

type Handler interface {
	Handle(APIContext) error
}

type HandlerFunc func(APIContext) error

func (hf HandlerFunc) Handle(ctx APIContext) error {
	return hf(ctx)
}

type Middleware struct {
	Name     string
	Priority int
	Handler  func(Handler) Handler
}
