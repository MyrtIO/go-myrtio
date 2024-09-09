// MyrtIO terminal example
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/MyrtIO/myrtio-go"
	"github.com/MyrtIO/myrtio-go/serial"
)

const (
	serialBaudRate = 28800
)

func randomByte() byte {
	return byte(rand.Int31n(256))
}

func main() {
	paths, err := serial.Discover()
	if err != nil {
		log.Panic(err)
	}
	if len(paths) == 0 {
		fmt.Println("Serial devices is not found")
		os.Exit(1)
	}
	port, err := serial.New(paths[0], serialBaudRate)
	if err != nil {
		log.Panic(err)
	}
	defer port.Close()

	writeMsg := myrtio.Message{
		Feature: 2,
		Action:  0,
		Payload: make([]byte, 3),
	}
	go attack("write", port, func() *myrtio.Message {
		writeMsg.Payload = make([]byte, 3)
		for i := 0; i < 3; i++ {
			writeMsg.Payload[i] = randomByte()
		}
		return &writeMsg
	})

	readMsg := myrtio.Message{
		Feature: 2,
		Action:  2,
	}
	go attack("read", port, func() *myrtio.Message {
		return &readMsg
	})

	<-make(chan struct{})
}

type MessageGenerator func() *myrtio.Message

func attack(ch string, port myrtio.Transport, generate MessageGenerator) {
	for {
		msg, err := port.RunAction(generate())
		if err != nil {
			log.Printf("%s: %v\n", ch, err)
		} else {
			log.Printf("%s: %+v\n", ch, msg.SkipStatus())
		}
	}
}
