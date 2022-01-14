package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Respose Type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(bytes)

	// fmt.Println(content)

	tours := tourtoJson(content)
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}

}

func tourtoJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	deoder := json.NewDecoder(strings.NewReader(content))
	_, err := deoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for deoder.More() {
		err := deoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
