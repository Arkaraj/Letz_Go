package main

import (
	"fmt"
)




func main()  {
	
	// Map => Key value 
	studentMarks := map[string]int32{

		"A": 9,
		"B": 7,
		"C": 5,
		"D": 10,

	}

	// Length function => len(studnetMarks)

	studentMarks["Z"] = 5
	delete(studentMarks, "C")
	// returns value and does it exist
	pop,ok := studentMarks["K"]

	/*
	* Note
	* Changes made to copy of the map will change the original map
	*/

	rmvD := studentMarks
	delete(rmvD, "D")

	fmt.Println(rmvD)
	fmt.Println(studentMarks)
	fmt.Println(pop,ok)

}