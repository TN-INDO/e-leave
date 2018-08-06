package api

// ReqRegister ...
type ReqRegister struct {
	EmployeeNumber   int64  `json:"employee_number" orm:"column(employee_number);pk"`
	Name             string `json:"name" orm:"column(name)"`
	Gender           string `json:"gender" orm:"column(gender)"`
	Position         string `json:"position" orm:"column(position)"`
	StartWorkingDate string `json:"start_working_date" orm:"column(start_working_date)"`
	MobilePhone      string `json:"mobile_phone" orm:"column(mobile_phone)"`
	Email            string `json:"email" orm:"column(email)"`
	Password         string `json:"password" orm:"column(password)"`
	Role             string `json:"role" orm:"column(role)"`
	SupervisorID     int64  `json:"supervisor_id" orm:"column(supervisor_id)"`
}
