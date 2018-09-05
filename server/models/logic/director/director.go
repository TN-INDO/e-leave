package director

import (
	"errors"
	"server/helpers"
	logicAdmin "server/models/db/pgsql/admin"
	logicLeave "server/models/db/pgsql/leave_request"
	logicUser "server/models/db/pgsql/user"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetEmployeePendingRequest ...
func GetEmployeePendingRequest() ([]structLogic.RequestPending, error) {
	respGet, errGet := DBDirector.GetEmployeePending()
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// GetEmployeeApprovedRequest ...
func GetEmployeeApprovedRequest() ([]structLogic.RequestAccept, error) {
	respGet, errGet := DBDirector.GetEmployeeApproved()
	if errGet != nil {
		helpers.CheckErr("Error get approved request @GetEmployeeApprovedRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// GetEmployeeRejectedRequest ...
func GetEmployeeRejectedRequest() ([]structLogic.RequestReject, error) {
	respGet, errGet := DBDirector.GetEmployeeRejected()
	if errGet != nil {
		helpers.CheckErr("Error get rejected request @GetEmployeeRejectedRequest - logicDirector", errGet)
	}

	return respGet, errGet
}

// ApproveByDirector ...
func ApproveByDirector(id int64, employeeNumber int64) error {
	var (
		user  logicUser.User
		leave logicLeave.LeaveRequest
		admin logicAdmin.Admin
	)

	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		helpers.CheckErr("Error begin @ApproveByDirector", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get director @ApproveByDirecto - logicDirectorr", errGetDirector)

	getEmployee, errGetEmployee := user.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @ApproveByDirector - logicDirector", errGetEmployee)

	getLeave, errGetLeave := leave.GetLeave(id)
	helpers.CheckErr("Error get leave @ApproveByDirector", errGetLeave)

	resGet, errGet := user.GetUserLeaveRemaining(getLeave.TypeLeaveID, employeeNumber)
	helpers.CheckErr("Error get leave remaining @ApproveByDirector - logicDirector", errGet)

	strTotal := strconv.FormatFloat(getLeave.Total, 'f', 1, 64)
	strBalance := strconv.FormatFloat(resGet.LeaveRemaining, 'f', 1, 64)

	if getLeave.Total > float64(resGet.LeaveRemaining) {
		beego.Warning("error leave balance @ApproveByDirector - logicDirector")
		return errors.New("Employee total leave is " + strTotal + " day and employee " + resGet.TypeName + " balance is " + strBalance + " day left")
	}

	actionBy := getDirector.Name

	errApprove := DBDirector.ApproveByDirector(id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @ApproveByDirector - logicDirector", errApprove)
		o.Rollback()
	}

	errUp := admin.UpdateLeaveRemaning(getLeave.Total, employeeNumber, getLeave.TypeLeaveID)
	if errUp != nil {
		helpers.CheckErr("Error update leave balance @ApproveByDirector - logicDirector", errUp)
		o.Rollback()
	}

	err = o.Commit()
	if err != nil {
		helpers.CheckErr("Error commit @ApproveByDirector - logicDirector", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	go func() {
		helpers.GoMailDirectorAccept(getEmployee.Email, getLeave.ID, getEmployee.Name, getDirector.Name)

	}()

	return errApprove
}

// RejectByDirector ...
func RejectByDirector(l *structDB.LeaveRequest, id int64, employeeNumber int64) error {
	var (
		user  logicUser.User
		leave logicLeave.LeaveRequest
	)

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get director @RejectByDirector - logicDirector", errGetDirector)

	getEmployee, errGetEmployee := user.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @RejectByDirector - logicDirector", errGetEmployee)

	getLeave, errGetLeave := leave.GetLeave(id)
	helpers.CheckErr("Error get leave @RejectByDirector - logicDirector", errGetLeave)

	rejectReason := l.RejectReason
	actionBy := getDirector.Name

	errApprove := DBDirector.RejectByDirector(l, id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @RejectByDirector - logicDirector", errApprove)
	}

	go func() {
		helpers.GoMailDirectorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getDirector.Name, rejectReason)
	}()

	return errApprove
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
