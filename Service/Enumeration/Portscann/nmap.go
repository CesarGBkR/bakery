package nmap

import (
  "context"
  "fmt"
  "log"
  "time"
  
  "github.com/Ullaakut/nmap/v3"
)

func scann(TARGET string, Flags []string) {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
  defer cancel()
  
  options := nmap.NewScanner(ctx, nmap.WithTargets(TARGET))
  //options = append(options, nmap.WithTargets(TARGET))

    // Agregar las banderas proporcionadas a las opciones del scanner
  for _, flag := range Flags {
    switch flag {
      case "WithSkipHostDiscovery()":
        options = append(options, nmap.WithSkipHostDiscovery())
      case "WithDisableDNSResolution()":
        options = append(options, nmap.WithDisabledDNSResolution())
      case "WithPorts(1-63000)":
        options = append(options, nmap.WithPorts("1-63000"))
      case "WithMinRate(5000)":
        options = append(options, nmap.WithMinRate(5000))
        // Agregar otros casos según sea necesario
    }
  }

    // Crear el scanner con las opciones construidas
    scanner, err := nmap.NewScanner(ctx, options...)
    if err != nil {
        log.Fatalf("[-] Error creating nmap scanner:\n%v", err)
    }


  //if err != nil {
  //  log.Fatalf("[-] Error creating nmap scann:\n%v", err)
  //}

  result, warnings, err := scanner.Run()
  
  if len(*warnings) > 0 {
    log.Printf("[i] run finished with warnings:\n%s", *warnings)
  }

  if err != nil {
    log.Fatalf("[-] Error executing nmap scann:\n%v", err)
  }

  for _, host := range result.Hosts {
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      continue
    }
    
    fmt.Printf("[+] Host %q:\n", host.Addresses[0])

    for _,port := range host.Ports {
      fmt.Printf("\t[%s] Port %d/%s %s\n", port.State, port.ID, port.Protocol, port.Service.Name)
    }

  }

  fmt.Printf("[+] Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

}

// TODO add List
// Example to Nmap pharams: Nmap("127.0.0.0", "1-10000")
func Nmap(TARGET string, Flags []string) {

  fmt.Println("[i] Starting Scann ")

  scann(TARGET, Flags)

  fmt.Println("Nmap")
}
