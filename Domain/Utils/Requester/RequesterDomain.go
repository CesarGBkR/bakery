package RequesterDomain

import (
	"bakery/Domain/Object"
	"bakery/Domain/Object/RequesterObjects"
	"bakery/Service/Utils/Requester"
)

func Requester(url string) RequesterObjects.Response {
	return requester.Requester(url)
}

func Post(TARGET objects.TargetObject) {
	requester.Post()
}

func Put(TARGET objects.TargetObject) {
	requester.Put()
}

func Delete(TARGET objects.TargetObject) {
	requester.Delete()
}
