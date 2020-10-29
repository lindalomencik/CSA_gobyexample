package main
import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13}
	var s []int = primes[1:4]
	fmt.Println(s)
	sl(6)

	p := make([]string, 4)
	p[0] = "l"
	p[1] = "i"
	p[2] = "n"
	p = append(p, "d")
	p = append(p, "a")
	fmt.Println(p)

	c := make([]string, len(p))
	copy(c,p)
	fmt.Println(c)

	l := p[2:4]
	fmt.Println(l)
	l = p[:4]
	fmt.Println(l)
	l = l[:2]
	fmt.Println(l)

	t := []string{"a","b","c"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j< innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	names := [4]string{"John","Paul","George","Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)
	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)
}

func sl(x int) []string{
	s := make([]string, x)
	fmt.Println(s)
	return s
}

