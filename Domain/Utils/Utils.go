package Utils

import (
  //"fmt"
  
  "bakery/Domain/Utils/Requester"
  "bakery/Domain/Utils/Curl"
  "bakery/Domain/Object"
)

func Requester(TARGET objects.TargetObject){
 RequesterDomain.Get(TARGET) 
}

func Curl() {
  CurlDomain.Curl()
}

