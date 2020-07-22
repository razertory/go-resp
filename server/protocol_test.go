package server

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var requestExpects = []struct {
	str     string
	request *Request
}{
	{"ping\r\n",
		&Request{CMD: "ping"},
	},
	{"*2\r\n$3\r\nGET\r\n$4\r\nname\r\n",
		&Request{
			CMD:  "get",
			Args: [][]byte{[]byte("name")},
		}},
	{"*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n",
		&Request{CMD: CMDSet, Args: [][]byte{[]byte("mykey"), []byte("myvalue")}}},
}

func TestParseRequest(t *testing.T) {
	for _, expect := range requestExpects {
		line := expect.str
		r := toReader(line)
		request, err := parseRequest(r)
		assert.NoError(t, err)
		assert.Equal(t, expect.request, request)
	}
}

var replyExpects = []struct {
	reply *Reply
	str   string
}{
	{
		&Reply{
			Type:   ReplyTypeCode,
			Code:   "PONG",
			Number: 0,
			Data:   nil,
		},
		"+PONG\r\n",
	},
	{
		&Reply{
			Type:   ReplyTypeCode,
			Code:   "OK",
			Number: 0,
			Data:   nil,
		},
		"+OK\r\n",
	},
	{
		&Reply{
			Type:   ReplyTypeNumber,
			Code:   "",
			Number: 12,
			Data:   nil,
		},
		":12\r\n",
	},
	{
		&Reply{
			Type:   ReplyTypeBulk,
			Code:   "",
			Number: 0,
			Data:   12,
		},
		":12\r\n",
	},
	{
		&Reply{
			Type:   ReplyTypeBulk,
			Code:   "",
			Number: 0,
			Data:   "love",
		},
		"$4\r\nlove\r\n",
	},
}

func TestParseReply(t *testing.T) {
	for _, expect := range replyExpects {
		expectStr := expect.str
		reply := expect.reply
		actual := parseReply(reply)
		assert.Equal(t, expectStr, actual)
	}
}

func toReader(str string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(str))
}
