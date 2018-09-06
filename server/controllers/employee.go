package controllers

import (
	"errors"
	"server/helpers"
	"strconv"

	logicEmployee "server/models/logic/employee"
	structAPI "server/structs/api"

	"github.com/astaxie/beego"
)

//EmployeeController ...
type EmployeeController struct {
	beego.Controller
}

// GetRequestPending ...
func (c *EmployeeController) GetRequestPending() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetRequestPending - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetPend := logicEmployee.GetPendingRequest(employeeNumber)
	if errGetPend != nil {
		resp.Error = errGetPend.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetRequestPending - controller", err)
	}
}

// GetRequestAccept ...
func (c *EmployeeController) GetRequestAccept() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetRequestAccept - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetAccept := logicEmployee.GetApprovedRequest(employeeNumber)
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetRequestAccept - controller", err)
	}
}

// GetRequestReject ...
func (c *EmployeeController) GetRequestReject() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetRequestReject - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetReject := logicEmployee.GetRejectedRequest(employeeNumber)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetRequestReject - controller", err)
	}
}
