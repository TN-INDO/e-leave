package employee

import (
	structLogic "server/structs/logic"
)

// IBaseEmployee ...
type IBaseEmployee interface {
	// GetPendingRequest
	GetPendingRequest(employeeNumber int64) (
		reqPending []structLogic.RequestPending,
		err error,
	)
	// GetApprovedRequest
	GetApprovedRequest(employeeNumber int64) (
		reqApprove []structLogic.RequestAccept,
		err error,
	)
	// GetRejectedRequest
	GetRejectedRequest(employeeNumber int64) (
		reqReject []structLogic.RequestReject,
		err error,
	)
}
