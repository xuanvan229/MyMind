package main

// three operator for import in golang
// 1. Dot operator. Some time people use following way to import package. The dot operator means you can omit the package name when you call functions inside of that package. Now fmt.Println() becomes to Printf()
// 2. Allias Operator. like that `import (f"fmt")`. It changes the name of package that we imported when we call function that belong to that package. Now fmt.Println() becomes to f.Println()
// 3. _ operator. This is operator actually means we just want to import that package ad execute its init function, and we not sure if we want to use the functions belonging that package
import (
	"./mylib"
	"errors"
	"fmt"
	"math"
)

// we cant define new type of containers of tother properties or fields in Go just like orther programming languages
// define a new type
type person struct {
	name string
	age  int
}
type Student struct {
	person
	school string
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

// method
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color                Color
}

type Color byte
type BoxList []Box // a slice of boxes

//method
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// method with a pointer receiver
func (b *Box) SetColor(c Color) {
	b.color = c
}

//method
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

// method
func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

//method
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func Goto() {
	i := 0
Here:
	fmt.Println(i)
	i++
	goto Here

}
func init() {
	fmt.Println("init 3")
}
func main() {
	// khai báo nhiều biến
	var (
		i      int     = 1
		pi     float32 = 3.14
		prefix string  = "hello"
	)
	newfile.Return_value("Hong Xuan")
	// array
	var arr [10]int
	arr[0] = 12
	arr[1] = 50
	fmt.Println("the first element  \n", arr[0])
	fmt.Println("the last element \n", arr[9])
	var x1, x2, x3 int = 1, 2, 3
	var s string = "hello"
	c := []byte(s)
	c[0] = 'c' // thay thế ký tự đầu tiên :))
	s2 := string(c)
	fmt.Println(s2)
	s1 := "world"
	s1 = "m" + s1[1:] // loại bỏ ký tự đầu tiên
	m := `Hello
		World
	`
	err := errors.New("Emit macho dwarf")
	if err != nil {
		fmt.Println(err)
	}
	// slice use when we dont know how long the array will be when we define it
	fmt.Println("===========Slice============>")
	slice := []int{1, 2, 3, 4, 5, 5}
	for k, v := range slice {
		fmt.Println("key", k)
		fmt.Println("value", v)
	}
	fmt.Println("<=======================")
	//MAP
	// use string as the key type, int as the value type, and `make` initialize it
	// var numbers map[string]int
	// another way define map
	fmt.Println("===========Map============>")
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 10
	fmt.Println("the second number is : ", numbers["two"]) // getvalues
	fmt.Println("<=======================")
	fmt.Println("===========For============>")
	sum := 0
	for index := 1; index < 10; index++ {
		sum += index
		fmt.Println("sum", sum)
	}
	fmt.Println("break")
	for index1 := 10; index1 > 0; index1-- {
		if index1 == 5 {
			break
		}
		fmt.Println(index1)
	}
	fmt.Println("continue")
	for index := 10; index > 0; index-- {
		if index == 5 {
			continue
		}
		fmt.Println(index)
	}

	fmt.Println("<=======================")
	fmt.Println(slice)
	fmt.Println(i)
	fmt.Println(pi)
	fmt.Println(prefix)
	fmt.Println(m)
	fmt.Println(s1)
	fmt.Println(s1[1:])
	fmt.Println(x1, x2, x3)
	fmt.Println("===========Struct============>")
	var tom person
	tom.name, tom.age = "Tom", 21
	xuan := person{age: 20, name: "Xuan"}
	t_bOlder, tb_diff := Older(tom, xuan)
	mark := Student{person{"Hoang", 20}, "bk DN"}
	fmt.Println(mark)
	fmt.Printf("Of %s and %s, %s is older by %d year \n", tom.name, xuan.name, t_bOlder.name, tb_diff)
	fmt.Println("<=======================")
	fmt.Println("===========Method============>")
	c1 := Circle{10}
	r1 := Rectangle{10, 5}
	fmt.Println("Area of c1 is", c1.Area())
	fmt.Println("Area of r1 is", r1.Area())
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

	// Let's paint them all black
	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
	fmt.Println("<=======================")

}
