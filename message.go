package myrtio

import (
	"errors"
)

var ErrWrongHeader = errors.New("unexpected header codes")
var ErrWrongTail = errors.New("unexpected tail code")
var ErrWrongLength = errors.New("message length is not correct")

type Message struct {
	Feature byte
	Action  byte
	Payload []byte
}

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
