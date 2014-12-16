package main

import (
	//	"bytes"
	"fmt"
)

var chanstr string
var ch1 = make(chan int, 10)
var ch2 = make(chan int)

func sum(a []int, c chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}
	c <- sum
}

// 有缓冲buffer
func buffered() {
	chanstr = "buffered"
	ch1 <- 0
}

// 无缓冲buffer
func no_buffer() {
	chanstr = "no_buffer"
	<-ch2
}

func main() {
	/*// {{{
	s := []byte("string")
	str := "string_hello"
	buf := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("nihaoshijie"))
	buf3 := bytes.NewBuffer([]byte{'n', 'h', 'c', 'd', 'd', 'm'})

	fmt.Println(buf, buf2, buf3)
	buf.Write(s)
	fmt.Println(buf, buf2, buf3)
	buf.WriteString(str)
	fmt.Println(buf, buf2, buf3)

	var c rune = '好'

	buf.WriteRune(c)
	fmt.Println(buf, buf2, buf3)
	*/ // }}}
	fmt.Printf("---------------------------\n")
	go buffered()
	<-ch1
	fmt.Println(chanstr)

	go no_buffer()
	ch2 <- 0
	fmt.Println(chanstr)
	fmt.Printf("---------------------------\n")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)

	go sum(a[1:4], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Printf("Sum:%d  %d\n", x, y)
}
