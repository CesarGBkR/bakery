package User

import (
  "flag"

  "bakery/Application"
)

func User() {
 
  var (
    FILE string
    IP string
    NS string
    PORTS string
    RATE int
    TYPE string    
  )

  flag.StringVar(&FILE, "F", "", "Tatget File Route")
  flag.StringVar(&IP, "IP", "127.0.0.1", "Target IP")
  flag.StringVar(&NS, "NS", "localhost.com", "Target Name Server")
  flag.StringVar(&PORTS, "P", "1-65535", "Target Ports")
  flag.IntVar(&RATE, "R", 5000, "Min RATE")
  flag.StringVar(&TYPE, "T", "P", "Action To Perform")
  flag.Parse()
  
  Application.ApplicationMain(FILE, IP, NS, PORTS, RATE, TYPE)
}
