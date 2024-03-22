package nmap

import (
  "context"
  "fmt"
  "log"
  "time"
  
  "github.com/Ullaakut/nmap/v3"
	osfamily "github.com/Ullaakut/nmap/v3/pkg/osfamilies"
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

//func warningsPring( Warnings string) {
  //if len(Warnings) > 0 {
    //log.Printf("[i] nmap run finished with warnings:\n%s", *Warnings)
  //}
//}

func scann(TARGET string, TYPE string, PORTS string, RATE int) {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()
  
  switch TYPE {
    case "ScanPorts": 
      scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(TARGET),
        nmap.WithPorts(PORTS),
        nmap.WithMinRate(RATE),
        nmap.WithDisabledDNSResolution(),
        nmap.WithSkipHostDiscovery(),
      )
      
      errorPrint(TYPE, err)
      
      result, warnings, err := scanner.Run()
      
      if len(*warnings) > 0 {
        log.Printf("[i] run finished with warnings:\n%s", *warnings)
      }
      
      errorPrint(TYPE, err)
      
      if len(result.Hosts) > 0 {
        for _, host := range result.Hosts {
          
          if len(host.Ports) == 0 || len(host.Addresses) == 0 {
            continue
          } 
          
          fmt.Printf("[+] Host %q:\n", host.Addresses[0])
          fmt.Printf("\tState\tPORT\tProtocol\tS.Name\n")
          
          for _,port := range host.Ports {
            fmt.Printf("\t[%s]\t%d\t%s\t\t%s\n", port.State, port.ID, port.Protocol, port.Service.Name)
          }
        }
      }
       
    case "OSScann":

      scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(TARGET),
        nmap.WithPorts(PORTS),
        nmap.WithMinRate(RATE),
        nmap.WithDisabledDNSResolution(),
        nmap.WithSkipHostDiscovery(),
        nmap.WithOSDetection(),
      )
      errorPrint(TYPE, err)

      result, warnings, err := scanner.Run()
      
      if len(*warnings) > 0 {
        log.Printf("[i] run finished with warnings:\n%s", *warnings)
      }
      
      errorPrint(TYPE, err)

      if len(result.Hosts) > 0 {
        var (Linux, Windows, Other int)

        for _, host := range result.Hosts {
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

          Os := ProbableOS(Linux, Windows, Other)
          
          fmt.Printf("[+] Host %q:\n", host.Addresses[0])
          fmt.Printf("[i] Probable OS: %s\n", Os)
        }
      }

    case "ServiceScann": 
      scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(TARGET),
        nmap.WithPorts(PORTS),
        nmap.WithMinRate(RATE),
        nmap.WithServiceInfo(),
      )

      errorPrint(TYPE, err)
      result, warnings, err := scanner.Run()
      if len(*warnings) > 0 {
        log.Printf("[i] run finished with warnings:\n%s", *warnings)
      }
      errorPrint(TYPE, err)
      fmt.Printf("%d",len(result.Hosts))
    
    case "ScriptScann":
      scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(TARGET),
        nmap.WithPorts(PORTS),
        nmap.WithMinRate(RATE),
        nmap.WithDefaultScript(),
      )
      errorPrint(TYPE, err)
      result, warnings, err := scanner.Run()
      if len(*warnings) > 0 {
        log.Printf("[i] run finished with warnings:\n%s", *warnings)
      }
      //warningsPrint(warnings)
      errorPrint(TYPE, err)
      //fmt.Printf(result)
      fmt.Printf("%d",len(result.Hosts))

      // TODO: Add optionsScanTechniques mngmnt (reference https://github.com/Ullaakut/nmap/blob/58d93393be5926c8b541e049cc357b6cb9e3eb5e/optionsScanTechniques.go)

    default: 
      log.Fatalf("[-] Error creating nmap scanner, Error on TYPE")
  }

}

// TODO add List
// Example to Nmap pharams: Nmap("127.0.0.0", "1-10000", "5000")
func Nmap(TARGET string,TYPE string, PORTS string, RATE int) {

  fmt.Println("[i] Starting Scann \n")

  scann(TARGET, TYPE, PORTS, RATE)
  fmt.Println("\n")
}
