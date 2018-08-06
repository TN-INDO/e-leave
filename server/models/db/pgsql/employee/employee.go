package employee

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Employee ...
type Employee struct{}

// GetPendingRequest ...
func (e *Employee) GetPendingRequest(employeeNumber int64) ([]structLogic.RequestPending, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqPending    []structLogic.RequestPending
	)

	statPendingSupervisor := constant.StatusPendingInSupervisor
	statPendingDirector := constant.StatusPendingInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorPendingRequest", errQB)
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
		// typeLeave.TableName()+".id as type_id",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
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
		Where(`(status = ? OR status = ? )`).And(user.TableName() + `.employee_number = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statPendingSupervisor, statPendingDirector, employeeNumber).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorPendingRequest", errRaw)
		return reqPending, errors.New("error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetApprovedRequest ...
func (e *Employee) GetApprovedRequest(employeeNumber int64) ([]structLogic.RequestAccept, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqAccept     []structLogic.RequestAccept
	)

	statAcceptDirector := constant.StatusSuccessInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorPendingRequest", errQB)
		return reqAccept, errQB
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
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
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
		Where(`status = ?`).And(user.TableName() + `.employee_number = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statAcceptDirector, employeeNumber).QueryRows(&reqAccept)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorPendingRequest", errRaw)
		return reqAccept, errors.New("error get leave request pending")
	}
	beego.Debug("Total approved request =", count)

	return reqAccept, errRaw
}

// GetRejectedRequest ...
func (e *Employee) GetRejectedRequest(employeeNumber int64) ([]structLogic.RequestReject, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqReject     []structLogic.RequestReject
	)

	statRejectSupervisor := constant.StatusRejectInSuperVisor
	statRejectDirector := constant.StatusRejectInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorPendingRequest", errQB)
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
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
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
		Where(`(status = ? OR status = ? )`).And(user.TableName() + `.employee_number = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statRejectSupervisor, statRejectDirector, employeeNumber).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorPendingRequest", errRaw)
		return reqReject, errors.New("error get leave request pending")
	}
	beego.Debug("Total rejected request =", count)

	return reqReject, errRaw
}
