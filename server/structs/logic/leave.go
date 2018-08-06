package logic

// GetLeave ...
type GetLeave struct {
	ID          string  `json:"id" orm:"column(id)"`
	TypeLeaveID int64   `json:"type_leave_id" orm:"column(type_leave_id)"`
	Total       float64 `json:"total" orm:"column(total)"`
}

// LeaveReason ...
type LeaveReason struct {
	// Status       string `json:"status" orm:"column(status)"`
	RejectReason string `json:"reject_reason" orm:"column(reject_reason)"`
}

// UserSumarry ...
type UserSumarry struct {
	TypeName string  `json:"type_name" orm:"column(type_name)"`
	Used     float64 `json:"used" orm:"column(used)"`
	// LeaveRemaining float64 `json:"leave_remaining" orm:"column(leave_remaining)"`
}

// UserTypeLeave ...
type UserTypeLeave struct {
	TypeName       string  `json:"type_name" orm:"column(type_name)"`
	LeaveRemaining float64 `json:"leave_remaining" orm:"column(leave_remaining)"`
}
