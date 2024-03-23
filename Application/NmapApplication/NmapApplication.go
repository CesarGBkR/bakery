package NmapApplication

import (
  "bakery/Domain/Enumeration/Portscann"
)

func ScannAllPorts(TARGET string, RATE int) {
  TYPE := "ScannAllPorts"
  PORTS := "1-65535"
  Portscann.Portscann(TARGET, PORTS, RATE, TYPE)
}

func ScannPort(TARGET string, PORTS string, RATE int) {
  TYPE := "ScannPorts"
  Portscann.Portscann(TARGET, PORTS, RATE, TYPE)
}

func ScannOS(TARGET string, PORTS string, RATE int) {  
  TYPE := "ScannOS"
  Portscann.OSScann(TARGET, PORTS, RATE, TYPE)
}

func ServiceScann(TARGET string, PORTS string, RATE int) {
  TYPE := "ServiceScann" 
  Portscann.ServiceScann(TARGET, PORTS, RATE, TYPE)
}
