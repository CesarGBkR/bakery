package requester


import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"

  "bakery/Domain/Object"
)

func requester(url string) {
    response, err := http.Get(url)
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
    fmt.Println("GET Response:", string(body))
}

func Get(TARGET objects.TargetObject) {
  var PORTS []string 
  if len(TARGET.PORTS) < 0 || TARGET.PORTS == "1-65535"{
    PORTS = append(PORTS, "80") 
    PORTS = append(PORTS, "8080") 
    PORTS = append(PORTS, "443") 
  }else{
    PORTS = strings.Split(TARGET.PORTS, ",")
  }
  for _, Port := range PORTS {

    urlIP := "http://"+TARGET.IP+":"+Port+"/"
    urlNS := "http://"+TARGET.NS+":"+Port+"/"
    requester(urlIP)
    requester(urlNS)
  }
}

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
