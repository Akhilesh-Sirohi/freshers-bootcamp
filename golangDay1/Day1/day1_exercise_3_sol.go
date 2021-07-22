package main

import (
	"fmt"
	"math/rand"
)

type Employee interface {
	Salary(timespan int) int
}
type fulltime struct {
	basic int
}
type contractor struct {
	basic int
}
type freelancer struct {
	basic int
}
func main() {
	var employee Employee
	fulltime_employee:=fulltime{500}
	employee=fulltime_employee;
	fmt.Println(employee.Salary(rand.Intn(200)))
	contractor_employee:=fulltime{100}
	employee=contractor_employee;
	fmt.Println(employee.Salary(rand.Intn(200)))
	freelancer_employee:=fulltime{10}
	employee=freelancer_employee;
	fmt.Println(employee.Salary(rand.Intn(28)))
	data:= [3]Employee{fulltime_employee,contractor_employee,freelancer_employee}
	for idx,employee:=range  data{
		//fmt.Println(idx)
		if idx==2{
			fmt.Println(employee.Salary(rand.Intn(200)))
		} else{
			fmt.Println(employee.Salary(rand.Intn(28)))
		}
	}

}
func (e fulltime) Salary(days int) int{
	return days*e.basic
}
func (e contractor) Salary(days int) int{
	return days*e.basic
}
func (e freelancer) Salary(hours int) int{
	if hours<20 {
		return 0
	}
	return hours*e.basic
}