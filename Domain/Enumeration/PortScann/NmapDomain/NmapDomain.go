package NmapDomain

import (
  "bakery/Service/Enumeration/Portscann"
  "bakery/Domain/Object/NmapObjects"
)

func ScannAllPorts(TARGET string, RATE int) NmapObjects.ScannResponse {
  TYPE := "ScannAllPorts"
  PORTS := "1-65535"
  return nmap.PortScann(TARGET, PORTS, RATE, TYPE)
}

func ScannPort(TARGET string, PORTS string, RATE int) NmapObjects.ScannResponse {
  TYPE := "ScannPorts"
  return nmap.PortScann(TARGET, PORTS, RATE, TYPE)

}

func ScannOS(TARGET string, PORTS string, RATE int) NmapObjects.ScannResponse {  
  TYPE := "ScannOS"
  return nmap.OSScann(TARGET, PORTS, RATE, TYPE)
}

func ScannService(TARGET string, PORTS string, RATE int) NmapObjects.ScannResponse {
  TYPE := "ServiceScann" 
  return nmap.ScannService(TARGET, PORTS, RATE, TYPE)
}

func ScannScript(TARGET string, PORTS string, RATE int) NmapObjects.ScannResponse {
  TYPE := "ServiceScann" 
  return nmap.ScannScript(TARGET, PORTS, RATE, TYPE)
}
