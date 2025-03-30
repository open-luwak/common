package common

type Handler interface {
	Process(APIContext) error
}

type HandlerFunc func(APIContext) error

func (hf HandlerFunc) Process(ctx APIContext) error {
	return hf(ctx)
}
