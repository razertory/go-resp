package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ping(t *testing.T) {
	request := &Request{
		CMD:  CMDPing,
		Args: nil,
	}
	reply, err := apply(request)
	assert.NoError(t, err)
	assert.Equal(t, CodePong, reply.Code)
}
