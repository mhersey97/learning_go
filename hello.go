package main

import (
	"fmt"
	"strings"
)

type building struct {
	walls  int
	floors int
	name   string
}

func main() {
	fmt.Println("Hello, World!")

	a := []float32{0.1, 0.2, 0.3}
	fmt.Println(a)

	a = append(a, 69.420)
	fmt.Println(a[3])

	b := make(map[float32]string)
	b[0.1] = "a"
	b[2.2] = "b"
	b[3.3] = "c"
	delete(b, 2.2)
	fmt.Println(b)
	fmt.Println(b[a[0]])

	for i := 0; i < 5; i++ {
		fmt.Print(i, "\n")
	}

	fmt.Println("sum is:", add(1, 2))

	fmt.Println(upperAndLen("guhguhguh doosh doosh SPLOOSH!"))

	skyscraper := building{name: "burj khalif", walls: 4, floors: 105}
	fmt.Println(skyscraper.name)
	fmt.Println(&skyscraper.name)

	foo()
}

func add(a int, b int) int {
	return a + b
}

func upperAndLen(c string) (string, int) {
	var l int
	for range c {
		l++
	}
	return strings.ToUpper(c), l
}
