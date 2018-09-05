package admin

import (
	"errors"
	"server/helpers"
	logicLeave "server/models/db/pgsql/leave_request"
	logicUser "server/models/db/pgsql/user"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego/orm"
)

// GetEmployeePendingRequest ...
func GetEmployeePendingRequest() ([]structLogic.RequestPending, error) {
	respGet, errGet := DBAdmin.GetEmployeePending()
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// GetEmployeeApprovedRequest ...
func GetEmployeeApprovedRequest() ([]structLogic.RequestAccept, error) {
	respGet, errGet := DBAdmin.GetEmployeeApproved()
	if errGet != nil {
		helpers.CheckErr("Error get approved request @GetEmployeeApprovedRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// GetEmployeeRejectedRequest ...
func GetEmployeeRejectedRequest() ([]structLogic.RequestReject, error) {
	respGet, errGet := DBAdmin.GetEmployeeRejected()
	if errGet != nil {
		helpers.CheckErr("Error get rejected request @GetEmployeeRejectedRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// CancelRequestLeave ...
func CancelRequestLeave(id int64, employeeNumber int64) (err error) {
	var (
		user  logicUser.User
		leave logicLeave.LeaveRequest
	)

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		helpers.CheckErr("Error begin @CancelRequestLeave - logicDirector", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get director @CancelRequestLeave - logicDirector", errGetDirector)

	getEmployee, errGetEmployee := user.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @CancelRequestLeave - logicDirector", errGetEmployee)

	getLeave, errGetLeave := leave.GetLeave(id)
	helpers.CheckErr("Error get leave @CancelRequestLeave - logicDirector", errGetLeave)

	errUp := leave.UpdateLeaveRemaningCancel(getLeave.Total, employeeNumber, getLeave.TypeLeaveID)
	if errUp != nil {
		helpers.CheckErr("Error update cancel leave request @CancelRequestLeave - logicDirector", errUp)
		o.Rollback()
	}

	errDelete := leave.DeleteRequest(id)
	if errDelete != nil {
		helpers.CheckErr("Error delete leave request @CancelRequestLeave - logicDirector", errDelete)
		o.Rollback()
	}

	err = o.Commit()
	if err != nil {
		helpers.CheckErr("Error commit @CancelRequestLeave - logicDirector", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	go func() {
		helpers.GoMailDirectorCancel(getDirector.Email, getLeave.ID, getEmployee.Name, getDirector.Name)
		helpers.GoMailEmployeeCancel(getEmployee.Email, getLeave.ID, getEmployee.Name)
	}()

	return err
}
