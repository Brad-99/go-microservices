package main

import (
	"fmt"
	"unsafe"

	geo "github.com/Brad-99/go-microservices/geometry"

	"rsc.io/quote"
)

func rectProps(length, width float64) (area1, perimeter float64) {
	area1 = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	x := 10
	var y, z = 2, 3
	var name string
	isWorking := false

	name = "Puppy"
	fmt.Println("Test")
	fmt.Println(quote.Go())
	fmt.Println(x, y, z, name, isWorking)
	fmt.Println("Type of name %T and size is %d\n", name, unsafe.Sizeof(name))

	a, p := rectProps(1, 2)
	fmt.Println("Area is %f and perimeter is %f", a, p)

	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
	fmt.Println(daysOfTheMonth)

	area := geo.Area(1, 2)
	diag := geo.Diagonal(1, 2)
	fmt.Println(area, diag)

}
