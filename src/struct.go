package main

import (
	"fmt"
	"reflect"
)

// User defined data type
type User struct{

	userName string
	uid int
	age int
	tag string
	topic []string
}

// Composition - kidda like Inheritance in OOPS

type Animal struct {
	Name string `required max:"100"` // Line is meaningless to go needs some validation library
	Origin string
	Genus string
}

type Bird struct {

	Animal
	SpeedKPH float32
	CanFly bool

}


func main()  {
	
	john := User{
		userName: "John",
		uid: 101,
		age: 408,
		tag: "Keep on living",
		topic: []string{
			"live",
			"immortal",
		},
	}

	// Not recommended to write like this
	/*gopher := User{
		"gopher",
		101,
		6,
		"Love beign the mascot of Golang",
		[] string{ 
			"gopher",
			"Go",
		},
	}*/

	// Anonymous Struct
	gopher := struct{name string}{name: "Golang"}

	anotherJohn := john
	anotherJohn.userName = "Joe"

	sameJohn := &john
	sameJohn.userName = "Mr. Doe"

	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false

	b := Bird{
		Animal: Animal{
			Name: "Golden Eagle",
			Origin: "IDK",
			Genus: "MKC",
		},
		CanFly: true,
		SpeedKPH: 280,
	}

	t := reflect.TypeOf(Animal{})

	field,_ := t.FieldByName("Name")

	fmt.Println(john)
	fmt.Println(john.topic[1])
	fmt.Println(gopher)
	fmt.Println(anotherJohn.userName)
	fmt.Println(bird)
	fmt.Println(b.Animal)
	fmt.Println(field.Tag)

}