package engine

type Command interface {
	Execute(handler Handler)
}
