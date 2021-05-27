package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Control Flows - Defer,Panic, Recover

func readRobots() {
	res,err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		// panic(err.Error())
		log.Fatal(err)
	}
	// opening and closing of a resource
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s",robots)
}

func panicker()  {
	fmt.Println("About To Panic")
	// Anonymus function
	defer func()  {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()

	panic("Something bad occured!")
}

func main() {

	// Defer - LIFO

	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")

	/* a := 5
	defer fmt.Println(a) // will print 5
	a = 10 */

	// readRobots()

	/* PANIC */

	// panic kindda like exception

	// defer - panic - return value in a function
	// a,b := 1,0
	// ans := a/b
	// fmt.Println(ans)
	// panic("Something bad occured!")

	fmt.Println("Start")
	panicker()
	fmt.Println("End")

}

