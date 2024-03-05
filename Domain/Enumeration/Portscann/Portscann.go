package Portscann

import (
  "bakery/Service/Enumeration/Portscann"
)

func Portscann(TARGET string, PORT string) {
  nmap.Nmap(TARGET, PORT)
}
