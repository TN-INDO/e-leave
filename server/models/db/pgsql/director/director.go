package director

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Director ...
type Director struct{}

// GetEmployeePending ...
func (u *Director) GetEmployeePending() (reqPending []structLogic.RequestPending, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetEmployeePending", errQB)
		return reqPending, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statPendingDirector := constant.StatusPendingInDirector

	count, errRaw := o.Raw(sql, statPendingDirector).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeePending", errRaw)
		return reqPending, errors.New("Error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetEmployeeApproved ...
func (u *Director) GetEmployeeApproved() (reqApprove []structLogic.RequestAccept, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetEmployeeApproved", errQB)
		return reqApprove, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statApproveDirector := constant.StatusSuccessInDirector

	count, errRaw := o.Raw(sql, statApproveDirector).QueryRows(&reqApprove)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeeApproved", errRaw)
		return reqApprove, errors.New("Error get leave request approved")
	}
	beego.Debug("Total approve request =", count)

	return reqApprove, errRaw
}

// GetEmployeeRejected ...
func (u *Director) GetEmployeeRejected() (reqReject []structLogic.RequestReject, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetEmployeeRejected", errQB)
		return reqReject, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".reject_reason",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	StatRejectInDirector := constant.StatusRejectInDirector

	count, errRaw := o.Raw(sql, StatRejectInDirector).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeeRejected", errRaw)
		return reqReject, errors.New("Error get leave request rejected")
	}
	beego.Debug("Total reject request =", count)

	return reqReject, errRaw
}

// ApproveByDirector ...
func (u *Director) ApproveByDirector(id int64, employeeNumber int64, actionBy string) (err error) {
	var dbLeave structDB.LeaveRequest

	o := orm.NewOrm()

	statAcceptDirector := constant.StatusSuccessInDirector

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statAcceptDirector, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("Error update status success @ApproveByDirector", errRAW)
	}

	return errRAW
}

// RejectByDirector ...
func (u *Director) RejectByDirector(l *structDB.LeaveRequest, id int64, employeeNumber int64, actionBy string) error {
	var dbLeave structDB.LeaveRequest

	o := orm.NewOrm()

	statRejectDirector := constant.StatusRejectInDirector
	rejectReason := l.RejectReason

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, reject_reason = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statRejectDirector, rejectReason, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("error update status reject @RejectByDirector", errRAW)
	}

	return errRAW
}
