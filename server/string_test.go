package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_string(t *testing.T) {
	k := "key"
	v := "value"
	doSet(k, v)
	reply := doGet(k)
	assert.Equal(t, v, reply.Data)
	assert.Equal(t, "", doGet("").Data)
}
