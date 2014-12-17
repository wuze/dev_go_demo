package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error", err.Error())
		os.Exit(1)
	}

}

func handConnection(conn net.Conn, i int) {

	fmt.Println("Conntion success! ID", i)
	i += 1
	time.Sleep(1 * time.Millisecond)
	_, _ = conn.Write([]byte("Msg from server \r\n"))
	time.Sleep(1 * time.Millisecond)
	conn.Close()

}

func main() {
	i := 0
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)

	fmt.Println("服务器处于监听状态....\n")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal Error", err.Error())
			continue
		}
		i += 1

		// 来一个就单独开一个携程
		go handConnection(conn, i)
	}

}
