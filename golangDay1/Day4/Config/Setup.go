package Config

type Job struct {
	Id      string   `json:"id"` //will save order Id
	Attempt int `json:"attempt"` //will save number of attempt
}
var MainChannel chan Job //will save order id of orders
const CooldownPeriod = 60


