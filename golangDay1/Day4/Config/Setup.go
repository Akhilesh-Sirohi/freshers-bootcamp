package Config

/*
Job will be used to provide order Id and attempt number to the worker
   Worker is ExecuteOrder function in Models/Order.go
Id of Job will represent orderId
Attempt of Job will represent number of times worker has tried to complete the order
 */
type Job struct {
	Id      string   `json:"id"` //will save order Id
	Attempt int `json:"attempt"` //will save number of attempt
}

/*
MainChannel will be a channel of Jobs. It is used in order to achieve concurrency
 */
var MainChannel chan Job //will save order id of orders


/*
CooldownPeriod is time in second between executing consecutive orders of same customer
 */
const CooldownPeriod = 120



