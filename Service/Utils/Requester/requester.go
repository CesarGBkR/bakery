package requester

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"

	//"bakery/Domain/Object"
	"bakery/Domain/Object/RequesterObjects"
)

func Requester(url string) RequesterObjects.Response {

	var RequesterResponse RequesterObjects.Response
	response, err := http.Get(url)
	fmt.Printf("\nRequestDone\n")
	if err != nil {
		RequesterResponse.Error = "[-] Error:\n\t" + err.Error()
		return RequesterResponse
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		RequesterResponse.Warning = "[!] Warning, response with error reading response body:\n\t" + err.Error()
		return RequesterResponse
	} else {
		RequesterResponse.Body = "GET Response:" + string(body)
		return RequesterResponse
	}
}

//func Get(TARGET objects.TargetObject) []RequesterObjects.Response {
//var PORTS []string
//var Responses []RequesterObjects.Response
//WebFinder := make(chan objects.PortScannResponse, 2)
//var wg sync.WaitGroup
//
//if len(TARGET.PORTS) < 0 || strings.Contains(TARGET.PORTS, "-"){
//wg.Add(3)
//PORTS = append(PORTS, "80","8080","443")
//}else{
//PORTS = strings.Split(TARGET.PORTS, ",")
//wg.Add(len(PORTS))
//}
//for id, Port := range PORTS {
//
//urlIP := "http://"+TARGET.IP+":"+Port+"/"
//urlNS := "http://"+TARGET.NS+":"+Port+"/"
//go requester(id, &wg, WebFinder, urlIP)
//NSResponse := requester(urlNS)
//
//}
//return Responses
//}

func Post() {
	url := "https://jsonplaceholder.typicode.com/posts"

	jsonPayload := []byte(`{"title": "foo", "body": "bar", "userId": 1}`)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("POST Response:", string(body))
}

func Put() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	jsonPayload := []byte(`{"id": 1, "title": "foo", "body": "bar", "userId": 1}`)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("PUT Response:", string(body))
}

func Delete() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	request, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code:", response.StatusCode)
		return
	}
	fmt.Println("DELETE Response:", response.Status)
}
