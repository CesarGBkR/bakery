package PortScann

import (
  //"fmt"
  
  "bakery/Domain/Enumeration/PortScann/NmapDomain"
  "bakery/Domain/Object"
)

func ScannAllPorts(TARGET string, RATE int) objects.PortScannResponse {
  
  var response objects.PortScannResponse 
  Nmap := NmapDomain.ScannAllPorts(TARGET, RATE)
  response.NmapResponse = Nmap
  return response
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
