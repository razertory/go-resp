package server

import "fmt"

var kvs = make(map[string]interface{})

func apply(request *Request) (*Reply, error) {
	var reply = &Reply{}
	switch request.CMD {
	case CMDPing, CMDCommand:
		return doPing(), nil
	case CMDGet:
		return doGet(string(request.Args[0])), nil
	case CMDSet:
		return doSet(string(request.Args[0]), string(request.Args[1])), nil
	default:
		return nil, fmt.Errorf("command not support")
	}
	return reply, nil
}
