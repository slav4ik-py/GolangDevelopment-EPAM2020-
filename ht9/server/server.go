package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, (1024 * 400))
		n, err := conn.Read(buf)
		if err != nil {
			conn.Write([]byte("ERROR!"))
			conn.Close()
			break
		}
		req := string(buf[:n])
		fmt.Println(req)
		if num, err := strconv.Atoi(req); err == nil {
			conn.Write([]byte(strconv.Itoa(num * 2)))
		} else {
			conn.Write([]byte(strings.ToUpper(req)))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":3535")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Write([]byte("ERROR!"))
			conn.Close()
			continue
		}
		go HandleConnection(conn)
	}
}
