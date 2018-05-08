package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	testInterface1()

}
func testInterface1() {
	//var peo People = Stduent{}  //compile error cannot use Stduent literal (type Stduent) as type People in assignment: Stduent does not implement People (Speak method has pointer receiver)
	//var peo = Stduent{}
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
