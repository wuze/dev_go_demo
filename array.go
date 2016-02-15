package main

import (
	"fmt"
)

func inArray(list []string, key string) int {
	for i, s := range list {
		if s == key {
			return i
		}
	}

	return -1
}

func main() {
	list := []string{"string1", "string2", "string3"}

	if k := inArray(list, "string1"); k >= 0 {
		fmt.Print("Found it\n")
	} else {
		fmt.Print("Not Found")
	}
}
