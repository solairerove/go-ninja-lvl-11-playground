package main

import (
	"encoding/json"
	"fmt"
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

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Who is Tyler Durden %v", err)
	}

	return bs, nil
}
