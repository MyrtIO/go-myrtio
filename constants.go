package myrtio

const MaxMessageLength = 128

const (
	FirstHeaderCode  byte = 0xFE
	SecondHeaderCode byte = 0xEF
	TailCode         byte = 0xAF
	SuccessCode      byte = 0xEE
	ErrorCode        byte = 0xFF
)

const MetaHeaderPadding = 3

// header (2) + length + feature + action + tail
const MinMessageLength = 6
