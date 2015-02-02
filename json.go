package main

import (
	"encoding/json"
	"fmt"
	//	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int
	Fruits []string
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	dat := map[string]int{"apple": 5, "peach": 10, "pear": 10}
	datD, _ := json.Marshal(dat)
	var vv map[string]interface{}

	if err := json.Unmarshal(datD, &vv); err != nil {
		panic(err)
	}

	fmt.Println(vv)

	for k, v := range vv {
		fmt.Println(k, v)
	}

	res1D := &Response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1, _ := json.Marshal(res1D)

	fmt.Println(string(res1))

	res := &Response2{}

	json.Unmarshal(res1, &res)

	fmt.Println(res)
	fmt.Println("=---------------------\n")
	fmt.Println("Page:", res.Page)

	fmt.Println("Fruits: ", res.Fruits[0], res.Fruits[1], res.Fruits[2])

}
