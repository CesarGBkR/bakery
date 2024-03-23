package NmapObjects

type ScannResponse struct {
  Hosts Host
  Warn []string
  ErrBuild error
  ErrExec error
}

type ScannRequest struct {
  TARGET string 
  PORTS string 
  RATE int
}

type Host struct {
  Ports []Port
  OS string
}

type Port struct {
  State string
  ID uint16
  Protocol string
  Services Service
}

type Service struct {
  Name string
  Version string
  Product string
  Extra string
}
