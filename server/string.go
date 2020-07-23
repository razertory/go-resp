package server

const (
	CMDSet = "set"
	CMDGet = "get"
	CodeOk = "OK"
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
