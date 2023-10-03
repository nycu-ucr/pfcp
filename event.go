package pfcp

import "net"

type ReceiveEvent struct {
	Type       ReceiveEventType
	RemoteAddr *net.UDPAddr
	RcvMsg     *Message
}

const (
	ReceiveEventTypeResendRequest EventType = iota
	ReceiveEventTypeValidResponse
)
