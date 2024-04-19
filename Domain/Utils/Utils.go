package Utils

import (
  //"fmt"
  
  "bakery/Domain/Utils/Requester"
  "bakery/Domain/Utils/Curl"
  "bakery/Domain/Object/RequesterObjects"
)

func Requester(url string) RequesterObjects.Response {
 return RequesterDomain.Requester(url) 
}

func Curl() {
  CurlDomain.Curl()
}

