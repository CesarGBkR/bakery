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

func ScannPort(TARGET string, PORTS string, RATE int) objects.PortScannResponse {

  var response objects.PortScannResponse
  Nmap := NmapDomain.ScannPort(TARGET, PORTS, RATE)
  response.NmapResponse = Nmap
  return response
}

func ScannOS(TARGET string, PORTS string, RATE int) objects.PortScannResponse {  
  
  var response objects.PortScannResponse
  Nmap := NmapDomain.ScannOS(TARGET, PORTS, RATE)
  response.NmapResponse = Nmap
  return response
}

func ServiceScann(TARGET string, PORTS string, RATE int) objects.PortScannResponse {
  
  var response objects.PortScannResponse
  Nmap := NmapDomain.ServiceScann(TARGET, PORTS, RATE)
  response.NmapResponse = Nmap
  return response
}
