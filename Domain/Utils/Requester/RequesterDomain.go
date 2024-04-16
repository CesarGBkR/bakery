package RequesterDomain

import (
  "bakery/Service/Utils/Requester"
  "bakery/Domain/Object"
)

func Get(TARGET objects.TargetObject) {
  requester.Get(TARGET)
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
