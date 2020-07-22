package server

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expects = []struct {
	str     string
	request *Request
}{
	{"*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n",
		&Request{CMD: "set", Args: [][]byte{[]byte("mykey"), []byte("myvalue")}}},
}

func TestParseRequest(t *testing.T) {
	for _, expect := range expects {
		line := expect.str
		r := toReader(line)
		request, err := parseRequest(r)
		assert.NoError(t, err)
		assert.Equal(t, request, expect.request)
	}
}

func toReader(str string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(str))
}
