package director

import (
	"errors"
	"server/helpers"
	"strconv"

	structDB "server/structs/db"
	structLogic "server/structs/logic"

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
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		helpers.CheckErr("Error begin @ApproveByDirector", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	getDirector, errGetDirector := DBUser.GetDirector()
	helpers.CheckErr("Error get director @ApproveByDirecto - logicDirectorr", errGetDirector)

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @ApproveByDirector - logicDirector", errGetEmployee)

	getLeave, errGetLeave := DBLeave.GetLeave(id)
	helpers.CheckErr("Error get leave @ApproveByDirector", errGetLeave)

	resGet, errGet := DBUser.GetUserLeaveRemaining(getLeave.TypeLeaveID, employeeNumber)
	helpers.CheckErr("Error get leave remaining @ApproveByDirector - logicDirector", errGet)

	strTotal := strconv.FormatFloat(getLeave.Total, 'f', 1, 64)
	strBalance := strconv.FormatFloat(resGet.LeaveRemaining, 'f', 1, 64)

	if getLeave.Total > float64(resGet.LeaveRemaining) {
		beego.Warning("Error leave balance @ApproveByDirector - logicDirector")
		return errors.New("Employee total leave is " + strTotal + " day and employee " + resGet.TypeName + " balance is " + strBalance + " day left")
	}

	actionBy := getDirector.Name

	errApprove := DBDirector.ApproveByDirector(id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @ApproveByDirector - logicDirector", errApprove)
		o.Rollback()
		return errApprove
	}

	errUp := DBLeave.UpdateLeaveRemaningApprove(getLeave.Total, employeeNumber, getLeave.TypeLeaveID)
	if errUp != nil {
		helpers.CheckErr("Error update leave balance @ApproveByDirector - logicDirector", errUp)
		o.Rollback()
		return errUp
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

	return err
}

// RejectByDirector ...
func RejectByDirector(l *structDB.LeaveRequest, id int64, employeeNumber int64) error {

	getDirector, errGetDirector := DBUser.GetDirector()
	helpers.CheckErr("Error get director @RejectByDirector - logicDirector", errGetDirector)

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @RejectByDirector - logicDirector", errGetEmployee)

	getLeave, errGetLeave := DBLeave.GetLeave(id)
	helpers.CheckErr("Error get leave @RejectByDirector - logicDirector", errGetLeave)

	rejectReason := l.RejectReason
	actionBy := getDirector.Name

	errApprove := DBDirector.RejectByDirector(l, id, employeeNumber, actionBy)
	if errApprove != nil {
		helpers.CheckErr("Error approved request @RejectByDirector - logicDirector", errApprove)
		return errApprove
	}

	go func() {
		helpers.GoMailDirectorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getDirector.Name, rejectReason)
	}()

	return errApprove
}
