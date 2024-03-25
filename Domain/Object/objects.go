package objects

import (
  "bakery/Domain/Object/NmapObjects"
)

type TargetObject struct {
  IP string
  NS string
  LVL int
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
