package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Formatter interface {
	Format(f State, c rune)
}

type State interface {
	Write(b []byte) (ret int, err error)
	Width() (wid int, ok bool)
	Presision() (prec int, ok bool)
	Flag(c int) bool
}

type Stringer interface {
	String() string
}

type GoStringer interface {
	GoString() string
}

type Ustr string

func (us Ustr) String() string {
	return string(us) + " Go Format"
}

func (us Ustr) GoString() string {
	return string(us) + "Go 格式化"
}

func (us Ustr) Format(f fmt.State, c rune) {

	switch c {
	case 'm', 'M':
		f.Write([]byte(us + "\n 扩展标记:["))

		if f.Flag('-') {
			f.Write([]byte(" -"))
		}

		if f.Flag('+') {
			f.Write([]byte(" +"))
		}

		if f.Flag('#') {
			f.Write([]byte(" #"))
		}

		if f.Flag(' ') {
			f.Write([]byte(" space"))
		}

		if f.Flag('0') {
			f.Write([]byte(" 0"))
		}

		f.Write([]byte(" ]\n"))

		if w, wok := f.Width(); wok {
			f.Write([]byte("Width:" + fmt.Sprint(w) + "\n"))
		}

		if p, pok := f.Width(); pok {
			f.Write([]byte("Presision:" + fmt.Sprint(p) + "\n"))
		}

	case 'v':
		if f.Flag('#') {

			f.Write([]byte("TESTESTEST\n"))
			f.Write([]byte(us.GoString()))
		} else {

			f.Write([]byte(us.String()))
		}

	default:
		f.Write([]byte(us.String()))
	}
}

func main() {

	us := Ustr("Hello World")
	fmt.Printf("% 0-+#8.5m\n", us)
	fmt.Printf("%#v\n", us)

	s := strings.NewReader("我是 Golang  我已经 4 岁了")
	var name string
	var age int

	fmt.Printf("------------------------------\n")
	fmt.Fscanf(s, "我是 %s  我已经 %d 岁了", &name, &age)

	fmt.Printf("%s %d\n", name, age)

	var a, b []byte

	a = []byte("helloBaby")
	b = []byte("heloBaby2")

	if bytes.Compare(a, b) < 0 {
		fmt.Printf("a is less than b\n")
	} else {
		fmt.Printf("a is bigger than b\n")
	}

	str := "你好世界我们都是"
	r := []rune(str)

	length := len(r)
	r[6] = '中'
	r[7] = '国'

	fmt.Printf("------------------------\n")
	fmt.Printf("%d %v\n", length, string(r))
	fmt.Printf("------------------------\n")

	fmt.Printf("%v\n", len(str))
	for i, v := range str {
		fmt.Printf("%2v = %c\n", i, v)
	}

}
