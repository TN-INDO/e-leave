package api

// RequestReport ...
type RequestReport struct {
	FromDate string
	ToDate   string
}

// RequestReportTypeLeave ...
type RequestReportTypeLeave struct {
	FromDate    string
	ToDate      string
	TypeLeaveID string
}
