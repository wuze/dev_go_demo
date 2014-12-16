package main

import (
	"bufio"
	"fmt"
	"os"
)

func menu() {
	// {{{
	fmt.Printf("=%-10s\n", "1.	ID")
	fmt.Printf("=%-10s\n", "2.  Menu")
	fmt.Printf("=%-10s\n", "3.  other")
	fmt.Printf("=%-10s\n", "4.  exit")

} // }}}

func main() {

	var i, j, k, l int = 0, 0, 0, 0
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	for _, rune := range input {
		switch {
		case (rune >= 'A' && rune <= 'Z'):
			i++
		case (rune >= 'a' && rune <= 'z'):
			j++
		case (rune == ' ' || rune == '\n'):
			k++
		case (rune >= '0' && rune <= '9'):
			l++
		default:
			i++
		}
	}

	fmt.Printf("BAl:%d SAl:%d SP:%d NUM:%d\n", i, j, k, l)
	fmt.Printf("-------\n")
	//menu()

	var m int

	_, err := fmt.Scanf("%d", &m)

	if err != nil {
		fmt.Printf("输入有异常")
		return
	}

	for {

		fmt.Printf("%d\n", m)

		switch m {
		case 1:
		case 2:
		case 3:
		default:
			fmt.Printf("ccc")
		}

		break
	}
}
