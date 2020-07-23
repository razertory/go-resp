package server

const (
	CMDPing    = "ping"
	CMDCommand = "command"
	CodePong   = "PONG"
)

func doPing() *Reply {
	return &Reply{
		Type:   ReplyTypeCode,
		Code:   CodePong,
		Number: 0,
		Data:   nil,
	}
}
