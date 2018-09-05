package user

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"

	userLogic "server/models/db/pgsql/admin"
	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego/orm"
)

// User ...
type User struct{}

// UserLogin ...
func (u *User) UserLogin(loginData *structAPI.ReqLogin) (user structDB.User, err error) {

	o := orm.NewOrm()
	errRaw := o.Raw(`SELECT * FROM `+user.TableName()+` WHERE email = ?`, loginData.Email).QueryRow(&user)
	if errRaw != nil {
		helpers.CheckErr("Error get user @UserLogin", errRaw)
		return user, errors.New("Email not register")
	}

	return user, err
}

// ForgotPassword ...
func (u *User) ForgotPassword(e *structLogic.PasswordReset) error {
	var dbUser structDB.User
	passwordReset := "JDJhJDEwJGtLeW42RFBOOEt2WUdpMlZHdHJ6bnVqY0gyU0lYUFNBMFVDQ0VQMW1kSWRIcHRmdWRsTmJl"

	o := orm.NewOrm()

	_, errRAW := o.Raw(`UPDATE `+dbUser.TableName()+` SET password = ? WHERE email = ?`, passwordReset, e.Email).Exec()
	if errRAW != nil {
		helpers.CheckErr("Error forgot password @ForgotPassword", errRAW)
		return errors.New("Error forgot password @ForgotPassword")
	}

	return errRAW
}

// CountUserEmail ...
func (u *User) CountUserEmail(email string) (int, error) {
	var (
		dbUser structDB.User
		count  int
	)

	o := orm.NewOrm()
	errGet := o.Raw(`SELECT count(*) as Count FROM `+dbUser.TableName()+` WHERE email = ?`, email).QueryRow(&count)
	if errGet != nil {
		helpers.CheckErr("Failed query select @GetTypeLeave", errGet)
		return count, errors.New("Error count user by email")
	}

	return count, errGet
}

// GetUser ...
func (u *User) GetUser(email string) (structLogic.GetEmployee, error) {
	var (
		dbUser   structDB.User
		employee structLogic.GetEmployee
	)

	o := orm.NewOrm()
	errGet := o.Raw(`SELECT name, email FROM `+dbUser.TableName()+` WHERE email = ?`, email).QueryRow(&employee)
	if errGet != nil {
		helpers.CheckErr("Failed query select @GetUser", errGet)
		return employee, errors.New("Error get user")
	}

	return employee, errGet
}

// UpdatePassword ...
func (u *User) UpdatePassword(p *structLogic.NewPassword, employeeNumber int64) (err error) {
	var (
		dbUser structDB.User
		admin  userLogic.Admin
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdatePassword", errQB)
		return errQB
	}

	bsNewPassword := []byte(p.NewPassword)
	bsConfirmPassword := []byte(p.ConfirmPassword)

	resGet, errGet := admin.GetUser(employeeNumber)
	helpers.CheckErr("Error get user @UpdatePassword", errGet)

	resComparePassword := helpers.ComparePassword(resGet.Password, p.OldPassword)

	if resComparePassword == true {
		if len(bsNewPassword) < 7 && len(bsConfirmPassword) < 7 {
			return errors.New("Password length minimum must be 7")
		} else if p.NewPassword == p.ConfirmPassword {
			qb.Update(dbUser.TableName()).Set("password = ?").
				Where(`employee_number = ?`)
			sql := qb.String()

			resPassword, errHash := helpers.HashPassword(p.NewPassword)
			if errHash != nil {
				helpers.CheckErr("err hash password @UpdatePassword", errHash)
			}

			res, errRaw := o.Raw(sql, resPassword, employeeNumber).Exec()

			if errRaw != nil {
				helpers.CheckErr("err update password @UpdatePassword", errRaw)
				return errors.New("update password failed")
			}

			_, errRow := res.RowsAffected()
			if errRow != nil {
				helpers.CheckErr("error get rows affected", errRow)
				return errRow
			}
		} else {
			return errors.New("wrong confirm password")
		}

	} else {
		return errors.New("wrong old password")
	}

	return err
}

// GetDirector ...
func (u *User) GetDirector() (result structLogic.GetDirector, err error) {
	var dbUser structDB.User
	role := "director"

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirector", errQB)
		return result, errQB
	}

	qb.Select(
		dbUser.TableName()+".name",
		dbUser.TableName()+".email").
		From(dbUser.TableName()).
		Where(dbUser.TableName() + `.role = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, role).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetDirector", errRaw)
		return result, errors.New("No role with director")
	}

	return result, err
}

// GetSupervisor ...
func (u *User) GetSupervisor(employeeNumber int64) (result structLogic.GetSupervisor, err error) {
	var (
		dbUser  structDB.User
		dbLeave structDB.LeaveRequest
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetSupervisor", errQB)
		return result, errQB
	}

	qb.Select(
		dbUser.TableName()+".supervisor_id",
		dbUser.TableName()+".name",
		dbUser.TableName()+".email").
		From(dbUser.TableName()).
		InnerJoin(dbLeave.TableName()).
		On(dbUser.TableName() + ".employee_number" + "=" + dbLeave.TableName() + ".employee_number").
		Where(dbUser.TableName() + `.employee_number = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, employeeNumber).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetSupervisor", errRaw)
		return result, errors.New("Employee number not exist")
	}

	return result, err
}

