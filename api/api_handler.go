package api

type Handler interface {
	Handle(Context) error
}

type HandlerFunc func(Context) error

func (hf HandlerFunc) Handle(ctx Context) error {
	return hf(ctx)
}

type Middleware struct {
	Name     string
	Priority int
	Handler  func(Handler) Handler
}
