// MyrtIO terminal example
package main

import (
	"fmt"
	"log"

	"github.com/MyrtIO/myrtio-go/serial"

	"github.com/MyrtIO/myrtio-go/cmd/terminal/repl"
)

const (
	serialPath     = "/dev/cu.wchusbserial14320"
	serialBaudRate = 9600
)

func main() {
	port, err := serial.New(serialPath, serialBaudRate)
	if err != nil {
		log.Panic(err)
	}
	defer port.Close()
	term := repl.New(port)
	name, err := term.Name()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to " + name)
	term.Start()
}
