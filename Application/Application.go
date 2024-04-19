package Application

import (
  "fmt"
  "strings"
  "encoding/json"
  "text/tabwriter"
  "io/ioutil"
  "sync"
  "reflect"
  
  "bakery/Domain/Object"
  "bakery/Domain/Enumeration/PortScann"
  "bakery/Domain/Utils"
  "bakery/Domain/Object/RequesterObjects"
)

// SCAN

func ApplicationScann(TARGET string) {

}

func ScannPort(wg *sync.WaitGroup, TARGET objects.TargetObject){
  defer wg.Done()
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannPort(TARGET) 
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\n")
  for _, port := range response.NmapResponse.Hosts.Ports {

    fmt.Fprintf(w,"%s\t%d\t%s\n", port.State, port.ID, port.Protocol)
  }
  w.Flush()
  fmt.Println(builder.String())
}


// DEPRECATED Not function yet
func ScannAllPorts(wg *sync.WaitGroup, TARGET objects.TargetObject){
  defer wg.Done()
  var builder strings.Builder
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  response := PortScann.ScannAllPorts(TARGET)
  fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\n")
  for _, port := range response.NmapResponse.Hosts.Ports {
    fmt.Fprintf(w, "%s\t%d\t%s\n", port.State, port.ID, port.Protocol)
  }
  w.Flush()
  fmt.Println(builder.String())
}

func ScannService(id int, wg *sync.WaitGroup, Enumeration chan<- objects.PortScannResponse, TARGET objects.TargetObject){
  defer wg.Done()
  response := PortScann.ScannService(TARGET) 
  Printer(objects.Response{
    Enumeration: objects.Enumeration{
      PortScann: response,
    },
  })  
  Enumeration <- response
} 

func ScannScript(id int, wg *sync.WaitGroup, Enumeration chan<- objects.PortScannResponse, TARGET objects.TargetObject) {
  defer wg.Done()
  response := PortScann.ScannScript(TARGET)
  Printer(objects.Response{
    Enumeration: objects.Enumeration{
      PortScann: response,
    },
  })
  Enumeration <- response
}

// FUZZING 
func ApplicationFuzzing() {
}

// UTILITIES
  // REQUESTER
func Requester(url string) RequesterObjects.Response{
  //defer wg.Done()
  return Utils.Requester(url)
}
  // Web Discover
func WebFinder(TARGET objects.TargetObject) {
  var PORTS []string 
  var Responses []RequesterObjects.Response
  //WebFinder := make(chan RequesterObjects.Response, 2)  
  //var wg sync.WaitGroup
  
  if len(TARGET.PORTS) < 0 || strings.Contains(TARGET.PORTS, "-"){
    //wg.Add(3)
    PORTS = append(PORTS, "80","8080","443") 
  }else{
    PORTS = strings.Split(TARGET.PORTS, ",")
    //wg.Add(len(PORTS))
  }
  for _, Port := range PORTS {

    urlIP := "http://"+TARGET.IP+":"+Port+"/"
    urlNS := "http://"+TARGET.NS+":"+Port+"/"
    fmt.Printf("\nRequest")
    Responses = append(Responses, Requester(urlIP))
    Responses = append(Responses, Requester(urlNS))
    fmt.Printf("\nRD")
  }
  //wg.Wait()
  fmt.Printf("\nWaited\n")
  //close(WebFinder)
  fmt.Printf("\nClosed\n")
  for _, Response := range Responses {
    fmt.Printf("\n%v", Response)

  }
}
  // PRINTER
func Printer(Response objects.Response) {
  var builder strings.Builder
  tipo := reflect.TypeOf(Response)
  w := tabwriter.NewWriter(&builder, 0, 0, 1, ' ', 0)
  
  for i := 0; i < tipo.NumField(); i++ {
    campo := tipo.Field(i)
    
    switch campo.Name {
      case "Enumeration":
        PORTS := Response.Enumeration.PortScann.NmapResponse.Hosts.Ports
        if len(PORTS) > 0 {
          if PORTS[0].Services.Version != "" {
            fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\tS.Name\tS.Product\tS.Version\tS.Extra\n")
            for _, port := range PORTS {


              fmt.Fprintf(w,"%s\t%d\t%s\t%s\t%s\t%s\t%s\n", port.State, port.ID, port.Protocol, port.Services.Name, port.Services.Product, port.Services.Version, port.Services.Extra)
            }
          }else {
            fmt.Fprintf(w,"STATE\tPORT\tPROTOCOL\tS.Name\tSC.ID\tSC.OUT\n")
            for _, port := range PORTS {
              for _, script := range port.Scripts {
                fmt.Fprintf(w,"%s\t%d\t%s\t%s\t%s\t%s\n", port.State, port.ID, port.Protocol, port.Services.Name, script.ID,script.Output)
              }
            }
          }
        
        
        w.Flush()
        fmt.Println(builder.String()) 
      }else {
        fmt.Printf("\n[-] No Open Ports\n")
      }
    }
  }
}

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
func TypeM(TARGET objects.TargetObject) {
  
  for _, TYPE := range TARGET.TYPE {
    TYPES := string(TYPE)
    switch TYPES {
      case "P":
        PortEnum := make(chan objects.PortScannResponse, 2)
        
        var wg sync.WaitGroup
        wg.Add(2)
        
        go ScannService(2, &wg, PortEnum, TARGET)
        go ScannScript(1, &wg, PortEnum, TARGET)
        
        wg.Wait()
        close(PortEnum)

      case "F":
        fmt.Printf("Testing")
        WebFinder(TARGET)
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
