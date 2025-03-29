package common

type Handler interface {
	Process(APIContext) error
}

type HandlerFunc func(APIContext) error

func (hf HandlerFunc) Process(ctx APIContext) error {
	return hf(ctx)
}

type Chain struct {
	list []func(Handler) Handler
}

func NewChain() *Chain {
	return &Chain{
		list: make([]func(Handler) Handler, 0, 8),
	}
}

func (c *Chain) Append(middleware func(Handler) Handler) {
	c.list = append(c.list, middleware)
}

func (c *Chain) Apply(h Handler) Handler {
	for i := len(c.list) - 1; i >= 0; i-- {
		h = c.list[i](h)
	}
	return h
}
