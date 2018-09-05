package user

import (
	"encoding/base64"
	"errors"
	"server/helpers"
	"time"

	"golang.org/x/crypto/bcrypt"

	structAPI "server/structs/api"
	structLogic "server/structs/logic"
)

// UserLogin ...
func UserLogin(loginData *structAPI.ReqLogin) (respLogin structAPI.RespLogin, err error) {

	respGet, errGet := DBPostUser.UserLogin(loginData)
	if errGet != nil {
		helpers.CheckErr("Error get user login @UserLogin - logicUser", errGet)
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
		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errCount)
	}

	respGet, errGet := DBPostUser.GetUser(e.Email)
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errGet)
	}

	if respCount == 0 {
		return errors.New("email not register")
	}

	errUp := DBPostUser.ForgotPassword(e)
	if errUp != nil {
		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errUp)
	}

	go func() {
		helpers.GoMailForgotPassword(respGet.Email, respGet.Name)
	}()

	return errUp
}

// // GetEmployeePendingRequest ...
// func GetEmployeePendingRequest() ([]structLogic.RequestPending, error) {
// 	respGet, errGet := DBAdmin.GetEmployeePending()
// 	if errGet != nil {
// 		helpers.CheckErr("Error get pending request @GetEmployeePendingRequest - logicDirector", errGet)
// 	}

// 	return respGet, errGet
// }

// // GetEmployeeApprovedRequest ...
// func GetEmployeeApprovedRequest() ([]structLogic.RequestAccept, error) {
// 	respGet, errGet := DBAdmin.GetEmployeeApproved()
// 	if errGet != nil {
// 		helpers.CheckErr("Error get approved request @GetEmployeeApprovedRequest - logicDirector", errGet)
// 	}

// 	return respGet, errGet
// }

// // GetEmployeeRejectedRequest ...
// func GetEmployeeRejectedRequest() ([]structLogic.RequestReject, error) {
// 	respGet, errGet := DBAdmin.GetEmployeeRejected()
// 	if errGet != nil {
// 		helpers.CheckErr("Error get rejected request @GetEmployeeRejectedRequest - logicDirector", errGet)
// 	}

// 	return respGet, errGet
// }
