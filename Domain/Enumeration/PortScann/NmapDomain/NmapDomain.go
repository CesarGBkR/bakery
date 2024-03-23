package NmapDomain

import (
  "bakery/Service/Enumeration/Portscann"
  "bakery/Domain/Object/NmapObjects"
)

func ScannAllPorts(TARGET string, RATE int) NmapObjects.ScannResponse{
  TYPE := "ScannAllPorts"
  PORTS := "1-65535"
  return nmap.PortScann(TARGET, PORTS, RATE, TYPE)
}

func ScannPort(TARGET string, PORTS string, RATE int) {
  TYPE := "ScannPorts"
  nmap.PortScann(TARGET, PORTS, RATE, TYPE)
}

func ScannOS(TARGET string, PORTS string, RATE int) {  
  TYPE := "ScannOS"
  nmap.OSScann(TARGET, PORTS, RATE, TYPE)
}

func ServiceScann(TARGET string, PORTS string, RATE int) {
  TYPE := "ServiceScann" 
  nmap.ServiceScann(TARGET, PORTS, RATE, TYPE)
}
