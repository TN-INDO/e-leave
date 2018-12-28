package supervisor

import (
	"server/helpers"

	structLogic "server/structs/logic"
)

// GetEmployeePending ...
func GetEmployeePending(supervisorID int64) ([]structLogic.LeavePending, error) {
	respGet, errGet := DBSupervisor.GetEmployeePending(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePending - logicSupervisor", errGet)
	}

	return respGet, errGet
}

// GetEmployeeApproved ...
func GetEmployeeApproved(supervisorID int64) ([]structLogic.LeaveAccept, error) {
	respGet, errGet := DBSupervisor.GetEmployeeApproved(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get approved request @GetEmployeeApproved - logicSupervisor", errGet)
	}

	return respGet, errGet
}

// GetEmployeeRejected ...
func GetEmployeeRejected(supervisorID int64) ([]structLogic.LeaveReject, error) {
	respGet, errGet := DBSupervisor.GetEmployeeRejected(supervisorID)
	if errGet != nil {
		helpers.CheckErr("Error get rejected request @GetEmployeeRejected - logicSupervisor", errGet)
	}

	return respGet, errGet
}

// ApproveBySupervisor ...
func ApproveBySupervisor(id int64, employeeNumber int64) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @ApproveBySupervisor - logicSupervisor", errGetEmployee)

	getSupervisorID, errGetSupervisorID := DBUser.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @ApproveBySupervisor - logicSupervisor", errGetSupervisorID)

	getSupervisor, errGetSupervisor := DBUser.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @ApproveBySupervisor - logicSupervisor", errGetSupervisor)

	getDirector, errGetDirector := DBUser.GetDirector()
	helpers.CheckErr("Error get director @ApproveBySupervisor - logicSupervisor", errGetDirector)

	getLeave, errGetLeave := DBLeave.GetLeave(id)
	helpers.CheckErr("Error get leave @ApproveBySupervisor - logicSupervisor", errGetLeave)

	actionBy := getSupervisor.Name

	errApprove := DBSupervisor.ApproveBySupervisor(id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @ApproveBySupervisor - logicSupervisor", errApprove)
		return errApprove
	}

	go func() {
		helpers.GoMailEmployee(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name)
		helpers.GoMailDirector(getDirector.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, getDirector.Name)
	}()

	return errApprove
}

// RejectBySupervisor ...
func RejectBySupervisor(l *structLogic.LeaveReason, id int64, employeeNumber int64) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @RejectBySupervisor - logicSupervisor", errGetEmployee)

	getSupervisorID, errGetSupervisorID := DBUser.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @RejectBySupervisor - logicSupervisor", errGetSupervisorID)

	getSupervisor, errGetSupervisor := DBUser.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @RejectBySupervisor - logicSupervisor", errGetSupervisor)

	getLeave, errGetLeave := DBLeave.GetLeave(id)
	helpers.CheckErr("Error get leave @RejectBySupervisor - logicSupervisor", errGetLeave)

	rejectReason := l.RejectReason
	actionBy := getSupervisor.Name

	errReject := DBSupervisor.RejectBySupervisor(l, id, employeeNumber, actionBy)
	if errReject != nil {
		helpers.CheckErr("Error rejected request @RejectBySupervisor - logicSupervisor", errReject)
		return errReject
	}

	go func() {
		helpers.GoMailSupervisorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, rejectReason)
	}()

	return errReject
}
