package employee

import (
	"server/helpers"
	structLogic "server/structs/logic"
)

// GetPendingRequest ...
func GetPendingRequest(employeeNumber int64) ([]structLogic.RequestPending, error) {
	respGet, errGet := DBPostEmployee.GetPendingRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("err get pending request", errGet)
	}

	return respGet, errGet
}

// GetApprovedRequest ...
func GetApprovedRequest(employeeNumber int64) ([]structLogic.RequestAccept, error) {
	respGet, errGet := DBPostEmployee.GetApprovedRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("err get approved request", errGet)
	}

	return respGet, errGet
}

// GetRejectedRequest ...
func GetRejectedRequest(employeeNumber int64) ([]structLogic.RequestReject, error) {
	respGet, errGet := DBPostEmployee.GetRejectedRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("err get rejected request", errGet)
	}

	return respGet, errGet
}
