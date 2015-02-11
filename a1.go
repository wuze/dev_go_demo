package main

import "fmt"

type Node struct {
	Name string
	Age  int64
}

type Animal interface { // {{{
	move()
}

type Human struct {
	i int
}

func (r Human) move() {
	fmt.Println("人类行走")
	r.i++
}

type Bird struct {
	i int
}

func (r *Bird) move() {
	fmt.Println("鸟类飞行")
	r.i++
}

func moveTest1(animal Animal) {
	animal.move()
}

func moveTest2(animal *Animal) {
	(*animal).move()
}

// }}}

func main() {
	var m map[string][]Node // {{{
	m = make(map[string][]Node)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%02d_key", i)
		for j := 1; j < 5; j++ {
			var tmp int64 = int64(i * j)
			nm := fmt.Sprintf("%d_key", tmp)
			nd := Node{nm, tmp}
			m[key] = append(m[key], nd)
		}
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	var t map[string]map[string]Node
	t = make(map[string]map[string]Node)

	t["test"] = make(map[string]Node)
	t["test"]["test2"] = Node{"CC", 1212}

	for _, v := range t {
		for kk, vv := range v {
			fmt.Println(kk, vv)
		}
	} // }}}

	h1 := Human{0}
	moveTest1(h1)
	moveTest1(h1)
	moveTest1(h1)
	moveTest1(h1)
	fmt.Println(h1.i)

	h2 := &Human{1}
	moveTest2(h2)
	moveTest2(h2)
	moveTest2(h2)
	moveTest2(h2)
	fmt.Println(h2.i)

	b2 := &Bird{3}
	moveTest2(b2)
	moveTest2(b2)
	moveTest2(b2)
	moveTest2(b2)
	fmt.Println(b2.i)
}
