package server

func doSet(key string, value string) {
	kvs[key] = value
}

func doGet(key string) string {
	v := kvs[key]
	if v == nil {
		return ""
	}
	return v.(string)
}
