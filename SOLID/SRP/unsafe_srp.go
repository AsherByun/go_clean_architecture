package main

import "fmt"

type Employee struct {
	WorkTime int
	Name     string
	Payment  int
}

func NewEmployee(name string, work_time int) Employee {
	return Employee{Name: name, WorkTime: work_time}
}

type EmployeeMethod interface {
	calculatePay()
	save()
	reportHours()
}

func (e *Employee) calculatePay() {
	e.Payment = regularHours(e) * 100000
	fmt.Println(e.Payment)
}

func (e *Employee) save() {
	fmt.Println(e.Name + "   saved")
}

func (e *Employee) reportHours() {
	fmt.Println(regularHours(e))
}

func regularHours(e *Employee) int {
	return e.WorkTime - 3
}

func main() {
	e := NewEmployee("sungjae", 100)

	e.calculatePay()
	e.save()
}

// ---> reqularHours 라는 함수를 reportHours, calculatePay 두 메서드에서 사용
// 이때 사용하는 주체가 다름(cto, coo) 따라서 한쪽에서 변경시 다른쪽도 변경해야함 --> srp 원칙 위배
// Table 스키마를 한 주체가 변경시 다른 주체에 영향을 줄 수 있음
