package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	httpMethod := "GET"
	url := "http://localhost:8080"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)

	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))

}
