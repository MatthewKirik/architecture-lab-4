package engine

type Handler interface {
	Post(cmd Command)
}
