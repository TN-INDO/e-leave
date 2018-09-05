package supervisor

import (
	"server/helpers"
	logicLeave "server/models/db/pgsql/leave_request"
	logicUser "server/models/db/pgsql/user"
	structLogic "server/structs/logic"
)

// GetEmployeePending ...
func GetEmployeePending(supervisorID int64) ([]structLogic.LeavePending, error) {
	respGet, errGet := DBSupervisor.GetEmployeePending(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePending - logic", errGet)
	}

	return respGet, errGet
}

// GetEmployeeApproved ...
func GetEmployeeApproved(supervisorID int64) ([]structLogic.LeaveAccept, error) {
	respGet, errGet := DBSupervisor.GetEmployeeApproved(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get approved request @GetEmployeeApproved - logic", errGet)
	}

	return respGet, errGet
}

// GetEmployeeRejected ...
func GetEmployeeRejected(supervisorID int64) ([]structLogic.LeaveReject, error) {
	respGet, errGet := DBSupervisor.GetEmployeeRejected(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get rejected request @GetEmployeeRejected - logic", errGet)
	}

	return respGet, errGet
}

// ApproveBySupervisor ...
func ApproveBySupervisor(id int64, employeeNumber int64) error {
	var (
		user  logicUser.User
		leave logicLeave.LeaveRequest
	)

	getEmployee, errGetEmployee := user.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @ApproveBySupervisor", errGetEmployee)

	getSupervisorID, errGetSupervisorID := user.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @ApproveBySupervisor", errGetSupervisorID)

	getSupervisor, errGetSupervisor := user.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @ApproveBySupervisor", errGetSupervisor)

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get director @ApproveBySupervisor", errGetDirector)

	getLeave, errGetLeave := leave.GetLeave(id)
	helpers.CheckErr("Error get leave @ApproveBySupervisor", errGetLeave)

	actionBy := getSupervisor.Name

	errApprove := DBSupervisor.ApproveBySupervisor(id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @ApproveBySupervisor - logic", errApprove)
	}

	go func() {
		helpers.GoMailEmployee(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name)
		helpers.GoMailDirector(getDirector.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, getDirector.Name)
	}()

	return errApprove
}

// RejectBySupervisor ...
func RejectBySupervisor(l *structLogic.LeaveReason, id int64, employeeNumber int64) error {
	var (
		user  logicUser.User
		leave logicLeave.LeaveRequest
	)

	getEmployee, errGetEmployee := user.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @RejectBySupervisor", errGetEmployee)

	getSupervisorID, errGetSupervisorID := user.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @RejectBySupervisor", errGetSupervisorID)

	getSupervisor, errGetSupervisor := user.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @RejectBySupervisor", errGetSupervisor)

	getLeave, errGetLeave := leave.GetLeave(id)
	helpers.CheckErr("Error get leave @RejectBySupervisor", errGetLeave)

	rejectReason := l.RejectReason
	actionBy := getSupervisor.Name

	errReject := DBSupervisor.RejectBySupervisor(l, id, employeeNumber, actionBy)
	if errReject != nil {
		helpers.CheckErr("Error rejected request @RejectBySupervisor - logic", errReject)
	}

	go func() {
		helpers.GoMailSupervisorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, rejectReason)
	}()

	return errReject
}
