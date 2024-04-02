package User

import (
  //"fmt"
  
  "bakery/Application"
)

func User() {

  // JSON FORMAT 
  
  jsonTest := `[{
    "IP": "10.10.11.8",
    "NS": "localhost.com",
    "LVL": 3,
    "RATE": 5000
  }]`

  //jsonTest := `[{
    //"IP": "127.0.0.0",
    //"NS": "localhost.com",
    //"LVL": 1,
    //"RATE": 5000
  //},
  //{
    //"IP": "127.0.0.0",
    //"NS": "localhost.com",
    //"LVL": 1,
    //"RATE": 5000
  //}]` 
  Application.Application(jsonTest)
}
