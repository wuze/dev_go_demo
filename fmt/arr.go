package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 99}

	fmt.Println(arr1, arr2)

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _, value := range arr2 {
		fmt.Println(value)
	}

	fmt.Printf("=========================\n")
	fmt.Println(reflect.TypeOf(arr1), arr1)

	fmt.Println(arr1[1:4])
	fmt.Println(arr1[2:4])

	arr3 := arr1[:]
	arr3 = append(arr3, 10, 4, 5, 6, 7)

	fmt.Println(arr3)
}
