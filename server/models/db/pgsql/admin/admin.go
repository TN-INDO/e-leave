package admin

import (
	"encoding/base64"
	"errors"
	"server/helpers"
	"server/helpers/constant"

	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

// Admin ...
type Admin struct{}

// AddUser ...
func (u *Admin) AddUser(user structDB.User) error {
	var (
		count               int
		countEmployeeNumber int
	)

	o := orm.NewOrm()

	o.Raw(`SELECT count(*) as Count FROM `+user.TableName()+` WHERE email = ?`, user.Email).QueryRow(&count)
	o.Raw(`SELECT count(*) as Count FROM `+user.TableName()+` WHERE employee_number = ?`, user.EmployeeNumber).QueryRow(&countEmployeeNumber)

	passwordString := user.Password
	bsEmployeeNumber := []byte(strconv.Itoa(int(user.EmployeeNumber)))
	arrPassword := []byte(passwordString)

	if len(bsEmployeeNumber) != 5 {
		return errors.New("Employee number must field and length must be 5")
	} else if countEmployeeNumber > 0 {
		return errors.New("Employee number already register")
	} else if user.Name == "" || user.Gender == "" || user.Position == "" || user.StartWorkingDate == "" || user.MobilePhone == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("Error empty field ")
	} else if count > 0 {
		return errors.New("Email already register")
	} else if len(arrPassword) < 7 {
		return errors.New("Password length must be 7")
	} else {
		hashedBytes, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		helpers.CheckErr("Error hash password @AddUser", errHash)

		user.Email = strings.ToLower(user.Email)
		user.Password = base64.StdEncoding.EncodeToString(hashedBytes)

		_, err := o.Insert(&user)
		if err != nil {
			helpers.CheckErr("Error insert @AddUser", err)
			return errors.New("Insert users failed")
		}

		go func() {
			helpers.GoMailRegisterPassword(user.Email, passwordString)

		}()

		return err
	}
}

// DeleteUser ...
func (u *Admin) DeleteUser(employeeNumber int64) (err error) {
	o := orm.NewOrm()
	v := structDB.User{EmployeeNumber: employeeNumber}

	err = o.Read(&v)
	if err == nil {
		var num int64
		if num, err = o.Delete(&structDB.User{EmployeeNumber: employeeNumber}); err == nil {
			beego.Debug("Number of records deleted in database:", num)
		} else if err != nil {
			helpers.CheckErr("Error delete user @DeleteUser", err)
			return errors.New("Error delete user")
		}
	}
	if err != nil {
		helpers.CheckErr("Error delete user @DeleteUser", err)
		return errors.New("Delete failed, id not exist")
	}

	return err
}

// GetUsers ...
func (u *Admin) GetUsers() (result []structDB.User, err error) {
	var (
		dbUser structDB.User
		roles  []string
	)
	roles = append(roles, "employee", "supervisor", "director")

	o := orm.NewOrm()
	count, err := o.Raw("SELECT * FROM "+dbUser.TableName()+" WHERE role IN (?,?,?)", roles).QueryRows(&result)
	if err != nil {
		helpers.CheckErr("Failed get users @GetUsers", err)
		return result, err
	}
	beego.Debug("Total user =", count)

	return result, err
}

