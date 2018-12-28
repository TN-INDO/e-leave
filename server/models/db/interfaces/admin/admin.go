package admin

import (
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseAdmin ...
type IBaseAdmin interface {
	// AddUser
	AddUser(user structDB.User) error
	// GetUsers
	GetUsers() (
		users []structDB.User,
		err error,
	)
	// GetUser
	GetUser(employeeNumber int64) (
		result structDB.User,
		err error,
	)
	// UpdateUser
	UpdateUser(
		e *structDB.User,
		employeeNumber int64,
	) (err error)
	// DeleteUser
	DeleteUser(employeeNumber int64) error

	// GetLeaveRequest
	GetLeaveRequestPending() (
		reqPending []structLogic.RequestPending,
		err error,
	)
	// GetLeaveRequestApproved
	GetLeaveRequestApproved() (
		reqApprove []structLogic.RequestAccept,
		err error,
	)
	// GetLeaveRequested
	GetLeaveRequestRejected() (
		reqReject []structLogic.RequestReject,
		err error,
	)

	// ResetUserTypeLeave
	ResetUserTypeLeave(
		leaveRemaining float64,
		typeLeaveID int64,
	) error
	// UpdateUserTypeLeave
	UpdateUserTypeLeave(
		leaveRemaining float64,
		typeLeaveID int64,
		employeeNumber int64,
	) error
}
