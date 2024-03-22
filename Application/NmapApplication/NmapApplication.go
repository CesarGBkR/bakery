package NmapApplication

import (
  "bakery/Domain/Enumeration/Portscann"
)

func AllPorts(TARGET string, RATE int) {
  TYPE := "OSScann"
  PORTS := "1-65535"
  Portscann.Portscann(TARGET, TYPE, PORTS, RATE)
}

func ScanPorts() {
  
} 
