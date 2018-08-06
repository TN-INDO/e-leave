package api

// ReqLeave ...
type (
	ReqLeave struct {
		TypeLeaveID    int64    `json:"type_leave_id" orm:"column(type_leave_id)"`
		Reason         string   `json:"reason" orm:"column(reason)"`
		DateFrom       string   `json:"date_from" orm:"column(date_from)"`
		DateTo         string   `json:"date_to" orm:"column(date_to)"`
		HalfDates      []string `json:"half_dates" orm:"column(half_dates)"`
		BackOn         string   `json:"back_on" orm:"column(back_on)"`
		ContactAddress string   `json:"contact_address" orm:"column(contact_address)"`
		ContactNumber  string   `json:"contact_number" orm:"column(contact_number)"`
	}

	UpdateLeaveRequest struct {
		TypeLeaveID    int64    `json:"type_leave_id" orm:"column(type_leave_id)"`
		Reason         string   `json:"reason" orm:"column(reason)"`
		DateFrom       string   `json:"date_from" orm:"column(date_from)"`
		DateTo         string   `json:"date_to" orm:"column(date_to)"`
		HalfDates      []string `json:"half_dates" orm:"column(half_dates)"`
		Total          float64  `json:"total" orm:"column(total)"`
		BackOn         string   `json:"back_on" orm:"column(back_on)"`
		ContactAddress string   `json:"contact_address" orm:"column(contact_address)"`
		ContactNumber  string   `json:"contact_number" orm:"column(contact_number)"`
	}
)
