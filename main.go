package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) GetUser() {
	person.Name = "liushiju"
	fmt.Println("GetUser() name=", person.Name)
}

func (p *Person) Modify(name string) {
	p.Name = name
	fmt.Println("Modify name=", p.Name)
}

func main() {
	var p Person
	p.Name = "xing"
	p.GetUser()
	p.Modify("xing")
	fmt.Println("main() p.Name=", p.Name)
}
