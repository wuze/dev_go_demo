package main

import (
	"fmt"
	//"os"
	//"reflect"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func fab(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	/*
		var fname string = "a.txt"

		f, _ := os.Create(fname)
		defer f.Close()
		fmt.Println(reflect.ValueOf(f).Type())

		t, _ := os.Stat(fname)
		fmt.Println(t.Size())
	*/

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	t := make(chan int, 10)
	go fab(cap(t)+2, t)

	for i := range t {
		fmt.Println(i, " ")
	}

}
