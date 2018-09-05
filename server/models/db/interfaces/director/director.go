package director

import (
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseDirector ...
type IBaseDirector interface {
	// GetEmployeePending
	GetEmployeePending() (
		[]structLogic.RequestPending,
		error,
	)
	// GetEmployeeApproved
	GetEmployeeApproved() (
		[]structLogic.RequestAccept,
		error,
	)
	// GetEmployeeRejected
	GetEmployeeRejected() (
		[]structLogic.RequestReject,
		error,
	)

	// ApproveByDirector
	ApproveByDirector(
		id int64,
		employeeNumber int64,
		actionBy string,
	) (err error)
	// RejectByDirector
	RejectByDirector(
		l *structDB.LeaveRequest,
		id int64,
		employeeNumber int64,
		actionBy string,
	) error
}
