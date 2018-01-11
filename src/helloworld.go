package main

import "fmt"

type Salutation struct{
	name string
	greeting string
}

func Greet (salutation Salutation){
	msg, alt := CreateMessage(salutation.name, salutation.greeting, "yo", "sup")
	fmt.Println(msg)
	fmt.Println(alt)

}

func CreateMessage(name string, greeting ...string) (message string, alternate string){
	//fmt.Println(len(greeting))
message = greeting[1] + " " + name
alternate = "Hey! " + name
return
}

func main() {

	s := Salutation{"Bob", "Hello"}
	Greet(s)
}
