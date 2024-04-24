package PortScann

import (
	//"fmt"

	"bakery/Domain/Enumeration/PortScann/NmapDomain"
	"bakery/Domain/Object"
)

func ScannAllPorts(TARGET objects.TargetObject) objects.PortScannResponse {

	var response objects.PortScannResponse
	Nmap := NmapDomain.ScannAllPorts(TARGET)
	response.NmapResponse = Nmap
	return response
}

func ScannPort(TARGET objects.TargetObject) objects.PortScannResponse {

	var response objects.PortScannResponse
	Nmap := NmapDomain.ScannPort(TARGET)
	response.NmapResponse = Nmap
	return response
}

func ScannOS(TARGET objects.TargetObject) objects.PortScannResponse {

	var response objects.PortScannResponse
	Nmap := NmapDomain.ScannOS(TARGET)
	response.NmapResponse = Nmap
	return response
}

func ScannService(TARGET objects.TargetObject) objects.PortScannResponse {

	var response objects.PortScannResponse
	Nmap := NmapDomain.ScannService(TARGET)
	response.NmapResponse = Nmap
	return response
}

func ScannScript(TARGET objects.TargetObject) objects.PortScannResponse {

	var response objects.PortScannResponse
	Nmap := NmapDomain.ScannScript(TARGET)
	response.NmapResponse = Nmap
	return response
}
