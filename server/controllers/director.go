package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"strconv"

	logicDirector "server/models/logic/director"
	structAPI "server/structs/api"
	structDB "server/structs/db"

	"github.com/astaxie/beego"
)

// DirectorController ...
type DirectorController struct {
	beego.Controller
}

// GetDirectorPendingLeave ...
func (c *DirectorController) GetDirectorPendingLeave() {
	var resp structAPI.RespData

	resGet, errGetPend := logicDirector.GetEmployeePendingRequest()
	if errGetPend != nil {
		resp.Error = errGetPend.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetDirectorPendingLeave - controller", err)
	}
}

// GetDirectorAcceptLeave ...
func (c *DirectorController) GetDirectorAcceptLeave() {
	var resp structAPI.RespData

	resGet, errGetAcc := logicDirector.GetEmployeeApprovedRequest()
	if errGetAcc != nil {
		resp.Error = errGetAcc.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetDirectorPendingLeave - controller", err)
	}
}

// GetDirectorRejectLeave ...
func (c *DirectorController) GetDirectorRejectLeave() {
	var resp structAPI.RespData

	resGet, errGetReject := logicDirector.GetEmployeeRejectedRequest()
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetDirectorPendingLeave - controller", err)
	}
}

// AcceptStatusByDirector ...
func (c *DirectorController) AcceptStatusByDirector() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @AcceptStatusByDirector - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert employee number failed @AcceptStatusByDirector - controller", errCon)
		resp.Error = errors.New("Convert employee number failed").Error()
		return
	}

	errUpStat := logicDirector.ApproveByDirector(id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "Leave request has been approved"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @AcceptStatusByDirector - controller", err)
	}
}

// RejectStatusByDirector ...
func (c *DirectorController) RejectStatusByDirector() {
	var (
		resp  structAPI.RespData
		leave structDB.LeaveRequest
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("REJECT-REASON=======>", string(body))

	errMarshal := json.Unmarshal(body, &leave)
	if errMarshal != nil {
		helpers.CheckErr("Failed unmarshall req body @RejectStatusByDirector - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @RejectStatusByDirector - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert employee number failed @RejectStatusByDirector - controller", errCon)
		resp.Error = errors.New("Convert employee number failed").Error()
		return
	}

	errUpStat := logicDirector.RejectByDirector(&leave, id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "Leave request has been rejected"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @RejectStatusByDirector - controller", err)
	}
}
