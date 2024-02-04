package myrtio

import (
	"errors"
)

// ErrWrongHeader is returned when given message contains wrong header codes
var ErrWrongHeader = errors.New("unexpected header codes")

// ErrWrongTail is returned when given message contains wrong tail code
var ErrWrongTail = errors.New("unexpected tail code")

// ErrWrongLength is returned when given message has incorrect length
var ErrWrongLength = errors.New("message length is not correct")

// Message represent MyrtIO message
type Message struct {
	Feature byte
	Action  byte
	Payload []byte
}

// Bytes returns message as byte slice
func (m *Message) Bytes() []byte {
	var payloadLength = 0
	if m.Payload != nil {
		payloadLength = len(m.Payload)
	}
	result := make([]byte, 0, payloadLength+MetaHeaderPadding+1)
	result = append(result, FirstHeaderCode, SecondHeaderCode)
	result = append(result, byte(payloadLength+2))
	result = append(result, m.Feature, m.Action)
	if m.Payload != nil && payloadLength > 0 {
		result = append(result, m.Payload...)
	}
	result = append(result, TailCode)
	return result
}

// Success checks status value and returns it as bool
func (m *Message) Success() bool {
	return m.Payload[0] == SuccessCode
}

// SkipStatus returns payload without first status value
func (m *Message) SkipStatus() []byte {
	return m.Payload[1:]
}

// ParseMessage parses message from raw bytes slice
func ParseMessage(message []byte) (*Message, error) {
	if len(message) < MinMessageLength || len(message) > MaxMessageLength {
		return nil, ErrWrongLength
	}
	if message[0] != FirstHeaderCode || message[1] != SecondHeaderCode {
		return nil, ErrWrongHeader
	}
	length := message[2]
	if len(message) < int(length)+MetaHeaderPadding {
		return nil, ErrWrongLength
	}
	if message[length+MetaHeaderPadding] != TailCode {
		return nil, ErrWrongTail
	}
	return &Message{
		Feature: message[MetaHeaderPadding],
		Action:  message[MetaHeaderPadding+1],
		Payload: message[MetaHeaderPadding+2 : MetaHeaderPadding+length],
	}, nil
}
