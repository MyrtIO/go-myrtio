package myrtio

type Transport interface {
	RunAction(message *Message) (*Message, error)
	Close() error
}
