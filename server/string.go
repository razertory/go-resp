package server

const (
	CodeOk = "OK"
)

func doSet(key string, value string) *Reply {
	kvs[key] = value
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
	v := kvs[key]
	if v != nil {
		reply.Data = v.(string)
	}
	return reply
}
