package leave

import (
	"server/helpers"
	"server/models/logic/user"

	structAPI "server/structs/api"
	structLogic "server/structs/logic"
)

// CreateLeaveRequestEmployee ...
func CreateLeaveRequestEmployee(
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
) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @CreateLeaveRequestEmployee", errGetEmployee)

	getSupervisorID, errGetSupervisorID := DBUser.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @CreateLeaveRequestEmployee", errGetSupervisorID)

	getSupervisor, errGetSupervisor := DBUser.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @CreateLeaveRequestEmployee", errGetSupervisor)

	errInsert := DBLeave.CreateLeaveRequestEmployee(employeeNumber, typeLeaveID, reason, dateFrom, dateTo, halfDates, backOn, total, address, contactLeave, status)
	if errInsert != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestEmployee - logicLeave", errInsert)
		return errInsert
	}

	go func() {
		helpers.GoMailSupervisor(getSupervisor.Email, getEmployee.Name, getSupervisor.Name)
	}()

	return errInsert
}

// CreateLeaveRequestSupervisor ...
func CreateLeaveRequestSupervisor(
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
) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @CreateLeaveRequestSupervisor", errGetEmployee)

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get employee @CreateLeaveRequestSupervisor", errGetDirector)

	errInsert := DBLeave.CreateLeaveRequestSupervisor(employeeNumber, typeLeaveID, reason, dateFrom, dateTo, halfDates, backOn, total, address, contactLeave, status)
	if errInsert != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestSupervisor - logicLeave", errInsert)
		return errInsert
	}

	go func() {
		helpers.GoMailDirectorFromSupervisor(getDirector.Email, getEmployee.Name, getDirector.Name)
	}()

	return errInsert
}

// // UpdateRequest ...
// func UpdateRequest(e *structAPI.UpdateLeaveRequest) error {
// 	errUpdate := DBLeave.UpdateRequest(e)
// 	if errUpdate != nil {
// 		helpers.CheckErr("Error update leave request @UpdateRequest - logicLeave", errUpdate)
// 	}

// 	return errUpdate
// }

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
		return errDelete
	}

	return errDelete
}

// UpdateLeaveRemaningApprove ...
func UpdateLeaveRemaningApprove(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningApprove(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningApprove - logicLeave", errUpdate)
		return errUpdate
	}

	return errUpdate
}

// UpdateLeaveRemaningCancel ...
func UpdateLeaveRemaningCancel(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningCancel(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningCancel - logicLeave", errUpdate)
		return errUpdate
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
