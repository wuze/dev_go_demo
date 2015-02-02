package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	fmt.Println("1+1=", 1+1)

	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 20
	m["k3"] = 30
	m["k4"] = 40
	m["k5"] = 50

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "k2")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)

	sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("\n----------------------------\n")

	seq := intSeq()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	fmt.Println("\n----------------------------\n")

	i := 1

	zeroptr(&i)
	fmt.Println("*ptr: ", i)
}
