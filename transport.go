package myrtio

// Transport represent MyrtIO transport
type Transport interface {
	// RunAction sends command to device and return response
	RunAction(message *Message) (*Message, error)
	// Close connection with device
	Close() error
}
