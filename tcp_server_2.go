package main

import (
	//	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"net"
	//"os"
	//"time"
)

type Client struct {
	Name       string
	Incoming   chan string
	Outgoing   chan string
	Conn       net.Conn
	Quit       chan bool
	ClientList *list.List
}

func (c *Client) Read(buffer []byte) bool {

	bytesRead, err := c.Conn.Read(buffer)

	if bytesRead == 0 {
		c.Close()
		return false
	}

	if err != nil {
		fmt.Println("3. Read Error: ", err)
		c.Close()
		return false
	}

	Log("\n3. Read: ", bytesRead, " bytes\n")
	return true
}

func (c *Client) Close() {
	c.Quit <- true
	c.Conn.Close()
	c.RemoveMe()
}

func (c *Client) RemoveMe() {

	for entry := c.ClientList.Front(); entry != nil; entry = entry.Next() {
		client := entry.Value.(Client)
		if c.Equal(&client) {
			Log("4. RemoveMe", c.Name, "\n")
			c.ClientList.Remove(entry)
		}
	}

}

func Log(v ...interface{}) {
	fmt.Print(v...)
}

// 判断是否为同一个客户端
func (c *Client) Equal(other *Client) bool {
	if bytes.Equal([]byte(c.Name), []byte(other.Name)) {
		if c.Conn == other.Conn {
			return true
		}
	}

	return false
}

// 读取客户端发来的消息
func ClientReader(client *Client) {
	buffer := make([]byte, 2048)

	for client.Read(buffer) {
		if bytes.Equal(buffer, []byte("/quit")) {
			client.Close()
			break
		}

		Log("ClientRead RCV:", client.Name, " > ", string(buffer))
		send := client.Name + ">" + string(buffer)

		client.Outgoing <- send

		//memset()
		for i := 0; i < 2048; i++ {
			buffer[i] = 0x00
		}
	}

	client.Outgoing <- client.Name + " has left chat"
	Log("ClientReader stopped For:  ", client.Name)
}

// 发送消息给客户端
func ClientSender(client *Client) {

	for {
		select {
		case buffer := <-client.Incoming:
			Log("2. ClientSender SND: ", string(buffer), "\n")
			count := 0

			for i := 0; i < len(buffer); i++ {
				if buffer[i] == 0x00 {
					break
				}
				count++
			}

			//Log("XSend Size: ", count, "\n")
			client.Conn.Write([]byte(buffer)[0:count])

		case <-client.Quit:
			Log("Client  Name; ", client.Name, " quiting\n")
			client.Conn.Close()
			break
		}
	}
}

func ClientHandler(conn net.Conn, ch chan string, clientList *list.List) {
	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)

	if err != nil {
		Log("Client connection error: ", err, "\n")
	}

	name := string(buffer[0:bytesRead])

	newClient := &Client{name, make(chan string), ch, conn, make(chan bool), clientList}

	go ClientSender(newClient)
	go ClientReader(newClient)

	clientList.PushBack(*newClient)
	//	fmt.Printf("当前用户数为: %d 人\n", clientList.Len())

	ch <- string(name + " has Joined the chat")
}

func IOHandler(Incoming <-chan string, clientList *list.List) {

	for {
		Log("1. (IOHandler)Waiting for input...\n")
		// 会阻塞在此处
		input := <-Incoming

		Log("Handling:", input, "\r\n\r\n")

		for e := clientList.Front(); e != nil; e = e.Next() {
			client := e.Value.(Client)
			client.Incoming <- input
		}
	}
}

func main() {

	Log("Hello Server\n")

	clientList := list.New()
	in := make(chan string)

	// 没有链接的时候一直阻塞着 等待ClientHandler 信号
	go IOHandler(in, clientList)

	service := ":9988"
	tcpAddr, error := net.ResolveTCPAddr("tcp", service)

	if error != nil {

		Log("Error: Could not resolve address\n")

	} else {

		netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())

		if err != nil {
			Log(err, "\n")
		} else {

			defer netListen.Close()

			for {
				Log("Waiting for client...\n")
				connection, error := netListen.Accept()

				if error != nil {
					Log("Client error: ", error, "\n")
				} else {
					// 这里往channel 里面塞数据
					go ClientHandler(connection, in, clientList)
				}

			}
		}

	}
}
