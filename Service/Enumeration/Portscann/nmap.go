package nmap

import (
	"context"
	"fmt"
	//"log"
	"time"

	"bakery/Domain/Object"
	"bakery/Domain/Object/NmapObjects"
	"github.com/Ullaakut/nmap/v3"
	osfamily "github.com/Ullaakut/nmap/v3/pkg/osfamilies"
)

func errorPrint(TYPE string, Err error) {

	if Err != nil {
		fmt.Printf("[-] Error in %s nmap scanner:\n%v", TYPE, Err)
	}
}

func ProbableOS(Linux, Windows, Other int) string {
	osNames := []string{"Linux", "Windows", "Other"}
	osValues := []int{Linux, Windows, Other}

	max := Other
	maxIndex := 2

	for i, value := range osValues {
		if value > max {
			max = value
			maxIndex = i
		}
	}

	return osNames[maxIndex]
}

func PortScann(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var Response NmapObjects.ScannResponse

	scanner, errBuild := nmap.NewScanner(
		ctx,
		nmap.WithTargets(TARGET.IP),
		nmap.WithPorts(TARGET.PORTS),
		nmap.WithMinRate(TARGET.RATE),
		nmap.WithDisabledDNSResolution(),
		nmap.WithSkipHostDiscovery(),
	)

	Response.ErrBuild = errBuild
	errorPrint(TARGET.TYPE, errBuild)
	result, warnings, errExec := scanner.Run()
	Response.ErrExec, Response.Warn = errExec, *warnings

	if len(result.Hosts) > 0 {
		var HostResponse NmapObjects.Host
		host := result.Hosts[0]
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			fmt.Printf("[-] No Ports or Addresses for: %s\n", TARGET)
		} else {
			var PortsResponse []NmapObjects.Port
			for _, port := range host.Ports {
				Port := NmapObjects.Port{
					State:    port.State.String(),
					ID:       port.ID,
					Protocol: port.Protocol,
				}
				PortsResponse = append(PortsResponse, Port)
			}
			HostResponse.Ports = PortsResponse
		}
		Response.Hosts = HostResponse
	}
	return Response
}

func ScannService(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var Response NmapObjects.ScannResponse

	scanner, errBuild := nmap.NewScanner(
		ctx,
		nmap.WithTargets(TARGET.IP),
		nmap.WithPorts(TARGET.PORTS),
		nmap.WithMinRate(TARGET.RATE),
		nmap.WithServiceInfo(),
	)

	Response.ErrBuild = errBuild
	errorPrint(TARGET.TYPE, errBuild)
	result, warnings, errExec := scanner.Run()
	Response.ErrExec, Response.Warn = errExec, *warnings

	if len(result.Hosts) > 0 {
		var HostResponse NmapObjects.Host
		host := result.Hosts[0]
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			fmt.Printf("[-] No Ports or Addresses for: %s", TARGET)
		} else {
			var PortsResponse []NmapObjects.Port
			for _, port := range host.Ports {

				ServiceResponse := NmapObjects.Service{
					Name:    port.Service.Name,
					Version: port.Service.Version,
					Product: port.Service.Product,
					Extra:   port.Service.ExtraInfo,
				}

				Port := NmapObjects.Port{
					State:    port.State.String(),
					ID:       port.ID,
					Protocol: port.Protocol,
					Services: ServiceResponse,
				}

				PortsResponse = append(PortsResponse, Port)
			}
			HostResponse.Ports = PortsResponse
		}
		Response.Hosts = HostResponse
	}
	return Response
}

func OSScann(TARGET objects.TargetObject) NmapObjects.ScannResponse {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var Response NmapObjects.ScannResponse

	scanner, errBuild := nmap.NewScanner(
		ctx,
		nmap.WithTargets(TARGET.IP),
		nmap.WithPorts(TARGET.PORTS),
		nmap.WithMinRate(TARGET.RATE),
		nmap.WithDisabledDNSResolution(),
		nmap.WithSkipHostDiscovery(),
		nmap.WithOSDetection(),
	)

	Response.ErrBuild = errBuild
	errorPrint(TARGET.TYPE, errBuild)
	result, warnings, errExec := scanner.Run()
	Response.ErrExec, Response.Warn = errExec, *warnings

	if len(result.Hosts) > 0 {
		var HostResponse NmapObjects.Host
		var (
			Linux, Windows, Other int
		)
		host := result.Hosts[0]
		for _, match := range host.OS.Matches {
			for _, class := range match.Classes {
				switch class.OSFamily() {
				case osfamily.Linux:
					Linux++
				case osfamily.Windows:
					Windows++
				default:
					Other++
				}
			}
		}
		HostResponse.OS = ProbableOS(Linux, Windows, Other)
		Response.Hosts = HostResponse
	}
	return Response
}

// TODO: Add functionality
func ScannScript(TARGET objects.TargetObject) NmapObjects.ScannResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var Response NmapObjects.ScannResponse

	scanner, errBuild := nmap.NewScanner(
		ctx,
		nmap.WithTargets(TARGET.IP),
		nmap.WithPorts(TARGET.PORTS),
		nmap.WithMinRate(TARGET.RATE),
		nmap.WithDefaultScript(),
	)

	Response.ErrBuild = errBuild
	errorPrint(TARGET.TYPE, errBuild)
	result, warnings, errExec := scanner.Run()
	Response.ErrExec, Response.Warn = errExec, *warnings
	if len(result.Hosts) > 0 {
		var HostResponse NmapObjects.Host
		host := result.Hosts[0]
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			fmt.Printf("[-] No Ports or Addresses for: %s", TARGET)
		} else {
			var PortsResponse []NmapObjects.Port
			for _, port := range host.Ports {
				ServiceResponse := NmapObjects.Service{
					Name:    port.Service.Name,
					Version: port.Service.Version,
					Product: port.Service.Product,
					Extra:   port.Service.ExtraInfo,
				}

				Port := NmapObjects.Port{
					State:    port.State.String(),
					ID:       port.ID,
					Protocol: port.Protocol,
					Services: ServiceResponse,
				}
				var ScriptsResponse []NmapObjects.Script
				for _, script := range port.Scripts {
					Script := NmapObjects.Script{
						ID:     script.ID,
						Output: script.Output,
					}
					ScriptsResponse = append(ScriptsResponse, Script)
				}
				Port.Scripts = ScriptsResponse
				PortsResponse = append(PortsResponse, Port)
			}
			HostResponse.Ports = PortsResponse
		}
		Response.Hosts = HostResponse
	}
	return Response
}

// TODO: Add optionsScanTechniques mngmnt (reference https://github.com/Ullaakut/nmap/blob/58d93393be5926c8b541e049cc357b6cb9e3eb5e/optionsScanTechniques.go)
