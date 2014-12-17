package main

import (
	"bytes"
	"container/list"
	"fmt"
	"net"
	//	"os"
	//	"time"
)

type Client struct {
	Name       string
	Incoming   chan string
	Outgoint   chan string
	Conn       net.Conn
	Quit       chan bool
	ClientList *list.List
}

func (c *Client) Read(buffer []byte) bool {
	bytesRead, err := c.Conn.Read(buffer)

	if err != nil {
		fmt.Println("Read Error")
		c.Close()
		Log(err, "\n")
		return false
	}

	Log("Read: ", bytesRead, " bytes\n")
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
			Log("RemoveMe", c.Name, "\n")
			c.ClientList.Remove(entry)
		}
	}
}

func Log(v ...interface{}) {
	fmt.Print(v...)
}

func (c *Client) Equal(other *Client) bool {
	if bytes.Equal([]byte(c.Name), []byte(other.Name)) {
		if c.Conn == other.Conn {
			return true
		}
	}

	return false
}

func IOHandler(Incoming <-chan string, clientList *list.List) {
	for {
		Log("(IOHandler)Waiting for input...\n")
		input := <-Incoming
		Log("Handling ", input, "\n")

		for e := clientList.Front(); e != nil; e = e.Next() {
			client := e.Value.(Client)
			client.Incoming <- input
		}
	}
}

func ClientReader(client *Client) {
	buffer := make([]byte, 2048)

	for client.Read(buffer) {
		if bytes.Equal(buffer, []byte("/quit")) {
			client.Close()
			break
		}

		Log("ClientRead received", client.Name, ">", string(buffer), "\n")
		send := client.Name + ">" + string(buffer)

		client.Outgoint <- send

		//memset()
		for i := 0; i < 2048; i++ {
			buffer[i] = 0x00
		}
	}

	client.Outgoint <- client.Name + " has left chat"
	Log("ClientReader stopped for ", client.Name)
}

func ClientSender(client *Client) {
	for {
		select {
		case buffer := <-client.Incoming:
			Log("ClientSender sending ", string(buffer), " to ", client.Name, "\n")
			count := 0
			for i := 0; i < len(buffer); i++ {
				if buffer[i] == 0x00 {
					break
				}
				count++
			}

			Log("Send Size: ", count, "\n")
			client.Conn.Write([]byte(buffer)[0:count])

		case <-client.Quit:
			Log("Client ", client.Name, " quiting\n")
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
	//	newClient := &Client(name, make(chan string), ch, conn, make(chan bool), clientList)

	go ClientSender(newClient)
	go ClientReader(newClient)

	clientList.PushBack(*newClient)
	ch <- string(name + " has Joined the chat")
}

func main() {
	Log("Hello Server\n")
	clientList := list.New()
	in := make(chan string)

	go IOHandler(in, clientList)

	service := ":9988"
	tcpAddr, error := net.ResolveTCPAddr("tcp", service)

	if error != nil {
		Log("Error: Could not resolve address\n")
	} else {
		netListen, error := net.Listen(tcpAddr.Network(), tcpAddr.String())
		if error != nil {
			Log(error, "\n")
		} else {
			defer netListen.Close()

			for {
				Log("Waiting for client...\n")
				connection, error := netListen.Accept()

				if error != nil {
					Log("Client error: ", error, "\n")
				} else {
					go ClientHandler(connection, in, clientList)
				}
			}
		}
	}
}
