package handler

// HandlerRegistrator is an interface that defines the methods that the handler registration requires.
type HandlerRegistrator interface {
	On(string, any) error
}
