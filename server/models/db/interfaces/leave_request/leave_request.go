package leave

import (
	structAPI "server/structs/api"
	structLogic "server/structs/logic"
)

// IBaseLeaveRequest ...
type IBaseLeaveRequest interface {
	// CreateLeaveRequestEmployee
	CreateLeaveRequestEmployee(
		employeeNumber int64,
		typeLeaveID int64,
		reason string,
		dateFrom string,
		dateTo string,
		halfDates []string,
		backOn string,
		total float64,
		address string,
		contactLeave string,
		status string,
	) error
	// CreateLeaveRequestSupervisor
	CreateLeaveRequestSupervisor(
		employeeNumber int64,
		typeLeaveID int64,
		reason string,
		dateFrom string,
		dateTo string,
		halfDates []string,
		backOn string,
		total float64,
		address string,
		contactLeave string,
		status string,
	) error
	// UpdateRequest
	UpdateRequest(
		e *structAPI.UpdateLeaveRequest,
		id int64,
	) (err error)
	// DeleteRequest
	DeleteRequest(id int64) (err error)
	// GetLeave
	GetLeave(id int64) (
		result structLogic.GetLeave,
		err error,
	)

	// UpdateLeaveRemaningApprove
	UpdateLeaveRemaningApprove(
		total float64,
		employeeNumber int64,
		typeID int64,
	) (err error)
	// UpdateLeaveRemaningCancel ...
	UpdateLeaveRemaningCancel(
		total float64,
		employeeNumber int64,
		typeID int64,
	) (err error)

	// DownloadReportCSV
	DownloadReportCSV(
		fromDate string,
		toDate string,
		path string,
	) (err error)
	// WriteCsv
	WriteCsv(
		path string,
		res []structLogic.ReportLeaveRequest,
	) error

	// ReportLeaveRequest
	ReportLeaveRequest(
		fromDate string,
		toDate string,
	) (
		report []structLogic.ReportLeaveRequest,
		err error,
	)
	// ReportLeaveRequestTypeLeave
	ReportLeaveRequestTypeLeave(
		fromDate string,
		toDate string,
		typeLeaveID string,
	) (
		report []structLogic.ReportLeaveRequest,
		err error,
	)
}
