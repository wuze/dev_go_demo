package main

import (
	"fmt"
	"time"
)

func producer(c chan int64, max int) {
	defer close(c)

	for i := 0; i < max; i++ {
		c <- time.Now().Unix()
	}
}

func consumer(c chan int64) {
	var v int64
	ok := true

	for ok {
		if v, ok = <-c; ok {
			fmt.Println(v)
		}
	}
}

func main() {
	//c := make(chan int64)
	//go producer(c, 100)
	//consumer(c)

	defer func() {
		fmt.Println("C")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()

	f()
}

func f() {
	fmt.Println("a")
	panic(44)
	fmt.Println("b")
	fmt.Println("f")

}
