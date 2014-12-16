package main

import "fmt"

func deferFunc() int {

	index := 0
	fc := func() {
		fmt.Println(index, "匿名函数1")
		index++
		defer func() {
			fmt.Println(index, "匿名函数1-1")
			index++
		}()
	}

	defer func() {
		fmt.Println(index, "匿名函数2-1")
		index++
	}()

	defer fc()

	return func() int {
		fmt.Println(index, "匿名函数3")
		index++
		return index
	}()
}

func Log(title string, GetMsg func() string) {

	if true {
		fmt.Println("Log:", GetMsg())
	}
}

func main() {
	index := deferFunc()

	fmt.Println(" 作用域LAST:", index)

	msg := func() string {
		return "Do not remind me that i broke the rule"
	}

	Log("Warning", msg)
}
