package server

const (
	CMDSet  = "set"
	CMDGet  = "get"
	CMDMGet = "mget"
	CMDMSet = "mset"
	CodeOk  = "OK"
)

func doSet(key string, value string) *Reply {
	applicationKV.Store(key, value)
	return &Reply{
		Type: ReplyTypeCode,
		Code: CodeOk,
	}
}

func doGet(key string) *Reply {
	reply := &Reply{
		Type: ReplyTypeBulk,
		Data: "",
	}
	v, _ := applicationKV.Load(key)
	if v != nil {
		reply.Data = v.(string)
	}
	return reply
}

func doMSet(req *Request) *Reply {
	args := req.Args
	for i := 1; i < len(args); i++ {
		applicationKV.Store(string(args[i-1]), string(args[i]))
	}
	return &Reply{
		Type: ReplyTypeCode,
		Code: CodeOk,
	}
}

func doMGet(req *Request) *Reply {
	args := req.Args
	var replyData []string
	for _, arg := range args {
		if data, _ := applicationKV.Load(string(arg)); data != nil {
			replyData = append(replyData, data.(string))
		}
	}
	return &Reply{
		Type: ReplyTypeBulk,
		Data: replyData,
	}
}
