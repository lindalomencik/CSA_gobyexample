package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("hello world")
	fmt.Println("go" + " lang")
	fmt.Println(true && false)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)
	const n = 240000
	const h = 3e20 / n
	fmt.Println(h)
	fmt.Println(int64(h))

	fmt.Println(math.Sin(n))

}
