package main

type PayCalculator interface {
	calculatePay()
}

type Hourrepoter interface {
	reportHours()
}

type Employeesaver interface {
	saveEmployee()
}

// --> 각 주체별로 클래스를 생성하여 사용함으로써 srp를 만족시킨다
