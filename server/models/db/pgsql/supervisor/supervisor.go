package user

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Supervisor ...
type Supervisor struct{}

// GetEmployeePending ...
func (u *Supervisor) GetEmployeePending(supervisorID int64) (reqPending []structLogic.LeavePending, err error) {
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
		Where(`status = ? `).And(`supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statPendingInSupervisor := constant.StatusPendingInSupervisor

	count, errRaw := o.Raw(sql, statPendingInSupervisor, supervisorID).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeePending", errRaw)
		return reqPending, errors.New("Error get leave request pending")
	}
	beego.Debug("Total request pending  =", count)

	return reqPending, errRaw
}

// GetEmployeeApproved ...
func (u *Supervisor) GetEmployeeApproved(supervisorID int64) (reqApprove []structLogic.LeaveAccept, err error) {
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
		Where(`(status = ? OR status = ? OR status = ? )`).And(user.TableName() + `.supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statPendingInDirector := constant.StatusPendingInDirector
	statSuccessInDirector := constant.StatusSuccessInDirector
	statsRejectInDirector := constant.StatusRejectInDirector

	count, errRaw := o.Raw(sql, statPendingInDirector, statSuccessInDirector, statsRejectInDirector, supervisorID).QueryRows(&reqApprove)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeeApproved", errRaw)
		return reqApprove, errors.New("Error get leave request approved")
	}
	beego.Debug("Total request approved  =", count)

	return reqApprove, errRaw
}

// GetEmployeeRejected ...
func (u *Supervisor) GetEmployeeRejected(supervisorID int64) (reqReject []structLogic.LeaveReject, err error) {
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
		Where(`status = ?`).And(user.TableName() + `.supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statRejectInSupervisor := constant.StatusRejectInSuperVisor

	count, errRaw := o.Raw(sql, statRejectInSupervisor, supervisorID).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployeeRejected", errRaw)
		return reqReject, errors.New("Error get leave request reject")
	}
	beego.Debug("Total request rejected  =", count)

	return reqReject, errRaw
}

// ApproveBySupervisor ...
func (u *Supervisor) ApproveBySupervisor(id int64, employeeNumber int64, actionBy string) error {
	var dbLeave structDB.LeaveRequest

	statPendingDirector := constant.StatusPendingInDirector

	o := orm.NewOrm()

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statPendingDirector, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("Error update status approve @ApproveBySupervisor", errRAW)
		return errors.New("Update status approve failed")
	}

	return errRAW
}

// RejectBySupervisor ...
func (u *Supervisor) RejectBySupervisor(l *structLogic.LeaveReason, id int64, employeeNumber int64, actionBy string) error {
	var dbLeave structDB.LeaveRequest

	statRejectSupervisor := constant.StatusRejectInSuperVisor
	rejectReason := l.RejectReason

	o := orm.NewOrm()

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, reject_reason = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statRejectSupervisor, rejectReason, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("Error update status reject @RejectBySupervisor", errRAW)
		return errors.New("Update status reject failed")
	}

	return errRAW
}
