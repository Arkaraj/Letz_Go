package main

import "fmt"

func sort(arr* [7]int){

	n := len(arr)

	for i:=0; i< n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (arr[j] > arr[j+1]) {
				temp := arr[j]
				arr[j] = arr[j+1];
    			arr[j+1] = temp;
			}
		}
	}

}

func main() {

	list := [7]int{4,17,13,55,32,149,7}

	sort(&list)

	fmt.Println("list: ", list)
	
}