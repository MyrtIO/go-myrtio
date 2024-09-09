// Package serial contains utilities for MyrtIO API through serial connection
package serial

import (
	"sync"
	"time"

	"github.com/MyrtIO/myrtio-go"

	"go.bug.st/serial"
)

// MyrtIO typical serial timings
const (
	startDelayMs   = 500
	commandDelayMs = 20
	readDelayMs    = 50
	readTimeoutMs  = 100
)

// Transport represents MyrtIO serial transport
type Transport struct {
	port serial.Port
	mu   sync.Mutex
}

// Close connection with device
func (t *Transport) Close() error {
	return t.port.Close()
}

// RunAction sends command to device and return response
func (t *Transport) RunAction(message *myrtio.Message) (*myrtio.Message, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	_, err := t.port.Write(message.Bytes())
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * readDelayMs)
	response := make([]byte, myrtio.MaxMessageLength)
	_, err = t.port.Read(response)
	if err != nil {
		return nil, err
	}
	responseMessage, err := myrtio.ParseMessage(response)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * commandDelayMs)
	return responseMessage, nil
}

// New creates new serial transport
func New(path string, baudRate int) (*Transport, error) {
	port, err := serial.Open(path, &serial.Mode{
		BaudRate: baudRate,
	})
	if err != nil {
		return nil, err
	}
	err = port.SetReadTimeout(time.Millisecond * readTimeoutMs)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * startDelayMs)
	return &Transport{
		port: port,
	}, nil
}
