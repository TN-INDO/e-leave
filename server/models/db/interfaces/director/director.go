package director

import (
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseDirector ...
type IBaseDirector interface {
	// GetEmployeePending
	GetEmployeePending() (
		reqPending []structLogic.RequestPending,
		err error,
	)
	// GetEmployeeApproved
	GetEmployeeApproved() (
		reqApprove []structLogic.RequestAccept,
		err error,
	)
	// GetEmployeeRejected
	GetEmployeeRejected() (
		reqReject []structLogic.RequestReject,
		err error,
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
