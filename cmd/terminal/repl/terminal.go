package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/MyrtIO/myrtio-go"
)

type MyrtIOTerminal struct {
	Prompt string
	port   myrtio.Transport
	reader *bufio.Reader
}

func New(port myrtio.Transport) *MyrtIOTerminal {
	return &MyrtIOTerminal{
		Prompt: "$ ",
		port:   port,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (m *MyrtIOTerminal) Name() (string, error) {
	resp, err := m.port.RunAction(&myrtio.Message{
		Feature: 0,
		Action:  1,
	})
	if err != nil {
		return "", err
	}
	return string(resp.Payload[1:]), nil
}

func (m *MyrtIOTerminal) Start() {
	for {
		fmt.Print(m.Prompt)
		input, err := m.read()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		request := &myrtio.Message{
			Feature: input[0],
			Action:  input[1],
			Payload: input[2:],
		}
		fmt.Println(">", prettyMessage(request.Bytes(), false))
		result := m.eval(input)
		fmt.Println("<", prettyMessage(result.Bytes(), true))
	}
}

func (m *MyrtIOTerminal) read() ([]byte, error) {
	input, err := m.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	symbols, err := ParsePayload(input)
	if err != nil {
		return nil, err
	}
	return symbols, nil
}

func (m *MyrtIOTerminal) eval(command []byte) *myrtio.Message {
	response, err := m.port.RunAction(&myrtio.Message{
		Feature: command[0],
		Action:  command[1],
		Payload: command[2:],
	})
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
	}
	return response
}
