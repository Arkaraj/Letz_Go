package main

import "fmt"

// func sort(){

// }

func main() {

	var classCode [3]string
	classCode[0] = "xyDf6"
	classCode[1] = "jui9$"
	classCode[2] = "hGo8@a"

	var identityMatrix [3][3]int = [3][3]int{{1,0,0},{0,1,0},{0,0,1}}

	arr := [7]int{4,17,13,55,32,149,7}
	arr2 := [...]int{4,17,13,55,32,149}
	copy := arr2 // normal copy
	copy[3] = 9000 // won't affect arr2
	// copy := &arr2 // address copy
	// copy[3] = 9000 // affects arr2

	/*
	* Array Functions
	* len(arr) - returns length of array 
	* cap(arr) - returns capacity of array 
	* Slices:
	* arr[:] - slices all elements
	* arr[n:] - slices nth element to the end
	* arr[:n] - slices from 1st ele to the nth, for last ele [:len(arr)-1]
	* arr[n:m] - slices from nth element to the mth element
	
	* make(array type, size, capacity)
	*/

	a := make([] int, 3)

	a = append(a, 1) // adds 1 at the end
	// a = append(a, 3,4,5,6)

	fmt.Println(classCode)
	fmt.Println("classCode #1: ",classCode[1])
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(copy)
	fmt.Println(a)

	fmt.Println("IdentityMatrix is: ")
	for i := 0; i < len(identityMatrix); i++ {
		fmt.Println(identityMatrix[i])
	}
	
}