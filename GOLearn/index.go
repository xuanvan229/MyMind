package main

import "fmt"

func swap(x , y string ) (string, string ) {
	return y, x
}
func split (sum int ) (x, y int) {
	x = sum / 2
	y = sum % 2 
	return
}
func swapnumber(x, y int ) (int, int) {
	return y, x
}
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}
func inra() string {
	return "Van Hong Xuan "
}
func main() {
	fmt.Println(inra())
	fmt.Println("hello world!")
	f := fib()
	fmt.Println(f(), f(), f(), f(), f())
	var pi float64 = 3.14159265
	fmt.Printf("%.3f ", pi)
	const n = 10000
	fmt.Println(n)
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	o, p := swapnumber(6, 3)
	fmt.Println(o,p)
	s,d := split(29)
	fmt.Println(s,d)
	a, b := swap("world", "hello")
	fmt.Println(a, b)	

	var xuan = 2
	switch xuan {
	case 1:
		fmt.Println("here is one")
	case 2:
		fmt.Println("here is two")
	case 3:
		fmt.Println("here is three")

	}
}
