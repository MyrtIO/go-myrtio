// MyrtIO terminal example
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MyrtIO/myrtio-go/serial"

	"github.com/MyrtIO/myrtio-go/cmd/terminal/repl"
)

const (
	serialBaudRate = 9600
)

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
	term := repl.New(port)
	name, err := term.Name()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to " + name)
	term.Start()
}
