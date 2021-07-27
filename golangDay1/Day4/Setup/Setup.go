package Setup

import (
	"retailer-api/Config"
	"retailer-api/Models/Order"
	"sync"
)

func Initiate(){
	Config.MainChannel =make(chan Config.Job,1000)

	//mutex will be used to stop overlapping if multiple customer are trying to by the same product
	mutex:=sync.Mutex{}
	go func(){
		for{
			select {
			case job:=<-Config.MainChannel:
				Order.ExecuteOrder(job,&mutex)
			}
		}
	}()
}
