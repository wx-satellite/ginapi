package main

import (
	"fmt"
)

type Person struct {
}

func (person *Person) say() {
	fmt.Println(person)
}


type Student struct {
	Person
	name string
}


type Dog struct {
	Name string
	Age int
}

func (dog Dog) String() string {
	return  "it is a test"
}



func main()  {
	var (
		dog Dog
	)
	dog.Name = "weixin"
	dog.Age = 12
	fmt.Println(fmt.Sprintf("%v", dog))
	fmt.Println(fmt.Sprintf("%s", dog))
}
