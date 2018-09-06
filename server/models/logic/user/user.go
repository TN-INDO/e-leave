package user

import (
	"encoding/base64"
	"errors"
	"server/helpers"
	"time"

	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"golang.org/x/crypto/bcrypt"
)

// UserLogin ...
func UserLogin(loginData *structAPI.ReqLogin) (respLogin structAPI.RespLogin, err error) {

	respGet, errGet := DBPostUser.UserLogin(loginData)
	if errGet != nil {
		helpers.CheckErr("Error get user login @UserLogin - logicUser", errGet)
		return respLogin, errGet
	}

	if errGet == nil {
		hashBytes, errDecode := base64.StdEncoding.DecodeString(respGet.Password)
		helpers.CheckErr("Error decode password @UserLogin - logicUser", errDecode)

		errCompare := bcrypt.CompareHashAndPassword(hashBytes, []byte(loginData.Password))
		if errCompare != nil {
			helpers.CheckErr("Error compare password @UserLogin - logicUser", errCompare)
			return respLogin, errors.New("Wrong password")
		}

		ezT := helpers.EzToken{
			ID:      respGet.EmployeeNumber,
			Email:   respGet.Email,
			Expires: time.Now().Unix() + 3600,
		}

		token, err := ezT.GetToken()
		if err != nil {
			helpers.CheckErr("Error get token @UserLogin - logicUser", err)
			return respLogin, errors.New("Failed generating token")
		}

		respLogin.Token = token
		respLogin.ID = respGet.EmployeeNumber
		respLogin.Role = respGet.Role
	}

	return respLogin, err
}

// ForgotPassword ...
func ForgotPassword(e *structLogic.PasswordReset) error {

	respCount, errCount := DBPostUser.CountUserEmail(e.Email)
	if errCount != nil {
		helpers.CheckErr("Error count user @ForgotPassword - logicUser", errCount)
	}

	respGet, errGet := DBPostUser.GetUser(e.Email)
	if errGet != nil {
		helpers.CheckErr("Error get user @ForgotPassword - logicUser", errGet)
	}

	if respCount == 0 {
		return errors.New("Email not register")
	}

	errUp := DBPostUser.ForgotPassword(e)
	if errUp != nil {
		helpers.CheckErr("Error forgot password @ForgotPassword - logicUser", errUp)
		return errUp
	}

	go func() {
		helpers.GoMailForgotPassword(respGet.Email, respGet.Name)
	}()

	return errUp
}

// GetDirector ...
func GetDirector() (structLogic.GetDirector, error) {
	respGet, errGet := DBPostUser.GetDirector()
	if errGet != nil {
		helpers.CheckErr("Error get director @GetDirector - logicUser", errGet)
	}

	return respGet, errGet
}

// GetSupervisors ...
func GetSupervisors() ([]structLogic.GetSupervisors, error) {
	respGet, errGet := DBPostUser.GetSupervisors()
	if errGet != nil {
		helpers.CheckErr("Error get supervisors @GetSupervisors - logicUser", errGet)
	}

	return respGet, errGet
}

// GetSupervisor ...
func GetSupervisor(employeeNumber int64) (structLogic.GetSupervisor, error) {
	respGet, errGet := DBPostUser.GetSupervisor(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get supervisor @GetSupervisor - logicUser", errGet)
	}

	return respGet, errGet
}

// GetEmployee ...
func GetEmployee(employeeNumber int64) (structLogic.GetEmployee, error) {
	respGet, errGet := DBPostUser.GetEmployee(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get employee @GetEmployee - logicUser", errGet)
	}

	return respGet, errGet
}

// CreateUserTypeLeave ...
func CreateUserTypeLeave(
	employeeNumber int64,
	typeLeaveID int64,
	leaveRemaining float64,
) error {
	errInsert := DBPostUser.CreateUserTypeLeave(employeeNumber, typeLeaveID, leaveRemaining)
	if errInsert != nil {
		helpers.CheckErr("Error insert user type leave @CreateUserTypeLeave - logicUser", errInsert)
	}

	return errInsert
}

// GetTypeLeave ...
func GetTypeLeave() ([]structDB.TypeLeave, error) {
	respGet, errGet := DBPostUser.GetTypeLeave()
	if errGet != nil {
		helpers.CheckErr("Error get type leave @GetTypeLeave - logicUser", errGet)
	}

	return respGet, errGet
}

// GetUserTypeLeave ...
func GetUserTypeLeave(employeeNumber int64) ([]structLogic.UserTypeLeave, error) {
	respGet, errGet := DBPostUser.GetUserTypeLeave(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get user type leave @GetUserTypeLeave - logicUser", errGet)
	}

	return respGet, errGet
}

// GetSumarry ...
func GetSumarry(employeeNumber int64) ([]structLogic.UserSumarry, error) {
	respGet, errGet := DBPostUser.GetSumarry(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get summary @GetSumarry - logicUser", errGet)
	}

	return respGet, errGet
}

// GetUserLeaveRemaining ...
func GetUserLeaveRemaining(typeID int64, employeeNumber int64) (structLogic.UserTypeLeave, error) {
	respGet, errGet := DBPostUser.GetUserLeaveRemaining(typeID, employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get leave remaining @GetUserLeaveRemaining - logicUser", errGet)
	}

	return respGet, errGet
}
