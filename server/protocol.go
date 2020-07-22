package server

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	// CMD
	CMDPing     = "ping"
	CMDSet      = "set"
	ProtoPrefix = '*'
)

type Request struct {
	CMD  string
	Args [][]byte
}

type Reply struct {
	Data interface{}
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
		return nil, malformed("$<argumentLength>", line)
	}
	var argSize int
	if _, err := fmt.Sscanf(line, "$%d\r", &argSize); err != nil {
		return nil, malformed("$<argumentSize>", line)
	}
	data, err := ioutil.ReadAll(io.LimitReader(r, int64(argSize)))
	if err != nil {
		return nil, err
	}
	if len(data) != argSize {
		return nil, malformedLength(argSize, len(data))
	}
	if b, err := r.ReadByte(); err != nil || b != '\r' {
		return nil, malformedMissingCRLF()
	}
	if b, err := r.ReadByte(); err != nil || b != '\n' {
		return nil, malformedMissingCRLF()
	}
	return data, nil
}

func malformed(expected string, got string) error {
	return fmt.Errorf("Mailformed request:'%str does not match %str\\r\\n'", got, expected)
}

func malformedLength(expected int, got int) error {
	return fmt.Errorf(
		"Mailformed request: argument length '%d does not match %d\\r\\n'",
		got, expected)
}

func malformedMissingCRLF() error {
	return fmt.Errorf("Mailformed request: line should end with \\r\\n")
}

func parseReply(reply *Reply) []byte {
	return []byte("$4\r\nPONG\r\n")
}
