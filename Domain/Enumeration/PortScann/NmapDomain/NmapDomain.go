package NmapDomain

import (
	"bakery/Domain/Object"
	"bakery/Domain/Object/NmapObjects"
	"bakery/Service/Enumeration/Portscann"
)

func ScannAllPorts(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	TARGET.TYPE = "ScannAllPorts"
	TARGET.PORTS = "1-65535"
	return nmap.PortScann(TARGET)
}

func ScannPort(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	TARGET.TYPE = "ScannPorts"
	return nmap.PortScann(TARGET)

}

func ScannOS(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	TARGET.TYPE = "ScannOS"
	return nmap.OSScann(TARGET)
}

func ScannService(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	TARGET.TYPE = "ServiceScann"
	return nmap.ScannService(TARGET)
}

func ScannScript(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	TARGET.TYPE = "ServiceScann"
	return nmap.ScannScript(TARGET)
}
