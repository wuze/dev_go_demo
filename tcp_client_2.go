package main

import (
	"fmt"
	//	"io"
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

func main() {

	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:9988")
		checkError(err)

		//fmt.Printf("connect success! ID: %d\n", i+1)

		var buf [512]byte

		str := fmt.Sprintf("Data From client ID: %d ", i)
		conn.Write([]byte(str))

		/*
			for {
				_, err := conn.Read(buf[0:])
				if err != nil {
					if err == io.EOF {
						break
					}
				}
			}
		*/

		conn.Read(buf[0:])
		fmt.Println(string(buf[:]))
		time.Sleep(1000 * time.Millisecond)
		conn.Close()
	}
}
