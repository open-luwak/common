package common

type LogicFunc func(ctx APIContext)

type LogicMiddleware func(LogicFunc) LogicFunc

func LogicChain(handler LogicFunc, middleware ...LogicMiddleware) LogicFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
