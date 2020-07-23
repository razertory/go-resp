package server

import (
	"net"
)

func Run() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		reply, err := func() (*Reply, error) {
			request, err := parseRequest(conn)
			if err != nil {
				return nil, err
			}
			reply, err := apply(request)
			if err != nil {
				return nil, err
			}
			return reply, nil
		}()
		if err != nil {
			conn.Write([]byte(redisErr(err)))
			continue
		}
		data := parseReply(reply)
		conn.Write([]byte(data))
	}
}

func redisErr(err error) string {
	return "-ERR " + err.Error() + "\r\n"
}
