package engine

type Command interface {
	Execute(handler Handler)
}

type FuncCommand func(handler Handler)

func (fCmd FuncCommand) Execute(handler Handler) {
	fCmd(handler)
}
