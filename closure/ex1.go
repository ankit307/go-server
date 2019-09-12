package main

import "fmt"

//Closure implimentation
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	intSeq := seq()
	fmt.Println(intSeq())
	fmt.Println(intSeq())

	fmt.Println(intSeq())
}
