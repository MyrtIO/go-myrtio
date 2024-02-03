package serial

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/MyrtIO/myrtio-go"

	"go.bug.st/serial"
)

// MyrtIO typical serial timings
const (
	startDelayMs   = 100
	readDelayMs    = 50
	commandDelayMs = 10
)

type Transport struct {
	port serial.Port
	mu   sync.Mutex
}

func (t *Transport) Close() error {
	return t.port.Close()
}

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
		return nil, errors.Join(
			err,
			fmt.Errorf("%v", response),
		)
	}
	time.Sleep(time.Millisecond * commandDelayMs)
	return responseMessage, nil
}

func New(path string, baudRate int) (*Transport, error) {
	port, err := serial.Open(path, &serial.Mode{
		BaudRate: baudRate,
	})
	port.SetReadTimeout(time.Millisecond * 500)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * startDelayMs)
	return &Transport{
		port: port,
	}, nil
}
