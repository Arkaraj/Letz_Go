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
		for j := 1; j <= i; j++ { // j,k = j+1, k+1
			fmt.Printf("%d ", k)
			k ++
		}
		fmt.Println()
	}
}

func main () {
	
	// floyd(10)

	// Float - IEEE standards => 32bit floating point nos & 64 bits...

	var i int = 64
	j := float64(i)
	k := strconv.Itoa(i)
	sez := 8 // 2^3
	var pi float64 = 3.14
	pi = 13.7e72
	t := true

	if t {
		fmt.Printf("%v, %T\n", i, i)
		fmt.Printf("%v, %T\n", math.Sqrt(j) , j)
		fmt.Printf("%s, %T\n", k , k)
		fmt.Println(sez << 3) // 2^3 * 2^3
		fmt.Println(sez >> 3) // 2^3 / 2^3
		fmt.Printf("%v , %T\n", pi, pi)
	}
}