// GetUser ...
func (u *Admin) GetUser(employeeNumber int64) (result structDB.User, err error) {
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUser", errQB)
		return result, errQB
	}

	qb.Select("*").From(result.TableName()).
		Where(`employee_number = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, employeeNumber).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select item @GetUser", errRaw)
		return result, errors.New("Employee number not exist")
	}

	return result, err
}

// UpdateUser ...
func (u *Admin) UpdateUser(e *structDB.User, employeeNumber int64) (err error) {
	var (
		user  structLogic.GetEmployee
		count int
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateUser", errQB)
		return errQB
	}

	o.Raw(`SELECT name, email FROM `+e.TableName()+` WHERE employee_number = ?`, employeeNumber).QueryRow(&user)

	if e.Email != user.Email {
		o.Raw(`SELECT count(*) as Count FROM `+e.TableName()+` WHERE email = ?`, e.Email).QueryRow(&count)
		if count > 0 {
			return errors.New("Email already register")
		} else {
			qb.Update(e.TableName()).
				Set("name = ?",
					"gender = ?",
					"position = ?",
					"start_working_date = ?",
					"mobile_phone = ?",
					"email= ?",
					"role = ?",
					"supervisor_id = ?",
					"updated_at = ?").Where("employee_number = ? ")
			sql := qb.String()

			e.Email = strings.ToLower(e.Email)

			res, errRaw := o.Raw(sql,
				e.Name,
				e.Gender,
				e.Position,
				e.StartWorkingDate,
				e.MobilePhone,
				e.Email,
				e.Role,
				e.SupervisorID,
				e.UpdatedAt,
				employeeNumber).Exec()

			if errRaw != nil {
				helpers.CheckErr("Error update user @UpdateUser", errRaw)
				return errors.New("Update user failed")
			}

			_, errRow := res.RowsAffected()
			if errRow != nil {
				helpers.CheckErr("Error get rows affected @UpdateUser", errRow)
				return errRow
			}
		}
	} else {
		qb.Update(e.TableName()).
			Set("name = ?",
				"gender = ?",
				"position = ?",
				"start_working_date = ?",
				"mobile_phone = ?",
				"email= ?",
				"role = ?",
				"supervisor_id = ?",
				"updated_at = ?").Where("employee_number = ? ")
		sql := qb.String()

		e.Email = strings.ToLower(e.Email)

		res, errRaw := o.Raw(sql,
			e.Name,
			e.Gender,
			e.Position,
			e.StartWorkingDate,
			e.MobilePhone,
			e.Email,
			e.Role,
			e.SupervisorID,
			e.UpdatedAt,
			employeeNumber).Exec()

		if errRaw != nil {
			helpers.CheckErr("Error update user @UpdateUser", errRaw)
			return errors.New("Update user failed")
		}

		_, errRow := res.RowsAffected()
		if errRow != nil {
			helpers.CheckErr("Error get rows affected @UpdateUser", errRow)
			return errRow
		}
	}

	return err
}

// GetLeaveRequestPending ...
func (u *Admin) GetLeaveRequestPending() ([]structLogic.RequestPending, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqPending    []structLogic.RequestPending
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequestPending", errQB)
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
		Where(`(status = ? OR status = ? )`).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statPendingInSupervisor := constant.StatusPendingInSupervisor
	statPendingInDirector := constant.StatusPendingInDirector

	count, errRaw := o.Raw(sql, statPendingInSupervisor, statPendingInDirector).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetLeaveRequestPending", errRaw)
		return reqPending, errors.New("Error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetLeaveRequest ...
func (u *Admin) GetLeaveRequest() ([]structLogic.RequestAccept, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqApprove    []structLogic.RequestAccept
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequest", errQB)
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
		helpers.CheckErr("Failed query select @GetLeaveRequest", errRaw)
		return reqApprove, errors.New("Error get leave request approved")
	}
	beego.Debug("Total approved request =", count)

	return reqApprove, errRaw
}

// GetLeaveRequestReject ...
func (u *Admin) GetLeaveRequestReject() ([]structLogic.RequestReject, error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
		reqReject     []structLogic.RequestReject
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequestReject", errQB)
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
		Where(`(status = ? OR status = ? )`).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statRejectInSuperVisor := constant.StatusRejectInSuperVisor
	statRejectInDirector := constant.StatusRejectInDirector

	count, errRaw := o.Raw(sql, statRejectInSuperVisor, statRejectInDirector).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetLeaveRequestReject", errRaw)
		return reqReject, errors.New("Error get leave request reject")
	}
	beego.Debug("Total reject request =", count)

	return reqReject, errRaw
}

// CreateUserTypeLeave ...
func (u *Admin) CreateUserTypeLeave(
	employeeNumber int64,
	typeLeaveID int64,
	leaveRemaining float64,
) error {
	var typeLeave structDB.UserTypeLeave

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @CreateUserTypeLeave", errQB)
		return errQB
	}

	qb.InsertInto(
		typeLeave.TableName(),
		"employee_number",
		"type_leave_id",
		"leave_remaining").
		Values("?, ?, ?")
	sql := qb.String()

	values := []interface{}{
		employeeNumber,
		typeLeaveID,
		leaveRemaining,
	}
	_, err := o.Raw(sql, values).Exec()
	if err != nil {
		helpers.CheckErr("Error insert @CreateUserTypeLeave", err)
		return errors.New("Insert user type leave failed")
	}

	return err
}

// UpdateLeaveRemaning ...
func (u *Admin) UpdateLeaveRemaning(total float64, employeeNumber int64, typeID int64) (err error) {
	var e *structDB.UserTypeLeave

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateLeaveRemaning", errQB)
		return errQB
	}

	qb.Update(e.TableName()).Set("leave_remaining = leave_remaining - ?").
		Where(`(employee_number = ? AND type_leave_id = ? )`)
	sql := qb.String()

	res, errRaw := o.Raw(sql, total, employeeNumber, typeID).Exec()
	if errRaw != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaning", errRaw)
		return errors.New("Update leave balance failed")
	}

	_, errRow := res.RowsAffected()
	if errRow != nil {
		helpers.CheckErr("Error get rows affected @UpdateLeaveRemaning", errRow)
		return errRow
	}

	return err
}

// ResetUserTypeLeave ...
func (u *Admin) ResetUserTypeLeave(leaveRemaining float64, typeLeaveID int64) error {
	var typeLeave structDB.UserTypeLeave

	o := orm.NewOrm()

	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @ResetUserTypeLeave", errQB)
		return errQB
	}

	qb.Update(typeLeave.TableName()).
		Set("leave_remaining = ?").
		Where("type_leave_id = ?")
	sql := qb.String()

	res, errRaw := o.Raw(sql, leaveRemaining, typeLeaveID).Exec()
	if errRaw != nil {
		helpers.CheckErr("Error update leave balance @ResetUserTypeLeave", errRaw)
		return errors.New("Reset leave balance failed")
	}

	_, errRow := res.RowsAffected()
	helpers.CheckErr("Error get rows affected @ResetUserTypeLeave", errRow)

	return errRow
}

// UpdateUserTypeLeave ...
func (u *Admin) UpdateUserTypeLeave(
	leaveRemaining float64,
	typeLeaveID int64,
	employeeNumber int64,
) error {
	var typeLeave structDB.UserTypeLeave

	o := orm.NewOrm()

	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateUserTypeLeave", errQB)
		return errQB
	}

	qb.Update(typeLeave.TableName()).
		Set(
			"leave_remaining = ?",
			"type_leave_id = ?",
		).
		Where("employee_number = ?")
	sql := qb.String()

	res, errRaw := o.Raw(sql, leaveRemaining, typeLeaveID, employeeNumber).Exec()
	if errRaw != nil {
		helpers.CheckErr("Error update @UpdateUserTypeLeave", errRaw)
		return errors.New("Update request failed")
	}

	_, errRow := res.RowsAffected()
	helpers.CheckErr("Error get rows affected @UpdateUserTypeLeave", errRow)

	return errRow
}
