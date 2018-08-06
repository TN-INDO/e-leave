package logic

// GetSupervisorID ...
type GetSupervisorID struct {
	SupervisorID int64 `json:"supervisor_id" orm:"column(supervisor_id)"`
}

// GetSupervisor ...
type GetSupervisor struct {
	SupervisorID int64  `json:"supervisor_id" orm:"column(supervisor_id)"`
	Name         string `json:"name" orm:"column(name)"`
	Email        string `json:"email" orm:"column(email)"`
}

// GetSupervisors ...
type GetSupervisors struct {
	SupervisorID int64  `json:"supervisor_id" orm:"column(employee_number)"`
	Name         string `json:"name" orm:"column(name)"`
}
