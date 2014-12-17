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

	/*
		defer func() {
			fmt.Println("C")
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("d")
		}()

		f()
	*/
	var mp = map[string]int{
		"aaa": 1,
		"bbb": 2,
		"ccc": 3,
		"ddd": 4,
	}

	fmt.Println("===================\n")
	var m []string

	for k, v := range mp {
		m = append(m, k)
		fmt.Printf("KEY:%s   VALUE: %d\n", k, v)
	}

	fmt.Println("All Keys:\n")
	fmt.Println(m)
}

func f() {
	fmt.Println("a")
	panic(44)
	fmt.Println("b")
	fmt.Println("f")

}
