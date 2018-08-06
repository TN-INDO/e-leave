package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	"server/helpers/constant"
	"strconv"
	"time"

	db "server/models/db/pgsql/leave_request"
	userLogic "server/models/db/pgsql/user"

	structAPI "server/structs/api"

	"github.com/astaxie/beego"
)

//LeaveController ...
type LeaveController struct {
	beego.Controller
}

// PostLeaveRequest ...
func (c *LeaveController) PostLeaveRequest() {
	var (
		resp    structAPI.RespData
		req     structAPI.ReqLeave
		dbLeave db.LeaveRequest
		dbUser  userLogic.User
	)

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @PostLeaveRequest", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-LEAVE-REQUEST=======>", string(body))

	errMarshal := json.Unmarshal(body, &req)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @PostLeaveRequest", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	totalDay := float64(helpers.GetTotalDay(req.DateFrom, req.DateTo))
	reqHalfDay := float64(len(req.HalfDates))
	valueHalfDay := float64(0.5)
	result := helpers.Multiply(totalDay, reqHalfDay, valueHalfDay)

	resGet, errGet := dbUser.GetUserLeaveRemaining(req.TypeLeaveID, employeeNumber)
	helpers.CheckErr("err get", errGet)

	leave := structAPI.CreateLeaveRequest{
		EmployeeNumber: employeeNumber,
		TypeLeaveID:    req.TypeLeaveID,
		Reason:         req.Reason,
		DateFrom:       req.DateFrom,
		DateTo:         req.DateTo,
		HalfDates:      req.HalfDates,
		Total:          result,
		BackOn:         req.BackOn,
		ContactAddress: req.ContactAddress,
		ContactNumber:  req.ContactNumber,
		Status:         constant.StatusPendingInSupervisor,
	}

	strBalance := strconv.FormatFloat(resGet.LeaveRemaining, 'f', 1, 64)
	strTotal := strconv.FormatFloat(result, 'f', 1, 64)

	if req.TypeLeaveID != 11 && req.TypeLeaveID != 22 && req.TypeLeaveID != 33 && req.TypeLeaveID != 44 && req.TypeLeaveID != 55 && req.TypeLeaveID != 66 {
		beego.Warning("error empty field type leave @PostLeaveRequest")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("error empty field").Error()
	} else if req.DateFrom == "" || req.DateTo == "" || req.BackOn == "" || req.ContactAddress == "" || req.ContactNumber == "" {
		beego.Warning("error empty field @PostLeaveRequest")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("error empty field").Error()
	} else if result > float64(resGet.LeaveRemaining) {
		beego.Warning("error leave balance @PostLeaveRequest")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("your total leave is " + strTotal + " day and your " + resGet.TypeName + " balance is " + strBalance + " day left").Error()
	} else {
		errAddLeave := dbLeave.CreateLeaveRequest(
			leave.EmployeeNumber,
			leave.TypeLeaveID,
			leave.Reason,
			leave.DateFrom,
			leave.DateTo,
			leave.HalfDates,
			leave.BackOn,
			leave.Total,
			leave.ContactAddress,
			leave.ContactNumber,
			leave.Status,
		)

		if errAddLeave != nil {
			resp.Error = errAddLeave.Error()
			c.Ctx.Output.SetStatus(400)
		} else {
			resp.Body = leave
		}
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @PostLeaveRequest", err)
	}
}

// PostLeaveRequestSupervisor ...
func (c *LeaveController) PostLeaveRequestSupervisor() {
	var (
		resp    structAPI.RespData
		req     structAPI.ReqLeave
		dbLeave db.LeaveRequest
		dbUser  userLogic.User
	)

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @PostLeaveRequestSupervisor", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-LEAVE-REQUEST=======>", string(body))

	errMarshal := json.Unmarshal(body, &req)
	if errMarshal != nil {
		helpers.CheckErr("unmarshall req body failed @PostLeaveRequestSupervisor", errMarshal)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	totalDay := float64(helpers.GetTotalDay(req.DateFrom, req.DateTo))
	reqHalfDay := float64(len(req.HalfDates))
	valueHalfDay := float64(0.5)
	result := helpers.Multiply(totalDay, reqHalfDay, valueHalfDay)

	resGet, errGet := dbUser.GetUserLeaveRemaining(req.TypeLeaveID, employeeNumber)
	helpers.CheckErr("err get", errGet)

	leave := structAPI.CreateLeaveRequest{
		EmployeeNumber: employeeNumber,
		TypeLeaveID:    req.TypeLeaveID,
		Reason:         req.Reason,
		DateFrom:       req.DateFrom,
		DateTo:         req.DateTo,
		HalfDates:      req.HalfDates,
		Total:          result,
		BackOn:         req.BackOn,
		ContactAddress: req.ContactAddress,
		ContactNumber:  req.ContactNumber,
		Status:         constant.StatusPendingInDirector,
	}

	strBalance := strconv.FormatFloat(resGet.LeaveRemaining, 'f', 1, 64)
	strTotal := strconv.FormatFloat(result, 'f', 1, 64)

	if req.TypeLeaveID != 11 && req.TypeLeaveID != 22 && req.TypeLeaveID != 33 && req.TypeLeaveID != 44 && req.TypeLeaveID != 55 && req.TypeLeaveID != 66 {
		beego.Warning("error empty field type leave @PostLeaveRequestSupervisor")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("error empty field").Error()
	} else if req.DateFrom == "" || req.DateTo == "" || req.BackOn == "" || req.ContactAddress == "" || req.ContactNumber == "" {
		beego.Warning("error empty field @PostLeaveRequestSupervisor")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("error empty field").Error()
	} else if result > float64(resGet.LeaveRemaining) {
		beego.Warning("error leave balance @PostLeaveRequestSupervisor")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("your total leave is " + strTotal + " day and your " + resGet.TypeName + " balance is " + strBalance + " day left").Error()
	} else {
		errAddLeave := dbLeave.CreateLeaveRequest(
			leave.EmployeeNumber,
			leave.TypeLeaveID,
			leave.Reason,
			leave.DateFrom,
			leave.DateTo,
			leave.HalfDates,
			leave.BackOn,
			leave.Total,
			leave.ContactAddress,
			leave.ContactNumber,
			leave.Status,
		)

		if errAddLeave != nil {
			resp.Error = errAddLeave.Error()
			c.Ctx.Output.SetStatus(400)
		} else {
			resp.Body = leave
		}
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @PostLeaveRequestSupervisor", err)
	}
}