// GetEmployee ...
func (u *User) GetEmployee(employeeNumber int64) (result structLogic.GetEmployee, err error) {
	var dbUser structDB.User

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetEmployee", errQB)
		return result, errQB
	}

	qb.Select(
		dbUser.TableName()+".name",
		dbUser.TableName()+".email").
		From(dbUser.TableName()).
		Where(dbUser.TableName() + `.employee_number = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, employeeNumber).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetEmployee", errRaw)
		return result, errors.New("Employee number not exist")
	}

	return result, err
}

// GetSumarry ...
func (u *User) GetSumarry(employeeNumber int64) ([]structLogic.UserSumarry, error) {
	var (
		dbLeave     structDB.LeaveRequest
		dbTypeLeave structDB.TypeLeave
		sumarry     []structLogic.UserSumarry
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetSumarry", errQB)
		return sumarry, errQB
	}

	qb.Select(
		dbTypeLeave.TableName()+".type_name",
		"SUM("+dbLeave.TableName()+".total) as used").
		From(dbLeave.TableName()).
		InnerJoin(dbTypeLeave.TableName()).
		On(dbTypeLeave.TableName() + ".id" + " = " + dbLeave.TableName() + ".type_leave_id").
		Where(dbLeave.TableName() + `.employee_number = ? `).
		And(dbLeave.TableName() + `.status = ?`).
		GroupBy(dbTypeLeave.TableName() + `.type_name`)
	sql := qb.String()

	statSuccessInDirector := constant.StatusSuccessInDirector

	_, errRaw := o.Raw(sql, employeeNumber, statSuccessInDirector).QueryRows(&sumarry)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetSumarry", errRaw)
		return sumarry, errors.New("Error get user summary")
	}

	return sumarry, errRaw
}

// GetUserTypeLeave ...
func (u *User) GetUserTypeLeave(employeeNumber int64) (result []structLogic.UserTypeLeave, err error) {
	var (
		dbType          structDB.TypeLeave
		dbUserTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserTypeLeave", errQB)
		return result, errQB
	}

	qb.Select(
		dbType.TableName()+".type_name",
		dbUserTypeLeave.TableName()+".leave_remaining").
		From(dbType.TableName()).
		InnerJoin(dbUserTypeLeave.TableName()).
		On(dbUserTypeLeave.TableName() + ".type_leave_id" + " = " + dbType.TableName() + ".id").
		Where(dbUserTypeLeave.TableName() + `.employee_number = ? `)
	sql := qb.String()

	_, errRaw := o.Raw(sql, employeeNumber).QueryRows(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetUserTypeLeave", errRaw)
		return result, errors.New("Error get user type leave")
	}

	return result, err
}

// GetUserLeaveRemaining ...
func (u *User) GetUserLeaveRemaining(typeID int64, employeeNumber int64) (result structLogic.UserTypeLeave, err error) {
	var (
		dbType          structDB.TypeLeave
		dbUserTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUserLeaveRemaining", errQB)
		return result, errQB
	}

	qb.Select(
		dbType.TableName()+".type_name",
		dbUserTypeLeave.TableName()+".leave_remaining").
		From(dbUserTypeLeave.TableName()).
		InnerJoin(dbType.TableName()).
		On(dbType.TableName() + ".id" + " = " + dbUserTypeLeave.TableName() + ".type_leave_id").
		Where(dbUserTypeLeave.TableName() + `.type_leave_id = ? `).And(dbUserTypeLeave.TableName() + `.employee_number = ? `)
	sql := qb.String()

	errRaw := o.Raw(sql, typeID, employeeNumber).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetUserLeaveRemaining", errRaw)
		return result, errors.New("Error get user leave balance")
	}

	return result, err
}

// GetSupervisors ...
func (u *User) GetSupervisors() (result []structLogic.GetSupervisors, err error) {
	var dbSupervisor structDB.User

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetSupervisors", errQB)
		return result, errQB
	}

	qb.Select(
		dbSupervisor.TableName()+".employee_number",
		dbSupervisor.TableName()+".name").
		From(dbSupervisor.TableName()).
		Where(`role = ? `)
	sql := qb.String()
	role := "supervisor"

	_, errRaw := o.Raw(sql, role).QueryRows(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetSupervisors", errRaw)
		return result, errors.New("Error get supervisor")
	}

	return result, err
}

// GetTypeLeave ...
func (u *User) GetTypeLeave() (result []structDB.TypeLeave, err error) {
	var dbType structDB.TypeLeave

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetTypeLeave", errQB)
		return result, errQB
	}

	qb.Select(
		dbType.TableName()+".id",
		dbType.TableName()+".type_name").
		From(dbType.TableName())
	sql := qb.String()

	_, errRaw := o.Raw(sql).QueryRows(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetTypeLeave", errRaw)
		return result, errors.New("Error get type leave")
	}

	return result, err
}
