package NmapApplication

import (
  "bakery/Domain/Enumeration/Portscann"
)

func AllPorts(TARGET string) {
  Flags := []string{"WithSkipHostDiscovery()", "WithDisableDNSResolution()", "WithPorts(1-63000)", "WithMinRate(5000)"}
  Portscann.Portscann(TARGET, Flags)
}
