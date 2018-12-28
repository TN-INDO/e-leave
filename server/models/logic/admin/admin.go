package admin

import (
	"encoding/base64"
	"errors"
	"server/helpers"
	"strconv"
	"strings"

	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser ...
func CreateUser(user structDB.User) error {

	resCountEmployee, errCountEmployee := DBUser.CountUserEmployeeNumber(user.EmployeeNumber)
	helpers.CheckErr("Error get users @GetUsers - logicAdmin", errCountEmployee)

	resCountEmail, errCountEmail := DBUser.CountUserEmail(user.Email)
	helpers.CheckErr("Error get users @GetUsers - logicAdmin", errCountEmail)

	passwordString := user.Password
	bsEmployeeNumber := []byte(strconv.Itoa(int(user.EmployeeNumber)))
	arrPassword := []byte(passwordString)

	employeeNumberStr := strconv.FormatInt(user.EmployeeNumber, 10)

	if employeeNumberStr == "" || user.Name == "" || user.Gender == "" || user.Position == "" || user.StartWorkingDate == "" || user.MobilePhone == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("Error empty field ")
	}
	if len(bsEmployeeNumber) != 5 {
		return errors.New("Employee number must length must be 5")
	}
	if resCountEmployee > 0 {
		return errors.New("Employee number already register")
	}
	if resCountEmail > 0 {
		return errors.New("Email already register")
	}
	if len(arrPassword) < 7 {
		return errors.New("Password length must be 7")
	}

	hashedBytes, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	helpers.CheckErr("Error hash password @CreateUser", errHash)

	user.Email = strings.ToLower(user.Email)
	user.Password = base64.StdEncoding.EncodeToString(hashedBytes)

	errInsert := DBAdmin.AddUser(user)
	if errInsert != nil {
		helpers.CheckErr("Error insert @CreateUser", errInsert)
		return errInsert
	}

	go func() {
		helpers.GoMailRegisterPassword(user.Email, passwordString)
	}()

	return errInsert
}

// GetUsers ...
func GetUsers() ([]structDB.User, error) {
	respGet, errGet := DBAdmin.GetUsers()
	if errGet != nil {
		helpers.CheckErr("Error get users @GetUsers - logicAdmin", errGet)
	}

	return respGet, errGet
}

// GetUser ...
func GetUser(employeeNumber int64) (structDB.User, error) {
	respGet, errGet := DBAdmin.GetUser(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get user @GetUser - logicAdmin", errGet)
	}

	return respGet, errGet
}

// DeleteUser ...
func DeleteUser(employeeNumber int64) error {
	errDelete := DBAdmin.DeleteUser(employeeNumber)
	if errDelete != nil {
		helpers.CheckErr("Error delete user @DeleteUser - logicAdmin", errDelete)
		return errDelete
	}

	return errDelete
}

// GetLeaveRequestPending ...
func GetLeaveRequestPending() ([]structLogic.RequestPending, error) {
	respGet, errGet := DBAdmin.GetLeaveRequestPending()
	if errGet != nil {
		helpers.CheckErr("Error get leave request pending @GetLeaveRequestPending - logicAdmin", errGet)
	}

	return respGet, errGet
}

// GetLeaveRequestApproved ...
func GetLeaveRequestApproved() ([]structLogic.RequestAccept, error) {
	respGet, errGet := DBAdmin.GetLeaveRequestApproved()
	if errGet != nil {
		helpers.CheckErr("Error get leave request approve @GetLeaveRequestApproved - logicAdmin", errGet)
	}

	return respGet, errGet
}

// GetLeaveRequestRejected ...
func GetLeaveRequestRejected() ([]structLogic.RequestReject, error) {
	respGet, errGet := DBAdmin.GetLeaveRequestRejected()
	if errGet != nil {
		helpers.CheckErr("Error get leave request reject @GetLeaveRequestRejected - logicAdmin", errGet)
	}

	return respGet, errGet
}

// CancelRequestLeave ...
func CancelRequestLeave(id int64, employeeNumber int64) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		helpers.CheckErr("Error begin @CancelRequestLeave - logicAdmin", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	getDirector, errGetDirector := DBUser.GetDirector()
	helpers.CheckErr("Error get director @CancelRequestLeave - logicAdmin", errGetDirector)

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	helpers.CheckErr("Error get employee @CancelRequestLeave - logicAdmin", errGetEmployee)

	getLeave, errGetLeave := DBLeave.GetLeave(id)
	helpers.CheckErr("Error get leave @CancelRequestLeave - logicAdmin", errGetLeave)

	errUp := DBLeave.UpdateLeaveRemaningCancel(getLeave.Total, employeeNumber, getLeave.TypeLeaveID)
	if errUp != nil {
		helpers.CheckErr("Error update cancel leave request @CancelRequestLeave - logicAdmin", errUp)
		o.Rollback()
		return errUp
	}

	errDelete := DBLeave.DeleteRequest(id)
	if errDelete != nil {
		helpers.CheckErr("Error delete leave request @CancelRequestLeave - logicAdmin", errDelete)
		o.Rollback()
		return errDelete
	}

	err = o.Commit()
	if err != nil {
		helpers.CheckErr("Error commit @CancelRequestLeave - logicAdmin", err)
		o.Rollback()
		return errors.New("Failed transaction fench")
	}

	go func() {
		helpers.GoMailDirectorCancel(getDirector.Email, getLeave.ID, getEmployee.Name, getDirector.Name)
		helpers.GoMailEmployeeCancel(getEmployee.Email, getLeave.ID, getEmployee.Name)
	}()

	return err
}

// ResetUserTypeLeave ...
func ResetUserTypeLeave(leaveRemaining float64, typeLeaveID int64) error {
	errReset := DBAdmin.ResetUserTypeLeave(leaveRemaining, typeLeaveID)
	if errReset != nil {
		helpers.CheckErr("Error reset type leave @ResetUserTypeLeave - logicAdmin", errReset)
	}

	return errReset
}
