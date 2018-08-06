package db

type (
	// LeaveRequest ...
	LeaveRequest struct {
		ID             int64  `json:"id" orm:"column(id);pk"`
		EmployeeNumber int64  `json:"employee_number" orm:"column(employee_number)"`
		TypeLeaveID    int64  `json:"type_leave_id" orm:"column(type_leave_id)"`
		Reason         string `json:"reason" orm:"column(reason)"`
		DateFrom       string `json:"date_from" orm:"column(date_from)"`
		DateTo         string `json:"date_to" orm:"column(date_to)"`
		// HalfDates      []string `json:"half_dates" orm:"column(half_dates)"`
		Total          float64 `json:"total" orm:"column(total)"`
		BackOn         string  `json:"back_on" orm:"column(back_on)"`
		ContactAddress string  `json:"contact_address" orm:"column(contact_address)"`
		ContactNumber  string  `json:"contact_number" orm:"column(contact_number)"`
		Status         string  `json:"status" orm:"column(status)"`
		ActionBy       string  `json:"action_by" orm:"column(action_by)"`
		RejectReason   string  `json:"reject_reason" orm:"column(reject_reason)"`
	}
)

// TableName ...
func (u *LeaveRequest) TableName() string {
	return "leave_request"
}
