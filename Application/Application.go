package Application

import (
  "fmt"
  "strings"
  "encoding/json"
  "text/tabwriter"
  "io/ioutil"
  
  "bakery/Domain/Object"
  "bakery/Domain/Enumeration/PortScann"
)

// SCAN

func ApplicationScann(TARGET string) {

}

func ScannPort(TARGET string, PORTS string, RATE int){
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannPort(TARGET, PORTS, RATE) 
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\n")
  for _, port := range response.NmapResponse.Hosts.Ports {

    fmt.Fprintf(w,"%s\t%d\t%s\n", port.State, port.ID, port.Protocol)
  }
  w.Flush()
  fmt.Println(builder.String())
}

func ScannAllPorts(TARGET string, RATE int){
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannAllPorts(TARGET, RATE)
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\n")
  for _, port := range response.NmapResponse.Hosts.Ports {
    fmt.Fprintf(w, "%s\t%d\t%s\n", port.State, port.ID, port.Protocol)
  }
  w.Flush()
  fmt.Println(builder.String())
}

func ScannService(TARGET string, PORTS string, RATE int){
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannService(TARGET, PORTS, RATE) 
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\tS.Name\tS.Product\tS.Version\tS.Extra\n")
  for _, port := range response.NmapResponse.Hosts.Ports {
    fmt.Fprintf(w,"%s\t%d\t%s\t%s\t%s\t%s\t%s\n", port.State, port.ID, port.Protocol, port.Services.Name, port.Services.Product, port.Services.Version, port.Services.Extra)
  }
  w.Flush()
  fmt.Println(builder.String())
} 

func ScannScript(TARGET string, PORTS string, RATE int) {
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannScript(TARGET, PORTS, RATE)
  //fmt.Printf("\n%v",response.NmapResponse.Hosts.Ports)
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\tS.Name\tSC.ID\tSC.OUT\n")
  for _, port := range response.NmapResponse.Hosts.Ports {
    for _, script := range port.Scripts {
      fmt.Fprintf(w,"%s\t%d\t%s\t%s\t%s\t%s\n", port.State, port.ID, port.Protocol, port.Services.Name, script.ID,script.Output)
    }
  }
  w.Flush()
  fmt.Println(builder.String())

}

// FUZZING 
func ApplicationFuzzing() {
}

// UTILITIES
  // OBJECT CONVERSION

func JsonToObject(JSONLIST string) []objects.TargetObject {
  var objectList []objects.TargetObject
  err := json.Unmarshal([]byte(JSONLIST), &objectList)
  if err != nil { 
    fmt.Println("[-] Error on JSON deserialization:\n%s", err)
  }
  return objectList
}

  // TYPE MNGMNT

func TypeM(Target objects.TargetObject) {
  for _, TYPE := range Target.TYPE {
    TYPES := string(TYPE)
    switch TYPES {
      case "P":
        ScannScript(Target.IP, Target.PORTS, Target.RATE)
        //ScannService(Target.IP, Target.PORTS, Target.RATE)
      
      case "F":
        fmt.Printf("TODO")
        //ScannPort(Target.IP, "1234", RATE)
      }
    }
}

// Main function
// This function manage and invoice otter functions in the file

func ApplicationMain(FILE string, IP string, NS string, PORTS string, RATE int, TYPE string) {
  
  var TargetList []objects.TargetObject

  if FILE != "" {
    contenido, err := ioutil.ReadFile(FILE)
    if err != nil {
        fmt.Printf("[!] Error reading file: %v\n", err)
      }
    TargetList = JsonToObject(string(contenido))
  } else {
    Target := objects.TargetObject {
      IP: IP,
      NS: NS,
      PORTS: PORTS,
      TYPE: TYPE,
      RATE: RATE,
    }
    TargetList = append(TargetList, Target)
  }
   
  if len(TargetList) > 0 {
    for _, target := range TargetList {
      TypeM(target)
    }
  }
}