// UpdateRequest ...
func (c *LeaveController) UpdateRequest() {
	var (
		resp    structAPI.RespData
		leave   structAPI.UpdateLeaveRequest
		dbLeave db.LeaveRequest
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("UPDATE-LEAVE-REQUEST=======>", string(body))

	err := json.Unmarshal(body, &leave)
	if err != nil {
		helpers.CheckErr("unmarshall req body failed @UpdateRequest", err)
		resp.Error = errors.New("type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @UpdateRequest", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	totalDay := float64(helpers.GetTotalDay(leave.DateFrom, leave.DateTo))
	reqHalfDay := float64(len(leave.HalfDates))
	valueHalfDay := float64(0.5)
	result := helpers.Multiply(totalDay, reqHalfDay, valueHalfDay)

	leave = structAPI.UpdateLeaveRequest{
		TypeLeaveID:    leave.TypeLeaveID,
		Reason:         leave.Reason,
		DateFrom:       leave.DateFrom,
		DateTo:         leave.DateTo,
		HalfDates:      leave.HalfDates,
		Total:          result,
		BackOn:         leave.BackOn,
		ContactAddress: leave.ContactAddress,
		ContactNumber:  leave.ContactNumber,
	}

	errUpdate := dbLeave.UpdateRequest(&leave, id)
	if errUpdate != nil {
		resp.Error = errUpdate.Error()
	} else {
		resp.Body = "Update leave success"
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @UpdateRequest", err)
	}
}

// DeleteRequest ...
func (c *LeaveController) DeleteRequest() {
	var (
		resp    structAPI.RespData
		dbLeave db.LeaveRequest
	)

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("convert id failed @DeleteRequest", errCon)
		resp.Error = errors.New("convert id failed").Error()
		return
	}

	if err := dbLeave.DeleteRequest(id); err == nil {
		resp.Body = "Deleted success"
	} else {
		resp.Error = err.Error()
		c.Ctx.Output.SetStatus(400)
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @DeleteRequest", err)
	}
}

// GetDownloadReportCSV ...
func (c *LeaveController) GetDownloadReportCSV() {
	var dbLeave db.LeaveRequest

	var reqDt = structAPI.RequestReport{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
	}

	dt := time.Now()
	fileName := "report_leave_request_" + dt.Format("20060102")
	path := constant.GOPATH + "/src/" + constant.GOAPP + "/storages/" + fileName + ".csv"

	errGet := dbLeave.DownloadReportCSV(&reqDt, path)
	if errGet != nil {
		beego.Debug("Error get csv @GetDownloadReportCSV", errGet)
	}

	// c.Ctx.Output.Header("Content-Disposition", "attachment; filename="+url.PathEscape(fileName+".csv"))
	// c.Ctx.Output.Header("Content-Description", "File Transfer")
	// c.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	// c.Ctx.Output.Header("Content-Type", "application/ctet-stream")
	// c.Ctx.Output.Header("Expires", "0")
	// c.Ctx.Output.Header("Cache-Control", "must-revalidate")
	// c.Ctx.Output.Header("Pragma", "public")

	// fileName := url.QueryEscape(fileName)
	// c.Ctx.Output.Header("Content-Type", "application/ctet-stream")
	// c.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s;filename*=utf-8''%s", fileName, fileName))
	// c.Ctx.Output.Header("Content-Description", "File Transfer")
	// c.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	// c.Ctx.Output.Header("Expires", "0")
	// c.Ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	// c.Ctx.Output.Header("Pragma", "no-cache")

	c.Ctx.Output.Download(path, fileName+".csv")
	// http.ServeFile(c.Ctx.Output.Context.ResponseWriter, c.Ctx.Output.Context.Request, path)
}

// GetReportLeaveRequest ...
func (c *LeaveController) GetReportLeaveRequest() {
	var (
		resp    structAPI.RespData
		dbLeave db.LeaveRequest
	)

	var reqDt = structAPI.RequestReport{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
	}

	resGet, errGetAcc := dbLeave.ReportLeaveRequest(&reqDt)
	if errGetAcc != nil {
		resp.Error = errGetAcc.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetReportLeaveRequest", err)
	}
}

// GetReportLeaveRequestTypeLeave ...
func (c *LeaveController) GetReportLeaveRequestTypeLeave() {
	var (
		resp    structAPI.RespData
		dbLeave db.LeaveRequest
	)

	var reqDt = structAPI.RequestReportTypeLeave{
		FromDate:    c.Ctx.Input.Query("fromDate"),
		ToDate:      c.Ctx.Input.Query("toDate"),
		TypeLeaveID: c.Ctx.Input.Query("typeID"),
	}

	resGet, errGetAcc := dbLeave.ReportLeaveRequestTypeLeave(&reqDt)
	if errGetAcc != nil {
		resp.Error = errGetAcc.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetReportLeaveRequest", err)
	}
}
