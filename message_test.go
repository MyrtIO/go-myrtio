package myrtio_test

import (
	"bytes"
	"testing"

	"github.com/MyrtIO/myrtio-go"
)

func TestMessageBytes(t *testing.T) {
	payload := []byte{0x01, 0x02, 0x03}
	message := &myrtio.Message{
		Feature: 0x10,
		Action:  0x20,
		Payload: payload,
	}

	expected := []byte{
		myrtio.FirstHeaderCode,
		myrtio.SecondHeaderCode,
		0x05,
		0x10,
		0x20,
		0x01,
		0x02,
		0x03,
		myrtio.TailCode,
	}
	result := message.Bytes()

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestMessageSuccess(t *testing.T) {
	payloadSuccess := []byte{myrtio.SuccessCode}
	payloadFailure := []byte{myrtio.ErrorCode}

	messageSuccess := &myrtio.Message{Payload: payloadSuccess}
	messageFailure := &myrtio.Message{Payload: payloadFailure}

	if !messageSuccess.Success() {
		t.Errorf("Expected success, but got failure")
	}

	if messageFailure.Success() {
		t.Errorf("Expected failure, but got success")
	}
}

func TestParseMessage(t *testing.T) {
	validMessage := []byte{
		myrtio.FirstHeaderCode,
		myrtio.SecondHeaderCode,
		0x05,
		0x10,
		0x20,
		0x01,
		0x02,
		0x03,
		myrtio.TailCode,
	}
	invalidHeader := []byte{
		0xAA,
		0xBB,
		0x04,
		0x10,
		0x20,
		0x01,
		0x02,
		0x03,
		myrtio.TailCode,
	}
	invalidLength := []byte{myrtio.FirstHeaderCode, myrtio.SecondHeaderCode, 0x01, 0x10, myrtio.TailCode}
	invalidTail := []byte{myrtio.FirstHeaderCode, myrtio.SecondHeaderCode, 0x04, 0x10, 0x20, 0x01, 0x02, 0xBB}

	validParsed, err := myrtio.ParseMessage(validMessage)
	invalidHeaderParsed, errHeader := myrtio.ParseMessage(invalidHeader)
	invalidLengthParsed, errLength := myrtio.ParseMessage(invalidLength)
	invalidTailParsed, errTail := myrtio.ParseMessage(invalidTail)

	if err != nil || validParsed == nil || !bytes.Equal(validParsed.Bytes(), validMessage) {
		t.Errorf("Expected valid message, but got error or nil. err: %v", err)
	}

	if errHeader != myrtio.ErrWrongHeader || invalidHeaderParsed != nil {
		t.Errorf("Expected invalid header error, but got different result")
	}

	if errLength != myrtio.ErrWrongLength || invalidLengthParsed != nil {
		t.Errorf("Expected invalid length error, but got different result")
	}

	if errTail != myrtio.ErrWrongTail || invalidTailParsed != nil {
		t.Errorf("Expected invalid tail error, but got different result")
	}
}
