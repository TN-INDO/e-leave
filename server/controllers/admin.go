package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"strconv"

	logic "server/models/logic/user"
	structAPI "server/structs/api"
	structDB "server/structs/db"

	"github.com/astaxie/beego"
)

// AdminController ...
type AdminController struct {
	beego.Controller
}

// CreateUser ...
func (c *AdminController) CreateUser() {
	var (
		resp    structAPI.RespData
		reqUser structAPI.ReqRegister
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-USER=======>", string(body))

	errMarshal := json.Unmarshal(body, &reqUser)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @CreateUser", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	user := structDB.User{
		EmployeeNumber:   reqUser.EmployeeNumber,
		Name:             reqUser.Name,
		Gender:           reqUser.Gender,
		Position:         reqUser.Position,
		StartWorkingDate: reqUser.StartWorkingDate,
		MobilePhone:      reqUser.MobilePhone,
		Email:            reqUser.Email,
		Password:         reqUser.Password,
		Role:             reqUser.Role,
		SupervisorID:     reqUser.SupervisorID,
	}

	errAddUser := logic.DBPostAdmin.AddUser(user)
	if errAddUser != nil {
		resp.Error = errAddUser.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = "create user success"
	}

	// startDate := helpers.GetDay(user.StartWorkingDate)
	// annualLeave := 12 + (12 - startDate)
	// beego.Debug("==>", annualLeave)

	if reqUser.Gender == "Male" && reqUser.Role == "employee" || reqUser.Role == "supervisor" {
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 11, 12)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 22, 3)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 33, 30)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 44, 2)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 66, 2)
	} else if reqUser.Gender == "Female" && reqUser.Role == "employee" || reqUser.Role == "supervisor" {
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 11, 12)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 22, 3)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 33, 30)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 44, 2)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 55, 90)
		logic.DBPostAdmin.CreateUserTypeLeave(user.EmployeeNumber, 66, 2)
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @CreateUser", err)
	}
}

// DeleteUser ...
func (c *AdminController) DeleteUser() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @DeleteUser", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	if err := logic.DBPostAdmin.DeleteUser(employeeNumber); err == nil {
		resp.Body = "Deleted success"
	} else {
		resp.Error = err.Error()
		c.Ctx.Output.SetStatus(400)
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @DeleteUser", err)
	}
}

// GetUsers ...
func (c *AdminController) GetUsers() {
	var resp structAPI.RespData

	res, errGet := logic.DBPostAdmin.GetUsers()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetUsers", err)
	}
}

// GetUser ...
func (c *AdminController) GetUser() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetRequestAccept", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	res, errGet := logic.DBPostAdmin.GetUser(employeeNumber)
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetUsers", err)
	}
}

// UpdateUser ...
func (c *AdminController) UpdateUser() {
	var (
		resp    structAPI.RespData
		reqUser structAPI.ReqRegister
	)

	body := c.Ctx.Input.RequestBody
	// fmt.Println("UPDATE-USER=======>", string(body))

	err := json.Unmarshal(body, &reqUser)
	if err != nil {
		helpers.CheckErr("unmarshall req body failed @UpdateUser", err)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @UpdateUser", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resTime, errTime := helpers.NowLoc("Asia/Jakarta")
	helpers.CheckErr("err time", errTime)

	user := structDB.User{
		EmployeeNumber:   reqUser.EmployeeNumber,
		Name:             reqUser.Name,
		Gender:           reqUser.Gender,
		Position:         reqUser.Position,
		StartWorkingDate: reqUser.StartWorkingDate,
		MobilePhone:      reqUser.MobilePhone,
		Email:            reqUser.Email,
		Password:         reqUser.Password,
		Role:             reqUser.Role,
		SupervisorID:     reqUser.SupervisorID,
		UpdatedAt:        resTime,
	}

	errUpdate := logic.DBPostAdmin.UpdateUser(&user, employeeNumber)
	if errUpdate != nil {
		resp.Error = errUpdate.Error()
	} else {
		resp.Body = "Update user success"
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @UpdateUser", err)
	}
}

// GetRequestPending ...
func (c *AdminController) GetRequestPending() {
	var resp structAPI.RespData

	resGet, errGetPending := logic.DBPostAdmin.GetLeaveRequestPending()
	if errGetPending != nil {
		resp.Error = errGetPending.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestPending", err)
	}
}

// GetRequestAccept ...
func (c *AdminController) GetRequestAccept() {
	var resp structAPI.RespData

	resGet, errGetAccept := logic.DBPostAdmin.GetLeaveRequest()
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestAccept", err)
	}
}

// GetRequestReject ...
func (c *AdminController) GetRequestReject() {
	var resp structAPI.RespData

	resGet, errGetReject := logic.DBPostAdmin.GetLeaveRequestReject()
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestReject", err)
	}
}
