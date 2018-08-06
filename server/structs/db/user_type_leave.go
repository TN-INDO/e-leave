package db

// UserTypeLeave ...
type UserTypeLeave struct {
	ID             int64   `json:"id" orm:"column(id);pk"`
	EmployeeNumber int64   `json:"employee_number" orm:"column(employee_number)"`
	TypeLeaveID    int64   `json:"type_leave_id" orm:"column(type_leave_id)"`
	LeaveRemaining float64 `json:"leave_remaining" orm:"column(leave_remaining)"`
}

// TableName ...
func (u *UserTypeLeave) TableName() string {
	return "user_type_leave"
}
