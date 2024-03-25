package nmap

import (
  "context"
  "fmt"
  //"log"
  "time"
  
  "github.com/Ullaakut/nmap/v3"
	osfamily "github.com/Ullaakut/nmap/v3/pkg/osfamilies"
  "bakery/Domain/Object/NmapObjects"
)

func errorPrint( TYPE string, Err error) {

  if Err != nil {
    fmt.Printf("[-] Error in %s nmap scanner:\n%v", TYPE, Err)
  }
}

func ProbableOS(Linux, Windows, Other int) string {
  Os := "Other"
  max := Other
  if Windows > max {
    max = Windows
    Os = "Windows"
  }
  if Linux > max {
    max = Linux
    Os = "Linux"
  }
  return Os
}

func PortScann(TARGET string, PORTS string, RATE int, TYPE string) NmapObjects.ScannResponse {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()

  var Response NmapObjects.ScannResponse
      
  scanner, errBuild := nmap.NewScanner(
    ctx,
    nmap.WithTargets(TARGET),
    nmap.WithPorts(PORTS),
    nmap.WithMinRate(RATE),
    nmap.WithDisabledDNSResolution(),
    nmap.WithSkipHostDiscovery(),
  )
      
  Response.ErrBuild = errBuild
  errorPrint(TYPE, errBuild) 
  result, warnings, errExec := scanner.Run()
  Response.ErrExec, Response.Warn = errExec, *warnings 
  
  if len(result.Hosts) > 0 {
    var HostResponse NmapObjects.Host
    host := result.Hosts[0]
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      fmt.Printf("[-] No Ports or Addresses for: %s\n", TARGET)
    }else{
      var PortsResponse []NmapObjects.Port
      for _,port := range host.Ports {
        Port := NmapObjects.Port{
          State: port.State.String(),
          ID: port.ID,
          Protocol: port.Protocol,
        }
        PortsResponse = append(PortsResponse,Port)
      }
      HostResponse.Ports = PortsResponse
    }
    Response.Hosts = HostResponse
  }
  return Response
}

func ServiceScann(TARGET string, PORTS string, RATE int, TYPE string) NmapObjects.ScannResponse {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()
  
  var Response NmapObjects.ScannResponse

  scanner, errBuild := nmap.NewScanner(
    ctx,
    nmap.WithTargets(TARGET),
    nmap.WithPorts(PORTS),
    nmap.WithMinRate(RATE),
    nmap.WithServiceInfo(),
  )
  
  Response.ErrBuild = errBuild
  errorPrint(TYPE, errBuild) 
  result, warnings, errExec := scanner.Run()
  Response.ErrExec, Response.Warn = errExec, *warnings 
    
  if len(result.Hosts) > 0 {
    var HostResponse NmapObjects.Host
    host := result.Hosts[0] 
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      fmt.Printf("[-] No Ports or Addresses for: %s", TARGET)
    }else{
      var PortsResponse []NmapObjects.Port
      for _,port := range host.Ports {

        ServiceResponse := NmapObjects.Service {
          Name: port.Service.Name,
          Version: port.Service.Version,
          Product: port.Service.Product,
          Extra: port.Service.ExtraInfo,
        }

        Port := NmapObjects.Port {
          State: port.State.String(),
          ID: port.ID,
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

func OSScann(TARGET string, PORTS string, RATE int, TYPE string) NmapObjects.ScannResponse {
  
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()

  var Response NmapObjects.ScannResponse

  scanner, errBuild := nmap.NewScanner(
    ctx,
    nmap.WithTargets(TARGET),
    nmap.WithPorts(PORTS),
    nmap.WithMinRate(RATE),
    nmap.WithDisabledDNSResolution(),
    nmap.WithSkipHostDiscovery(),
    nmap.WithOSDetection(),
  )

  Response.ErrBuild = errBuild
  errorPrint(TYPE, errBuild) 
  result, warnings, errExec := scanner.Run()
  Response.ErrExec, Response.Warn = errExec, *warnings 
 

  if len(result.Hosts) > 0 {
    var HostResponse NmapObjects.Host
    var (Linux, Windows, Other int)
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
    //fmt.Printf("[+] Host %q:\n", host.Addresses[0])
    //fmt.Printf("[i] Probable OS: %s\n", Os)
  }
  return Response
}


// TODO: Add functionality
func ScriptScann(TARGET string, PORTS string, RATE int, TYPE string) {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()
 
  var Response NmapObjects.ScannResponse

  scanner, errBuild := nmap.NewScanner(
    ctx,
    nmap.WithTargets(TARGET),
    nmap.WithPorts(PORTS),
    nmap.WithMinRate(RATE),
    nmap.WithDefaultScript(),
  )

  Response.ErrBuild = errBuild
  errorPrint(TYPE, errBuild) 
  result, warnings, errExec := scanner.Run()
  Response.ErrExec, Response.Warn = errExec, *warnings 
  fmt.Printf("%v",  result)
}


// TODO: Add optionsScanTechniques mngmnt (reference https://github.com/Ullaakut/nmap/blob/58d93393be5926c8b541e049cc357b6cb9e3eb5e/optionsScanTechniques.go)
