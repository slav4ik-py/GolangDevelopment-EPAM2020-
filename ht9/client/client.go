package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3535")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		fmt.Println("Enter number or smth else: ")
		var source string
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if source == "exit" {
			os.Exit(1)
		}
		if _, err := conn.Write([]byte(source)); err != nil {
			fmt.Println(err)
			continue
		}
		buf := make([]byte, (4 * 1024))
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}

}
