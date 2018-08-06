package logic

// LeaveReject ...
type LeaveReject struct {
	ID               int64   `json:"id" orm:"column(id);pk"`
	EmployeeNumber   int64   `json:"employee_number" orm:"column(employee_number);pk"`
	Name             string  `json:"name" orm:"column(name)"`
	Gender           string  `json:"gender" orm:"column(gender)"`
	Position         string  `json:"position" orm:"column(position)"`
	StartWorkingDate string  `json:"start_working_date" orm:"column(start_working_date)"`
	Email            string  `json:"email" orm:"column(email)"`
	Role             string  `json:"role" orm:"column(role)"`
	TypeName         string  `json:"type_name" orm:"column(type_name)"`
	LeaveRemaining   float64 `json:"leave_remaining" orm:"column(leave_remaining)"`
	Reason           string  `json:"reason" orm:"column(reason)"`
	DateFrom         string  `json:"date_from" orm:"column(date_from)"`
	DateTo           string  `json:"date_to" orm:"column(date_to)"`
	HalfDates        string  `json:"half_dates" orm:"column(half_dates)"`
	BackOn           string  `json:"back_on" orm:"column(back_on)"`
	Total            float64 `json:"total" orm:"column(total)"`
	ContactAddress   string  `json:"contact_address" orm:"column(contact_address)"`
	ContactNumber    string  `json:"contact_number" orm:"column(contact_number)"`
	Status           string  `json:"status" orm:"column(status)"`
	ActionBy         string  `json:"action_by" orm:"column(action_by)"`
	RejectReason     string  `json:"reject_reason" orm:"column(reject_reason)"`
}
