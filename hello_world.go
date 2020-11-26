package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("hello world")
	fmt.Println("go" + " lang")
)

	fmt.Println(s)
	const n = 240000
	const h = 3e20 / n
	fmt.Println(h)
	fmt.Println(int64(h))

	fmt.Println(math.Sin(n))

}
