package Utils

import (
	//"fmt"

	"bakery/Domain/Object/RequesterObjects"
	"bakery/Domain/Utils/Curl"
	"bakery/Domain/Utils/Requester"
)

func Requester(url string) RequesterObjects.Response {
	return RequesterDomain.Requester(url)
}

func Curl() {
	CurlDomain.Curl()
}
