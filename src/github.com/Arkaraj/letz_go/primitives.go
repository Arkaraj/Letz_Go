package main

import "fmt"

func main()  {

	var j float64 = 7.4
	var cmp complex64 = complex(7,4)

	// Strings are immutable
	s1 := "This is a String"
	s2 := "This String2"
	// Conv string to byte by []byte

	b := s1 + " " + s2

	// string - utf8
	// rune - utf32, rune is a int32...
	var r rune = 'a' // 72

	fmt.Printf("%v, %T\n", j , j)
	fmt.Printf("%v, %T\n", cmp , cmp)
	fmt.Printf("Real part: %v, %T\n", real(cmp) , cmp)
	fmt.Printf("Real part: %v, %T\n", imag(cmp) , cmp)
	fmt.Printf("%v, %T\n", b , b)
	fmt.Printf("%v, %T\n", r , r)
}