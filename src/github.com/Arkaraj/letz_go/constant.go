package main

// Package level

import (
	"fmt"
	"math"
)

// iota block
const(
	a = iota
	b
	c
) 

const (
	_ = iota
	KB = 1 << (10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main()  {

	const lyfe int = 42
	const gg = 27

	fileSize := 4*math.Pow(10,9)

	fmt.Printf("%v, %T\n", lyfe + gg , lyfe + gg)
	fmt.Printf("%v, %T\n", a , a)
	fmt.Printf("%v, %T\n", b , b)
	fmt.Printf("%v, %T\n", c , c)
	fmt.Printf("%.2fGB, %T\n", fileSize/GB , fileSize/GB)
	fmt.Printf("%.2fZB, %T\n", fileSize/ZB , fileSize/ZB)
}