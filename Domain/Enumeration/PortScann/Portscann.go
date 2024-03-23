package PortScann

import (
  "fmt"
  
  "bakery/Domain/Enumeration/PortScann/NmapDomain"
  //"bakery/Domain/Object/NmapObjects"
)

func ScannAllPorts(TARGET string, RATE int) {
  response := NmapDomain.ScannAllPorts(TARGET, RATE)
   
  fmt.Printf("%v", response.Hosts.Ports[0].State)

}

func ScannPort(TARGET string, PORTS string, RATE int) {
  NmapDomain.ScannPort(TARGET, PORTS, RATE)
}

func ScannOS(TARGET string, PORTS string, RATE int) {  
  NmapDomain.ScannOS(TARGET, PORTS, RATE)
}

func ServiceScann(TARGET string, PORTS string, RATE int) {
  NmapDomain.ServiceScann(TARGET, PORTS, RATE)
}
