package employee

import (
	"server/helpers"

	structLogic "server/structs/logic"
)

// GetPendingRequest ...
func GetPendingRequest(employeeNumber int64) ([]structLogic.RequestPending, error) {
	respGet, errGet := DBEmployee.GetPendingRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetPendingRequest - logicEmployee", errGet)
	}

	return respGet, errGet
}

// GetApprovedRequest ...
func GetApprovedRequest(employeeNumber int64) ([]structLogic.RequestAccept, error) {
	respGet, errGet := DBEmployee.GetApprovedRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get approved request @GetApprovedRequest - logicEmployee", errGet)
	}

	return respGet, errGet
}

// GetRejectedRequest ...
func GetRejectedRequest(employeeNumber int64) ([]structLogic.RequestReject, error) {
	respGet, errGet := DBEmployee.GetRejectedRequest(employeeNumber)
	if errGet != nil {
		helpers.CheckErr("Error get rejected request @GetRejectedRequest - logicEmployee", errGet)
	}

	return respGet, errGet
}
