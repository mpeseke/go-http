package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)




func main() {


}

func inputControl() int {
	var number int
	fmt.Println("What Pokemon number would you like to search for?")
	fmt.Scan(&number)

	return number
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
