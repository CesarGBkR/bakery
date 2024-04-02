package Application

import (
  "fmt"
  "strings"
  "encoding/json"
  "text/tabwriter"
  
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

func ServiceScann(TARGET string, PORTS string, RATE int){
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ServiceScann(TARGET, PORTS, RATE) 
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\tS.Name\tS.Product\tS.Version\tS.Extra\n")
  for _, port := range response.NmapResponse.Hosts.Ports {

    fmt.Fprintf(w,"%s\t%d\t%s\t%s\t%s\t%s\t%s\n", port.State, port.ID, port.Protocol, port.Services.Name, port.Services.Product, port.Services.Version, port.Services.Extra)
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
    fmt.Println("[-] Error on JSON deserialization:\n%v", err)
  }

  return objectList
  
}

  // LEVEL MNGMNT

func LvlM(Target objects.TargetObject) {
  switch Target.LVL {
  case 1:
    RATE := 5000
    ScannAllPorts(Target.IP, RATE)
  case 2:
    RATE := 5000
    ScannPort(Target.IP, "1234", RATE)
  case 3:
    RATE := 5000
    ServiceScann(Target.IP,"22,5000", RATE )
  }
}



// Main function
// This function manage and invoice otter functions in the file

func Application(JSONLIST string) {
   
  var targetList []objects.TargetObject

  targetList = JsonToObject(JSONLIST)

  if len(targetList) > 0 {
    for _, target := range targetList {
      LvlM(target)
    }
  }
}
