package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
	number int
	name   string
}

func main() {
	inputControl()
}

func inputControl() {
	var number string
	fmt.Println("What Pokemon number would you like to search for?")
	fmt.Scan(&number)

	searchMonster(number)
}

func searchMonster(number string) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + number + "/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
