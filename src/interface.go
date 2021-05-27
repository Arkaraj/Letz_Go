package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int,error)
}

// Empty interface, no methods in it
type Closer interface {}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int,error) {

	w,err := fmt.Println(string(data))

	return w,err

}

func main() {

	var w Writer = ConsoleWriter{} 
	w.Write([]byte("Hello Go!"))

	var i interface{} = string([]byte("Golang"))

	switch i.(type) {
	case int:
		fmt.Println("i is of type int")
	case string:
		fmt.Println("i is of type string")
	case bool:
		fmt.Println("i is of type boolean")
	default:
		fmt.Println("couldn't find type of i")
	}


}