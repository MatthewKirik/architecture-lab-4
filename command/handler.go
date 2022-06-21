package command

type Handler interface {
	Post(cmd Command)
}
