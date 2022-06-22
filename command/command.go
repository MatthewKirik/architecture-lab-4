package command

type Command interface {
	Execute(handler Handler)
}
