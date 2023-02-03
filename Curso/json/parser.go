package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Origin string
	User   string
	Active bool
}

func main() {

	bytedata, err := ioutil.ReadFile("users.json")

	if err != nil {
		fmt.Println("Error")
	}

	var payload Data
	err = json.Unmarshal(bytedata, &payload)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}

	fmt.Printf("origin: %s\n", payload.Origin)
	fmt.Printf("user: %s\n", payload.User)
	fmt.Printf("status: %t\n", payload.Active)

}
