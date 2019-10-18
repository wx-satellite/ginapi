package main

import "fmt"

type Person struct {
}

func (person *Person) say() {
	fmt.Println(person)
}


type Student struct {
	Person
	name string
}

func main()  {
	var (
		student *Student
	)
	student = &Student{name: "weixin"}

	fmt.Println(student)
	student.say()
}
