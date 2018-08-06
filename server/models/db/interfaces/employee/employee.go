package employee

import (
	structLogic "server/structs/logic"
)

// IBaseEmployee ...
type IBaseEmployee interface {
	// GetPendingRequest
	GetPendingRequest(employeeNumber int64) (
		[]structLogic.RequestPending,
		error,
	)
	// GetApprovedRequest
	GetApprovedRequest(employeeNumber int64) (
		[]structLogic.RequestAccept,
		error,
	)
	// GetRejectedRequest
	GetRejectedRequest(employeeNumber int64) (
		[]structLogic.RequestReject,
		error,
	)
}
