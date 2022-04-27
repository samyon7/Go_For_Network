package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func main(){
	person := Person{
		Name: Name {Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.namea"}, Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}
	}

	saveJSON("person.json", person)
}

func saveJSON(){

}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
                os.Exit(1)
        }
}
