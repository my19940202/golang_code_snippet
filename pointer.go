package main

import "fmt"

type MyInt struct {
    n int
    name string
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{100, "fuck me"}
    mi.Decrease()
    mi.Decrease()
    mi.Decrease()
	fmt.Println(mi.n, mi.name)
    mi.Increase()
    mi.Increase()
    mi.Increase()
	fmt.Println(mi.n, mi.name)
}