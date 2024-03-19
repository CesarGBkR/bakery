package Portscann

import (
  "bakery/Service/Enumeration/Portscann"
)

func Portscann(TARGET string, Flags []string) {
  nmap.Nmap(TARGET Flags)
}
