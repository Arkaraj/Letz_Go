package main

import (
	"fmt"
	"math"
	"strconv"
)

var lifeAnswer int = 42
var floydNumber int = 4

// Floyd's Triangle
func floyd(num int) {

	var k int = 1
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", k)
			k ++
		}
		fmt.Println()
	}
}

func main () {
	
	// floyd(10)

	var i int = 64
	j := float64(i)
	k := strconv.Itoa(i)
	t := true

	if t {
		fmt.Printf("%v, %T\n", i, i)
		fmt.Printf("%v, %T\n", math.Sqrt(j) , j)
		fmt.Printf("%s, %T\n", k , k)
	}
}