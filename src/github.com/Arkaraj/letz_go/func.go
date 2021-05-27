package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name string
}

func swap(name *string, action *string) {

	*name, *action = *action, *name

}

// Variatic parameters
func sum(values ...int) int {

	result := 0

	for _, v := range values {
		result += v
	}

	return result

}

func check(name string) (flag bool) {

	flag = name == "Arkaraj"
	return

}

func divide(a float32, b float32) (float32, error) {

	if b == 0.0 {
		panic("Error: Tried to divide with zero")
	} 

	return a/b, nil

}

func (g *greeter) greet() {

	g.name = "TypeScript"

}

func main() {

	a := "Arkaraj"
	b := "Learning Go"

	fmt.Println(a,b)

	swap(&a,&b)

	fmt.Println(a,b)

	fmt.Println("Sum is:", sum(1,2,3,4,5,6))
	fmt.Println(check("Arkaraj"))

	res,err := divide(7.98, 2.00)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	golang := greeter{
		greeting: "Hello",
		name: "Go",
	}

	golang.greet()

	fmt.Println(golang)


}