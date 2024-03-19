package Application

import (
  "fmt"
  "encoding/json"

  "bakery/Application/NmapApplication"
  "bakery/Domain/Object"
  )

// SCAN
func ApplicationScann(TARGET string) {
    NmapApplication.AllPorts(Target)

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
    ApplicationScann(Target.IP)
  }
}



// Main function
// This function manage and invoice otter functions in the file

func Application(JSONLIST string) {
  
  fmt.Println("[i] Application is connected")
  
  var targetList []objects.TargetObject

  targetList = JsonToObject(JSONLIST)

  if len(targetList) > 0 {
    for _, target := range targetList {
      LvlM(target)
    }
  }
}
