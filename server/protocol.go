package server

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ProtoPrefix     = '*'
	ReplyTypeCode   = "code"
	ReplyTypeNumber = "number"
	ReplyTypeBulk   = "bulk"
)

type Request struct {
	CMD  string
	Args [][]byte
}

type Reply struct {
	Type   string
	Code   string
	Number int
	Data   interface{}
}

func parseRequest(conn io.ReadCloser) (*Request, error) {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if line[0] == ProtoPrefix {
		return readProto(line, r)
	}
	fields := strings.Split(strings.Trim(line, "\r\n"), " ")
	var args [][]byte
	if len(fields) > 1 {
		for _, arg := range fields[1:] {
			args = append(args, []byte(arg))
		}
	}
	return &Request{
		CMD:  strings.ToLower(fields[0]),
		Args: args,
	}, nil
}

func readProto(line string, r *bufio.Reader) (*Request, error) {
	var argsCount int
	if _, err := fmt.Sscanf(line, "*%d\r", &argsCount); err != nil {
		return nil, err
	}
	firstArg, err := readArgument(r)
	if err != nil {
		return nil, err
	}
	args := make([][]byte, argsCount-1)
	for i := 0; i < argsCount-1; i += 1 {
		if args[i], err = readArgument(r); err != nil {
			return nil, err
		}
	}
	return &Request{
		CMD:  strings.ToLower(string(firstArg)),
		Args: args,
	}, nil
}

func readArgument(r *bufio.Reader) ([]byte, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("length err %s" + line)
	}
	var argSize int
	if _, err := fmt.Sscanf(line, "$%d\r", &argSize); err != nil {
		return nil, fmt.Errorf("length err %s", line)
	}
	data, err := ioutil.ReadAll(io.LimitReader(r, int64(argSize)))
	if err != nil {
		return nil, err
	}
	if len(data) != argSize {
		return nil, fmt.Errorf("length err %d", len(data))
	}
	if b, err := r.ReadByte(); err != nil || b != '\r' {
		return nil, fmt.Errorf("missing CRLF")
	}
	if b, err := r.ReadByte(); err != nil || b != '\n' {
		return nil, fmt.Errorf("missing CRLF")
	}
	return data, nil
}

func parseReply(reply *Reply) string {
	var str string
	switch reply.Type {
	case ReplyTypeCode:
		str = "+" + reply.Code
	case ReplyTypeNumber:
		str = ":" + strconv.Itoa(reply.Number)
	case ReplyTypeBulk:
		str = bulkReply(reply.Data)
	}
	str += "\r\n"
	return str
}

func bulkReply(data interface{}) string {
	switch v := data.(type) {
	case string:
		{
			l := len(v)
			if l == 0 {
				return "$-1"
			}
			return "$" + strconv.Itoa(l) + "\r\n" + v
		}
	case []byte:
		{
			l := len(v)
			if l == 0 {
				return "$-1"
			}
			return "$" + strconv.Itoa(l) + "\r\n" + string(v)
		}
	case []string:
		{
			l := len(v)
			if l == 0 {
				return "$-1"
			}
			var str = "*" + strconv.Itoa(l) + "\r\n"
			for i := 0; i < len(v)-1; i++ {
				s := v[i]
				str += "$" + strconv.Itoa(len(s)) + "\r\n" + s + "\r\n"
			}
			last := v[len(v)-1]
			str += "$" + strconv.Itoa(len(last)) + "\r\n" + last
			return str
		}
	case int:
		{
			return ":" + strconv.Itoa(v)
		}
	}
	return ""
}

func redisErr(err error) string {
	return "-ERR " + err.Error() + "\r\n"
}
