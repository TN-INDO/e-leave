package controllers

import (
	"encoding/json"
	"errors"
	"server/helpers"
	"strconv"

	logic "server/models/logic/user"
	structAPI "server/structs/api"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
)

//UserController ...
type UserController struct {
	beego.Controller
}

// Login ...
func (c *UserController) Login() {
	var (
		resp     structAPI.RespData
		reqLogin structAPI.ReqLogin
	)

	body := c.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &reqLogin)
	if err != nil {
		helpers.CheckErr("unmarshall req body failed @Login", err)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	result, errLogin := logic.DBPostUser.GetJWT(&reqLogin)
	if errLogin != nil {
		resp.Error = errLogin.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = result
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @Login", err)
	}
}

// PasswordReset ...
func (c *UserController) PasswordReset() {
	var (
		resp   structAPI.RespData
		dbUser structLogic.PasswordReset
	)

	body := c.Ctx.Input.RequestBody

	errMarshal := json.Unmarshal(body, &dbUser)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @PasswordReset", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	errUpStat := logic.DBPostUser.ForgotPassword(&dbUser)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "reset password success, please check your email"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @PasswordReset", err)
	}
}

// UpdateNewPassword ...
func (c *UserController) UpdateNewPassword() {
	var (
		resp   structAPI.RespData
		newPwd structLogic.NewPassword
	)

	body := c.Ctx.Input.RequestBody

	errMarshal := json.Unmarshal(body, &newPwd)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @UpdateNewPassword", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	employeeStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(employeeStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert enum failed @UpdateNewPassword", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	errUpPassword := logic.DBPostUser.UpdatePassword(&newPwd, employeeNumber)
	if errUpPassword != nil {
		resp.Error = errUpPassword.Error()
	} else {
		resp.Body = "Update password success"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @UpdateNewPassword", err)
	}
}

// GetUserSummary ...
func (c *UserController) GetUserSummary() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetUserSummary", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGetSummary := logic.DBPostUser.GetSumarry(employeeNumber)
	if errGetSummary != nil {
		resp.Error = errGetSummary.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetUserSummary", err)
	}
}

// GetUserTypeLeave ...
func (c *UserController) GetUserTypeLeave() {
	var (
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @GetUserTypeLeave", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	resGet, errGet := logic.DBPostUser.GetUserTypeLeave(employeeNumber)
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetUserTypeLeave", err)
	}
}

// GetSupervisors ...
func (c *UserController) GetSupervisors() {
	var resp structAPI.RespData

	res, errGet := logic.DBPostUser.GetSupervisors()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetSupervisors", err)
	}
}

// GetTypeLeave ...
func (c *UserController) GetTypeLeave() {
	var resp structAPI.RespData
	res, errGet := logic.DBPostUser.GetTypeLeave()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetTypeLeave", err)
	}
}
