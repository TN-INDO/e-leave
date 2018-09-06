package leave

import (
	"server/helpers"

	structAPI "server/structs/api"
	structLogic "server/structs/logic"
)

// GetLeave ...
func GetLeave(id int64) (structLogic.GetLeave, error) {
	respGet, errGet := DBLeave.GetLeave(id)
	if errGet != nil {
		helpers.CheckErr("Error get leave request @GetLeave - logicLeave", errGet)
	}

	return respGet, errGet
}

// DeleteRequest ...
func DeleteRequest(id int64) (err error) {
	errDelete := DBLeave.DeleteRequest(id)
	if errDelete != nil {
		helpers.CheckErr("Error delete leave request @DeleteRequest - logicLeave", errDelete)
	}

	return errDelete
}

// UpdateLeaveRemaningApprove ...
func UpdateLeaveRemaningApprove(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningApprove(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningApprove - logicLeave", errUpdate)
	}

	return errUpdate
}

// UpdateLeaveRemaningCancel ...
func UpdateLeaveRemaningCancel(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningCancel(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningCancel - logicLeave", errUpdate)
	}

	return errUpdate
}

// DownloadReportCSV ...
func DownloadReportCSV(query *structAPI.RequestReport, path string) error {
	errGet := DBLeave.DownloadReportCSV(query.FromDate, query.ToDate, path)
	if errGet != nil {
		helpers.CheckErr("Error get report @DownloadReportCSV - logicLeave", errGet)
	}

	return errGet
}

// ReportLeaveRequest ...
func ReportLeaveRequest(query *structAPI.RequestReport) (report []structLogic.ReportLeaveRequest, err error) {
	respGet, errGet := DBLeave.ReportLeaveRequest(query.FromDate, query.ToDate)
	if errGet != nil {
		helpers.CheckErr("Error get report @ReportLeaveRequest - logicLeave", errGet)
	}

	return respGet, errGet
}

// ReportLeaveRequestTypeLeave ...
func ReportLeaveRequestTypeLeave(query *structAPI.RequestReportTypeLeave) (report []structLogic.ReportLeaveRequest, err error) {
	respGet, errGet := DBLeave.ReportLeaveRequestTypeLeave(query.FromDate, query.ToDate, query.TypeLeaveID)
	if errGet != nil {
		helpers.CheckErr("Error get report type leave @ReportLeaveRequestTypeLeave - logicLeave", errGet)
	}

	return respGet, errGet
}
