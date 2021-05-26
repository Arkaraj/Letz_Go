package main

import (
	"fmt"
)

type myStruct struct {
	foo int
	bar string
}

func main()  {

	a := 42
	// b is holding the memory address of a
	var b *int = &a 

	fmt.Println(a,b)
	// Dereferencing b
	fmt.Println(a,*b)

	a = 17
	fmt.Println(a,*b)

	*b = 67 // will change a as well 

	fmt.Println(a,*b)

	arr := [3]int {1,2,3}
	arrb := &arr[0]
	arrc := &arr[1]

	fmt.Printf("%v, %p, %p \n", arr, arrb, arrc)

	var ms *myStruct

	ms = new(myStruct)
	(*ms).foo = 42 // same as ms.foo
	(*ms).bar = "life_Equation"
	fmt.Println(ms.foo)

}