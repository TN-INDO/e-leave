package leave

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	logicUser "server/models/db/pgsql/user"
	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// LeaveRequest ...
type LeaveRequest struct{}

// CreateLeaveRequest ...
func (l *LeaveRequest) CreateLeaveRequest(employeeNumber int64,
	typeLeaveID int64,
	reason string,
	dateFrom string,
	dateTo string,
	halfDates []string,
	backOn string,
	total float64,
	address string,
	contactLeave string,
	status string) error {

	var leave structDB.LeaveRequest
	var user logicUser.User
	isHalfDay := helpers.ArrayToString(halfDates, ",")

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserPending", errQB)
		return errQB
	}

	qb.InsertInto(
		leave.TableName(),
		"employee_number",
		"type_leave_id",
		"reason",
		"date_from",
		"date_to",
		"half_dates",
		"back_on",
		"total",
		"contact_address",
		"contact_number",
		"status").
		Values("?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?")
	sql := qb.String()

	values := []interface{}{employeeNumber,
		typeLeaveID,
		reason,
		dateFrom,
		dateTo,
		"{" + isHalfDay + "}",
		backOn,
		total,
		address,
		contactLeave,
		status}
	_, err := o.Raw(sql, values).Exec()
	if err != nil {
		helpers.CheckErr("error insert @CreateLeaveRequestEmployee", err)
		return errors.New("insert create leave request failed")
	}

	getEmployee, _ := user.GetEmployee(employeeNumber)
	getSupervisorID, _ := user.GetSupervisor(employeeNumber)
	getSupervisor, _ := user.GetEmployee(getSupervisorID.SupervisorID)

	defer helpers.GoMailSupervisor(getSupervisor.Email, getEmployee.Name, getSupervisor.Name)

	return err
}

// CreateLeaveRequestSupervisor ...
func (l *LeaveRequest) CreateLeaveRequestSupervisor(employeeNumber int64,
	typeLeaveID int64,
	reason string,
	dateFrom string,
	dateTo string,
	halfDates []string,
	backOn string,
	total float64,
	address string,
	contactLeave string,
	status string) error {

	var leave structDB.LeaveRequest
	var user logicUser.User

	isHalfDay := helpers.ArrayToString(halfDates, ",")

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserPending", errQB)
		return errQB
	}

	qb.InsertInto(
		leave.TableName(),
		"employee_number",
		"type_leave_id",
		"reason",
		"date_from",
		"date_to",
		"half_dates",
		"back_on",
		"total",
		"contact_address",
		"contact_number",
		"status").
		Values("?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?")
	sql := qb.String()

	values := []interface{}{employeeNumber,
		typeLeaveID,
		reason,
		dateFrom,
		dateTo,
		"{" + isHalfDay + "}",
		backOn,
		total,
		address,
		contactLeave,
		status}
	_, err := o.Raw(sql, values).Exec()
	if err != nil {
		helpers.CheckErr("error insert @CreateLeaveRequestSupervisor", err)
		return errors.New("insert create leave request failed")
	}

	getEmployee, _ := user.GetEmployee(employeeNumber)
	getDirector, _ := user.GetDirector()

	helpers.GoMailDirectorFromSupervisor(getDirector.Email, getEmployee.Name, getDirector.Name)

	return err
}

// UpdateRequest ...
func (l *LeaveRequest) UpdateRequest(e *structAPI.UpdateLeaveRequest, id int64) (err error) {
	var dbLeave structDB.LeaveRequest
	isHalfDay := helpers.ArrayToString(e.HalfDates, ",")
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateRequest", errQB)
		return errQB
	}

	qb.Update(dbLeave.TableName()).
		Set("type_leave_id = ?",
			"reason = ?",
			"date_from = ?",
			"date_to = ?",
			"half_dates = ?",
			"back_on = ?",
			"total = ?",
			"contact_address = ?",
			"contact_number = ?").Where("id = ? ")
	sql := qb.String()

	res, errRaw := o.Raw(sql,
		e.TypeLeaveID,
		e.Reason,
		e.DateFrom,
		e.DateTo,
		"{"+isHalfDay+"}",
		e.BackOn,
		e.Total,
		e.ContactAddress,
		e.ContactNumber,
		id).Exec()

	if errRaw != nil {
		helpers.CheckErr("err update @UpdateRequest", errRaw)
		return errors.New("update request failed")
	}

	_, errRow := res.RowsAffected()
	if errRow != nil {
		helpers.CheckErr("error get rows affected", errRow)
		return errRow
	}

	return err
}

