package supervisor

import (
	structLogic "server/structs/logic"
)

// IBaseSupervisor ...
type IBaseSupervisor interface {
	// GetUserPending
	GetEmployeePending(supervisorID int64) (
		reqPending []structLogic.LeavePending,
		err error,
	)
	// GetEmployeeApproved
	GetEmployeeApproved(supervisorID int64) (
		reqApprove []structLogic.LeaveAccept,
		err error,
	)
	// GetEmployeeRejected
	GetEmployeeRejected(supervisorID int64) (
		reqReject []structLogic.LeaveReject,
		err error,
	)
	// ApproveBySupervisor
	ApproveBySupervisor(
		id int64,
		employeeNumber int64,
		actionBy string,
	) error
	// RejectBySupervisor
	RejectBySupervisor(
		l *structLogic.LeaveReason,
		id int64,
		employeeNumber int64,
		actionBy string,
	) error
}
