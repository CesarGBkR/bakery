package objects

import (
  //"bakery/Domain/Object/CurlObjects"
  "bakery/Domain/Object/NmapObjects"
)

type TargetObject struct {
  IP string
  NS string
  PORTS string
  TYPE string
  RATE int
}

type Response struct {
  Enumeration Enumeration
}

type Enumeration struct {
  PortScann PortScannResponse
  Fuzzing FuzzingResponse
}

type FuzzingResponse struct {

}

type PortScannResponse struct {
  NmapResponse NmapObjects.ScannResponse

}
