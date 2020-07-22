package server

import "fmt"

var kvs = make(map[string]interface{})

func apply(request *Request) (*Reply, error) {
	var reply = &Reply{}
	fmt.Printf("request %+v \n", request)
	switch request.CMD {
	case CMDPing:
		{
			reply.Type = ReplyTypeCode
			reply.Code = "pong"
		}
	case CMDGet:
		{
			data := doGet(string(request.Args[0]))
			reply.Type = ReplyTypeBulk
			reply.Data = data
		}
	case CMDSet:
		{
			doSet(string(request.Args[0]), string(request.Args[1]))
			reply.Type = ReplyTypeCode
			reply.Code = "ok"
		}
	}
	return reply, nil
}
