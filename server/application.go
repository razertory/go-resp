package server

import (
	"fmt"
	"sync"
)

var applicationKV = sync.Map{}

func apply(request *Request) (*Reply, error) {
	var reply = &Reply{}
	fmt.Printf("request %+v \n", request)
	switch request.CMD {
	case CMDPing, CMDCommand:
		return doPing(), nil
	case CMDGet:
		return doGet(string(request.Args[0])), nil
	case CMDSet:
		return doSet(string(request.Args[0]), string(request.Args[1])), nil
	case CMDMGet:
		return doMGet(request), nil
	case CMDMSet:
		return doMSet(request), nil
	default:
		return nil, fmt.Errorf("command not support")
	}
	return reply, nil
}
