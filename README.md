# MyrtIO Go library [![Quality assurance](https://github.com/MyrtIO/myrtio-go/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/MyrtIO/myrtio-go/actions/workflows/quality-assurance.yaml) [![GitHub release](https://img.shields.io/github/v/tag/MyrtIO/myrtio-go)](https://GitHub.com/MyrtIO/myrtio-go/releases/)


<img src="./docs/logo.svg" align="right" width="100" />

Library for MyrtIO protocol support in Golang.

Allows you to format and parse messages, as well as use ready-made transports.

### Transports

- Serial
- ~~UDP~~ — not ready yet

## Usage

```go
func main() {
    port, err := serial.New(serialPath, serialBaudRate)
	if err != nil {
		log.Panic(err)
	}
    // Send ping
    response, err := m.port.RunAction(&myrtio.Message{
		Feature: 0,
		Action:  0,
	})
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
	}
	defer port.Close()
}
```