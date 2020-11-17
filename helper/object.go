package helper

import (
	"fmt"
)

type oObject struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func newObject(name string, breed string) *oObject {
	//Logic process
	return &oObject{name, breed}
}

func (p *oObject) secretBark() {
	//Logic
	fmt.Println("meow")
}

func (p *oObject) Bark() {
	p.secretBark()
	fmt.Println("woof boink boink 🐶")
}

func (p *oObject) Description() string {
	//Logic
	return fmt.Sprintf("I am a %s  and my name is %s  ", p.Breed, p.Name)
}

func CallObject() {
	p := newObject("Blacky🦄", "Cat✨")
	p.Bark()
	p.secretBark() // compilation error
	p.Breed = "💖"
	fmt.Println(p.Description()) // I am a Westie and my magic puppy name is 🦄💖
}
