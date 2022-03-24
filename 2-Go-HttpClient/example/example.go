package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alexfabianojr/Go-HttpClient/2-Go-HttpClient/example/gohttp"
)

var (
	gethubClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)

	client.SetHeaders(commonHeaders)

	return client
}

func exampleUsage() {

	response, err := gethubClient.Get("https://api.github.com", nil)

	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(response.Status)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