// DeleteRequest ...
func (l *LeaveRequest) DeleteRequest(id int64) (err error) {
	o := orm.NewOrm()
	v := structDB.LeaveRequest{ID: id}

	err = o.Read(&v)
	if err == nil {
		var num int64
		if num, err = o.Delete(&structDB.LeaveRequest{ID: id}); err == nil {
			beego.Debug("Number of records deleted in database:", num)
		} else if err != nil {
			helpers.CheckErr("error deleted @DeleteRequest", err)
			return errors.New("error deleted leave request")
		}
	}
	if err != nil {
		helpers.CheckErr("error deleted @DeleteRequest", err)
		return errors.New("Delete failed, id not exist")
	}
	return err
}

// GetLeave ...
func (l *LeaveRequest) GetLeave(id int64) (result structLogic.GetLeave, err error) {
	var dbLeave structDB.LeaveRequest

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeave", errQB)
		return result, errQB
	}

	qb.Select(dbLeave.TableName()+".id",
		dbLeave.TableName()+".type_leave_id",
		dbLeave.TableName()+".total").
		From(dbLeave.TableName()).
		Where(dbLeave.TableName() + `.id = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, id).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetLeave", errRaw)
		return result, errors.New("id not exist")
	}
	return result, err
}

// UpdateLeaveRemaningApprove ...
func (l *LeaveRequest) UpdateLeaveRemaningApprove(total float64, employeeNumber int64, typeID int64) (err error) {
	var e *structDB.UserTypeLeave
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateLeaveRemaningApprove", errQB)
		return errQB
	}

	qb.Update(e.TableName()).Set("leave_remaining = leave_remaining - ?").
		Where(`(employee_number = ? AND type_leave_id = ? )`)
	sql := qb.String()

	res, errRaw := o.Raw(sql, total, employeeNumber, typeID).Exec()

	if errRaw != nil {
		helpers.CheckErr("err update @UpdateLeaveRemaningApprove", errRaw)
		return errors.New("update leave remaining failed")
	}

	_, errRow := res.RowsAffected()
	if errRow != nil {
		helpers.CheckErr("error get rows affected", errRow)
		return errRow
	}

	return err
}

// UpdateLeaveRemaningCancel ...
func (l *LeaveRequest) UpdateLeaveRemaningCancel(total float64, employeeNumber int64, typeID int64) (err error) {
	var e *structDB.UserTypeLeave
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateLeaveRemaningCancel", errQB)
		return errQB
	}

	qb.Update(e.TableName()).Set("leave_remaining = leave_remaining + ?").
		Where(`(employee_number = ? AND type_leave_id = ? )`)
	sql := qb.String()

	res, errRaw := o.Raw(sql, total, employeeNumber, typeID).Exec()

	if errRaw != nil {
		helpers.CheckErr("err update @UpdateLeaveRemaningCancel", errRaw)
		return errors.New("update leave remaining failed")
	}

	_, errRow := res.RowsAffected()
	if errRow != nil {
		helpers.CheckErr("error get rows affected", errRow)
		return errRow
	}

	return err
}

// DownloadReportCSV ...
func (l *LeaveRequest) DownloadReportCSV(query *structAPI.RequestReport,
	path string) (err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		report        []structLogic.ReportLeaveRequest
	)
	statAcceptDirector := constant.StatusSuccessInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @DownloadReportCSV", errQB)
		return errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".email",
		typeLeave.TableName()+".type_name",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number").
		From(leave.TableName()).
		InnerJoin(user.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(leave.TableName() + `.status = ? `).
		And(leave.TableName() + `.created_at >= ?`).And(leave.TableName() + `.created_at <= ?`)

	sql := qb.String()

	count, errRaw := o.Raw(sql, statAcceptDirector, query.FromDate, query.ToDate).QueryRows(&report)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @DownloadReportCSV", errRaw)
		return errRaw
	}
	beego.Debug("Total leave request =", count)

	l.WriteCsv(path, report)

	return err
}

// WriteCsv ...
func (l *LeaveRequest) WriteCsv(path string, res []structLogic.ReportLeaveRequest) error {

	w, err := helpers.NewCsvWriter(path)
	if err != nil {
		beego.Debug(err)
		return err
	}

	w.Write([]string{
		"No.",
		"Request ID",
		"Employee Number",
		"Name",
		"Gender",
		"Position",
		"Start Working Date",
		"Email",
		"Type of Leave",
		"Reason",
		"From",
		"To",
		"Half Day",
		"Back To Work",
		"Total Leave",
		"Leave Balance",
		"Contact Address",
		"Contact Number",
	})

	count := len(res)

	for i := 0; i < count; i++ {
		w.Write([]string{
			strconv.Itoa(int(i+1)) + ".",
			strconv.Itoa(int(res[i].ID)),
			strconv.Itoa(int(res[i].EmployeeNumber)),
			res[i].Name,
			res[i].Gender,
			res[i].Position,
			res[i].StartWorkingDate,
			res[i].Email,
			res[i].TypeName,
			res[i].Reason,
			res[i].DateFrom,
			res[i].DateTo,
			res[i].HalfDates,
			res[i].BackOn,
			strconv.FormatFloat(res[i].Total, 'f', 1, 64) + " days",
			strconv.FormatFloat(res[i].LeaveRemaining, 'f', 1, 64) + " days",
			res[i].ContactAddress,
			res[i].ContactNumber,
		})

	}
	w.Flush()

	return err
}

// ReportLeaveRequest ...
func (l *LeaveRequest) ReportLeaveRequest(query *structAPI.RequestReport) (res []structLogic.ReportLeaveRequest, err error) {
	var (
		report        []structLogic.ReportLeaveRequest
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)
	statAcceptDirector := constant.StatusSuccessInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @ReportLeaveRequest", errQB)
		return report, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".email",
		typeLeave.TableName()+".type_name",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number").
		From(leave.TableName()).
		InnerJoin(user.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where("strftime(" + leave.TableName() + ".date_from) >= strftime(?) ").
		And("strftime(" + leave.TableName() + ".date_from) <= strftime(?) ").
		And(leave.TableName() + `.status = ? `).
		OrderBy("strftime(" + leave.TableName() + ".date_from) ASC ")
	sql := qb.String()

	count, errRaw := o.Raw(sql, query.FromDate, query.ToDate, statAcceptDirector).QueryRows(&report)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @ReportLeaveRequest", errRaw)
		return nil, errRaw
	}
	beego.Debug("Total leave request =", count)

	return report, err
}

// ReportLeaveRequestTypeLeave ...
func (l *LeaveRequest) ReportLeaveRequestTypeLeave(query *structAPI.RequestReportTypeLeave) (res []structLogic.ReportLeaveRequest, err error) {
	var (
		report        []structLogic.ReportLeaveRequest
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)
	statAcceptDirector := constant.StatusSuccessInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @ReportLeaveRequest", errQB)
		return report, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".email",
		typeLeave.TableName()+".type_name",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		// "array_to_string("+leave.TableName()+".half_dates, ', ') as half_dates",
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number").
		From(leave.TableName()).
		InnerJoin(user.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where("strftime(" + leave.TableName() + ".date_from) >= strftime(?) ").
		And("strftime(" + leave.TableName() + ".date_from) <= strftime(?) ").
		And(leave.TableName() + `.status = ?`).
		And(leave.TableName() + `.type_leave_id = ?`).
		OrderBy("strftime(" + leave.TableName() + ".date_from) ASC ")
	sql := qb.String()

	// id, errCon := strconv.ParseInt(query.TypeLeaveID, 0, 64)
	// if errCon != nil {
	// 	helpers.CheckErr("convert id failed @ReportLeaveRequestTypeLeave", errCon)
	// }

	count, errRaw := o.Raw(sql, query.FromDate, query.ToDate, statAcceptDirector, query.TypeLeaveID).QueryRows(&report)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @ReportLeaveRequest", errRaw)
		return nil, errRaw
	}
	beego.Debug("Total leave request =", count)

	return report, err
}
