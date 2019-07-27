package main

import (
	"encoding/json"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{
		First: "Tyler",
		Last:  "Durden",
		Sayings: []string{
			"I want you to hit me as hard as you can",
			"The things you own end up owning you.",
		},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(bs))
}
