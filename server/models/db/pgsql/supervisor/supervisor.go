package user

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	logicLeave "server/models/db/pgsql/leave_request"
	logicUser "server/models/db/pgsql/user"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Supervisor ...
type Supervisor struct{}

// GetUserPending ...
func (u *Supervisor) GetUserPending(supervisorID int64) ([]structLogic.LeavePending, error) {
	var (
		reqPending    []structLogic.LeavePending
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	statPendingInSupervisor := constant.StatusPendingInSupervisor
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
		Where(`status = ? `).And(`supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statPendingInSupervisor, supervisorID).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorPendingRequest", errRaw)
		return reqPending, errors.New("error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetUserAccept ...
func (u *Supervisor) GetUserAccept(supervisorID int64) ([]structLogic.LeaveAccept, error) {
	var (
		reqAccepts    []structLogic.LeaveAccept
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	statPendingInDirector := constant.StatusPendingInDirector
	statSuccessInDirector := constant.StatusSuccessInDirector
	statsRejectInDirector := constant.StatusRejectInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserAccept", errQB)
		return reqAccepts, errQB
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
		Where(`(status = ? OR status = ? OR status = ? )`).And(user.TableName() + `.supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statPendingInDirector, statSuccessInDirector, statsRejectInDirector, supervisorID).QueryRows(&reqAccepts)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetUserAccept", errRaw)
		return reqAccepts, errors.New("error get leave request accept")
	}
	beego.Debug("Total accept request =", count)

	return reqAccepts, errRaw
}

// GetUserReject ...
func (u *Supervisor) GetUserReject(supervisorID int64) ([]structLogic.LeaveReject, error) {
	var (
		reqRejects    []structLogic.LeaveReject
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	statRejectInSupervisor := constant.StatusRejectInSuperVisor

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserReject", errQB)
		return reqRejects, errQB
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
		Where(`status = ?`).And(user.TableName() + `.supervisor_id = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	count, errRaw := o.Raw(sql, statRejectInSupervisor, supervisorID).QueryRows(&reqRejects)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetUserReject", errRaw)
		return reqRejects, errors.New("error get leave request reject")
	}
	beego.Debug("Total reject request =", count)

	return reqRejects, errRaw
}

// AcceptBySupervisor ...
func (u *Supervisor) AcceptBySupervisor(id int64, employeeNumber int64) error {
	var (
		dbLeave structDB.LeaveRequest
		user    logicUser.User
		leave   logicLeave.LeaveRequest
	)

	o := orm.NewOrm()

	getEmployee, _ := user.GetEmployee(employeeNumber)
	getSupervisorID, _ := user.GetSupervisor(employeeNumber)
	getSupervisor, _ := user.GetEmployee(getSupervisorID.SupervisorID)
	getDirector, _ := user.GetDirector()
	getLeave, _ := leave.GetLeave(id)

	statPendingDirector := constant.StatusPendingInDirector
	actionBy := getSupervisor.Name

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statPendingDirector, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("error update status @AcceptBySupervisor", errRAW)
	}

	helpers.GoMailEmployee(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name)
	helpers.GoMailDirector(getDirector.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, getDirector.Name)

	return errRAW
}

// RejectBySv ...
func (u *Supervisor) RejectBySv(l *structLogic.LeaveReason, id int64, employeeNumber int64) error {
	var (
		dbLeave structDB.LeaveRequest
		user    logicUser.User
		leave   logicLeave.LeaveRequest
	)

	o := orm.NewOrm()

	getEmployee, _ := user.GetEmployee(employeeNumber)
	getSupervisorID, _ := user.GetSupervisor(employeeNumber)
	getSupervisor, _ := user.GetEmployee(getSupervisorID.SupervisorID)
	getLeave, _ := leave.GetLeave(id)

	statRejectSupervisor := constant.StatusRejectInSuperVisor
	rejectReason := l.RejectReason
	actionBy := getSupervisor.Name

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, reject_reason = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statRejectSupervisor, rejectReason, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("error update status @RejectBySv", errRAW)
	}

	helpers.GoMailSupervisorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getSupervisor.Name, rejectReason)

	return errRAW
}
