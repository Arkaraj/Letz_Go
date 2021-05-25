package main

import (
	"fmt"
	"math/rand"
)


func isPallindrome()  {
	
	var num, digit, rev, n int = 0,0,0,0

	fmt.Println("Enter a positive number: ")
	fmt.Scanln(&num)

	n = num

	// Do while loop
	for ok := true; ok; ok = num != 0 {
       
		digit = num % 10
		rev = (rev * 10) + digit
        num = num / 10

}


	fmt.Println(" The reverse of the number is: ", rev)

	if n == rev {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println(" The number is not a palindrome.")
	}

}

func proSwitch(send bool) {

	// i := int(math.Floor(float64(rand.Intn(1000))))
	i := 0
	var(
		score int = 0
		totalNoOfGames = 0
	)
	var x int = 0
	var cont string

	if(send) {

		for ok := true; ok; ok = cont != "n" {

			i = rand.Intn(1000)
			totalNoOfGames ++

			fmt.Println("Enter your number: ")
			fmt.Scanln(&x)

			switch {
			case x<i:
				fmt.Println("Your Number is less than",i)
				// fallthrough it will go to next case doesn't care about the condition
			case x == i:
				fmt.Println("Your Number equal to ",i)
				fmt.Println("Ohh so closeee....")
			default:
				fmt.Println("Your Number greater than ",i)
				score ++;
			}
			fmt.Println("Try again?(y/n)")
			fmt.Scanln(&cont)
		}

		fmt.Println("Your final Score is: ",score,"/", totalNoOfGames)

	} else {
		fmt.Println("Bye")
	}

}


func main()  {
	
	var play string

	fmt.Println("Do you want to play this game?")
	fmt.Scanln(&play)

	// Switch Case
	// switch play {
	// case "y":
	// 	fmt.Println("You are playing the game")
	// 	isPallindrome()
	// case "Y":
	// 	fmt.Println("You are playing the game")
	// 	isPallindrome()
	// default:
	// 	fmt.Println("Bye")
	// }

	if play == "y" || play == "Y" {
		fmt.Println("You are playing the game")

		isPallindrome()

	} else if play == "t" {
		
		s := "Hello from Go!"
		// for range of string, arrays, maps, etc
		for k, v := range s {
			fmt.Println(k,string(v))
		}

		studentMarks := map[string]int32{

			"A": 9,
			"B": 7,
			"C": 5,
			"D": 10,
	
		}

		for k, v := range studentMarks {
			fmt.Println(k,v)
		}

	} else {
		fmt.Println("Suprise game, guess greater than the computer number!")
		proSwitch(true)
	}


}