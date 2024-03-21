package User

import (
  //"fmt"
  
  "bakery/Application"
)

func User() {

  // JSON FORMAT 
  
  jsonTest := `[{
    "IP": "127.0.0.0",
    "NS": "localhost.com",
    "LVL": 1
  },
  {
    "IP": "127.0.0.0",
    "NS": "localhost.com",
    "LVL": 1
  }]` 
  Application.Application(jsonTest)
}
