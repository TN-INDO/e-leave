package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"strconv"

	logicSupervisor "server/models/logic/supervisor"
	structAPI "server/structs/api"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
)

// SupervisorController ...
type SupervisorController struct {
	beego.Controller
}

// GetPendingLeave ...
func (c *SupervisorController) GetPendingLeave() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetPendingLeave - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetPend := logicSupervisor.GetEmployeePending(supervisorID)
	if errGetPend != nil {
		resp.Error = errGetPend.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetPendingLeave - controller", err)
	}
}

// GetAcceptLeave ...
func (c *SupervisorController) GetAcceptLeave() {
	var (
		resp structAPI.RespData
	)

	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetAcceptLeave - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetAccept := logicSupervisor.GetEmployeeApproved(supervisorID)
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetAcceptLeave - controller", err)
	}
}

// GetRejectLeave ...
func (c *SupervisorController) GetRejectLeave() {
	var (
		resp structAPI.RespData
	)

	idStr := c.Ctx.Input.Param(":id")
	supervisorID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetRejectLeave - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetReject := logicSupervisor.GetEmployeeRejected(supervisorID)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetRejectLeave - controller", err)
	}
}

// AcceptStatusBySupervisor ...
func (c *SupervisorController) AcceptStatusBySupervisor() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @AcceptStatusBySupervisor - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert enum failed @AcceptStatusBySupervisor - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	errUpStat := logicSupervisor.ApproveBySupervisor(id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "Leave request has been approved"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @AcceptStatusBySupervisor - controller", err)
	}
}

// RejectStatusBySv ...
func (c *SupervisorController) RejectStatusBySv() {
	var (
		resp  structAPI.RespData
		leave structLogic.LeaveReason
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("REJECT-REASON=======>", string(body))

	errMarshal := json.Unmarshal(body, &leave)
	if errMarshal != nil {
		helpers.CheckErr("Unmarshall req body failed @RejectStatusBySv - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @RejectStatusBySv - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	employeeStr := c.Ctx.Input.Param(":enumber")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert employee number failed @RejectStatusBySv - controller", errCon)
		resp.Error = errors.New("convert employee number failed").Error()
		return
	}

	errUpStat := logicSupervisor.RejectBySupervisor(&leave, id, employeeNumber)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "Leave request has been rejected"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @RejectStatusBySv - controller", err)
	}
}
