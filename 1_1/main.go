package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Hi, my name is %s!", h.Name)
}

type Action struct {
	Human
	Skill string
}

func (a Action) Perform() {
	fmt.Printf("%s is doing skill: %s.\n", a.Name, a.Skill)
}

func main() {
	action := Action{
		Human: Human{Name: "Shokhrukh", Age: 21},
		Skill: "Coding",
	}

	action.Speak()
	action.Perform()
}
