package engine

type Command interface {
	Execute(handler Handler)
}

// type FuncCommand func()

// func (fCmd FuncCommand) Execute(handler Handler) {
// 	fCmd()
// }
