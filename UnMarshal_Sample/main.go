package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {

	fmt.Println("Welcome to Marshal Sample")
	p1 := person{
		First: "Jagadeesh",
	}

	p2 := person{
		First: "JJ",
	}

	xp := []person{p1, p2}

	//Converting objects to json format
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))

	xp2 := []person{}
	unerr := json.Unmarshal(bs, &xp2)
	if unerr != nil {
		log.Panic("Error", err)
	}
	println("json to obejcts - unMarshal :", xp2)

	println(xp2[0].First)
	println(xp2[1].First)
}
