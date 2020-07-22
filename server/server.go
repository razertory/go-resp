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
		request, err := parseRequest(conn)
		if err != nil {
			return // TODO return err code
		}
		reply, err := apply(request)
		if err != nil {
			return // TODO return err code
		}
		data := parseReply(reply)
		conn.Write(data)
	}
}
