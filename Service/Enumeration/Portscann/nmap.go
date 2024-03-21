package nmap

import (
  "context"
  "fmt"
  "log"
  "time"
  
  "github.com/Ullaakut/nmap/v3"
)

func errorPrint( TYPE string, Err error) {

  if Err != nil {
    log.Fatalf("[-] Error in %s nmap scanner:\n%v", TYPE, Err)
  }
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
      
      if err != nil {
        log.Fatalf("[-] Error in %s nmap scanner:\n%v", TYPE, err)
      }

      result, warnings, err := scanner.Run()
      if len(*warnings) > 0 {
        log.Printf("[i] run finished with warnings:\n%s", *warnings)
      }

      //warningsPrint(warnings)
      errorPrint(TYPE, err)
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

      //fmt.Printf("[+] Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
    
    case "OSScann":
      scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(TARGET),
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
      //warningsPrint(warnings)
      errorPrint(TYPE, err)
      fmt.Printf("%d",len(result.Hosts))

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
      //warningsPrint(warnings)
      errorPrint(TYPE, err)
      //fmt.Printf(result)
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

  fmt.Println("[i] Starting Scann ")

  scann(TARGET, TYPE, PORTS, RATE)

  //fmt.Println("Nmap")
}
