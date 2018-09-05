package user

import (
	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseUser ...
type IBaseUser interface {
	// UserLogin
	UserLogin(loginData *structAPI.ReqLogin) (
		result structDB.User,
		err error,
	)
	// ForgotPassword
	ForgotPassword(e *structLogic.PasswordReset) error
	CountUserEmail(email string) (int, error)
	GetUser(email string) (structLogic.GetEmployee, error)

	// UpdatePassword
	UpdatePassword(
		p *structLogic.NewPassword,
		employeeNumber int64,
	) (err error)

	// GetDirector
	GetDirector() (
		result structLogic.GetDirector,
		err error,
	)
	// GetSupervisors
	GetSupervisors() (
		result []structLogic.GetSupervisors,
		err error,
	)
	// GetEmployee
	GetEmployee(employeeNumber int64) (
		result structLogic.GetEmployee,
		err error,
	)

	// GetTypeLeave
	GetTypeLeave() (
		result []structDB.TypeLeave,
		err error,
	)
	// GetSumarry
	GetSumarry(employeeNumber int64) (
		[]structLogic.UserSumarry,
		error,
	)
	// GetUserLeaveRemaining
	GetUserLeaveRemaining(
		typeID int64,
		employeeNumber int64,
	) (
		result structLogic.UserTypeLeave,
		err error,
	)

	// GetSupervisor
	GetSupervisor(employeeNumber int64) (
		result structLogic.GetSupervisor,
		err error,
	)
	// GetUserTypeLeave
	GetUserTypeLeave(employeeNumber int64) (
		result []structLogic.UserTypeLeave,
		err error,
	)
}
