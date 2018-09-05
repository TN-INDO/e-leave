package supervisor

import (
	structLogic "server/structs/logic"
)

// IBaseSupervisor ...
type IBaseSupervisor interface {
	// GetUserPending
	GetEmployeePending(supervisorID int64) (
		[]structLogic.LeavePending,
		error,
	)
	// GetEmployeeApproved
	GetEmployeeApproved(supervisorID int64) (
		[]structLogic.LeaveAccept,
		error,
	)
	// GetEmployeeRejected
	GetEmployeeRejected(supervisorID int64) (
		[]structLogic.LeaveReject,
		error,
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
