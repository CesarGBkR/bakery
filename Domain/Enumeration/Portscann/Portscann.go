package Portscann

import (
  "bakery/Service/Enumeration/Portscann"
)

func Portscann(TARGET string, TYPE string, PORTS string, RATE int) {
  nmap.Nmap(TARGET, TYPE, PORTS, RATE)
}
