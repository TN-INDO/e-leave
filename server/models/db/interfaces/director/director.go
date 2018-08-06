package director

import (
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseDirector ...
type IBaseDirector interface {
	// AcceptByDirector
	AcceptByDirector(
		id int64,
		employeeNumber int64,
	) (err error)
	// RejectByDirector
	RejectByDirector(
		l *structDB.LeaveRequest,
		id int64,
		employeeNumber int64,
	) error
	// GetDirectorPendingRequest
	GetDirectorPendingRequest() (
		[]structLogic.RequestPending,
		error,
	)
	// GetDirectorAcceptRequest
	GetDirectorAcceptRequest() (
		[]structLogic.RequestAccept,
		error,
	)
	// GetDirectorRejectRequest
	GetDirectorRejectRequest() (
		[]structLogic.RequestReject,
		error,
	)
	// CancelRequestLeave
	CancelRequestLeave(
		id int64,
		employeeNumber int64,
	) (err error)
}
