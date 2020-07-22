package server

var kvs map[string]interface{}

func apply(request *Request) (*Reply, error) {
	var reply = &Reply{}
	switch request.CMD {
	case CMDPing:
		{
			reply.Data = "pong"
		}
	case CMDSet:
		{
			doSet(string(request.Args[1]), string(request.Args[2]))
			reply.Data = "ok"
		}
	}
	return reply, nil
}
