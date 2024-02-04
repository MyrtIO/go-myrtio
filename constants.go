// Package myrtio adds support for MyrtIO APIs
package myrtio

const (
	// MaxMessageLength represents the maximum allowed length for a message.
	MaxMessageLength = 128
	// MetaHeaderPadding represents header size (header (2) + length)
	MetaHeaderPadding = 3
	// MinMessageLength represents the minimum allowed message length.
	// header + feature + action + tail
	MinMessageLength = MetaHeaderPadding + 3
)

const (
	// FirstHeaderCode represents the first header code.
	FirstHeaderCode byte = 0xFE
	// SecondHeaderCode represents the second header code.
	SecondHeaderCode byte = 0xEF
	// TailCode represents the tail code.
	TailCode byte = 0xAF
	// SuccessCode represents the success code.
	SuccessCode byte = 0xEE
	// ErrorCode represents the error code.
	ErrorCode byte = 0xFF
)
