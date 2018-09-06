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
	logicLeave "server/models/logic/leave"
	logicUser "server/models/logic/user"
	structAPI "server/structs/api"

	"github.com/astaxie/beego"
)

//LeaveController ...
type LeaveController struct {
	beego.Controller
}

// PostLeaveRequestEmployee ...
func (c *LeaveController) PostLeaveRequestEmployee() {
	var (
		req  structAPI.ReqLeave
		resp structAPI.RespData
	)

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @PostLeaveRequestEmployee - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-LEAVE-REQUEST=======>", string(body))

	errMarshal := json.Unmarshal(body, &req)
	if errMarshal != nil {
		helpers.CheckErr("Failed unmarshall req body @PostLeaveRequestEmployee - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	totalDay := float64(helpers.GetTotalDay(req.DateFrom, req.DateTo))
	reqHalfDay := float64(len(req.HalfDates))
	valueHalfDay := float64(0.5)
	result := helpers.Multiply(totalDay, reqHalfDay, valueHalfDay)

	resGet, errGet := logicUser.GetUserLeaveRemaining(req.TypeLeaveID, employeeNumber)
	helpers.CheckErr("Error get leave balance @PostLeaveRequestEmployee - controller", errGet)

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
		beego.Warning("Error empty field type leave @PostLeaveRequestEmployee - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("error empty field").Error()

	} else if req.DateFrom == "" || req.DateTo == "" || req.BackOn == "" || req.ContactAddress == "" || req.ContactNumber == "" {
		beego.Warning("Error empty field @PostLeaveRequestEmployee - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("Error empty field").Error()

	} else if result > float64(resGet.LeaveRemaining) {
		beego.Warning("Error leave balance @PostLeaveRequestEmployee - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("Your total leave is " + strTotal + " day and your " + resGet.TypeName + " balance is " + strBalance + " day left").Error()

	} else {
		errAddLeave := logicLeave.CreateLeaveRequestEmployee(
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
		helpers.CheckErr("Failed giving output @PostLeaveRequestEmployee - controller", err)
	}
}

// PostLeaveRequestSupervisor ...
func (c *LeaveController) PostLeaveRequestSupervisor() {
	var (
		req  structAPI.ReqLeave
		resp structAPI.RespData
	)

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @PostLeaveRequestSupervisor - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-LEAVE-REQUEST=======>", string(body))

	errMarshal := json.Unmarshal(body, &req)
	if errMarshal != nil {
		helpers.CheckErr("Failed unmarshall req body @PostLeaveRequestSupervisor - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	totalDay := float64(helpers.GetTotalDay(req.DateFrom, req.DateTo))
	reqHalfDay := float64(len(req.HalfDates))
	valueHalfDay := float64(0.5)
	result := helpers.Multiply(totalDay, reqHalfDay, valueHalfDay)

	resGet, errGet := logicUser.GetUserLeaveRemaining(req.TypeLeaveID, employeeNumber)
	helpers.CheckErr("Error get leave balance @PostLeaveRequestSupervisor - controller", errGet)

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
		beego.Warning("Error empty field type leave @PostLeaveRequestSupervisor - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("Error empty field").Error()

	} else if req.DateFrom == "" || req.DateTo == "" || req.BackOn == "" || req.ContactAddress == "" || req.ContactNumber == "" {
		beego.Warning("Error empty field @PostLeaveRequestSupervisor - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("Error empty field").Error()

	} else if result > float64(resGet.LeaveRemaining) {
		beego.Warning("Error leave balance @PostLeaveRequestSupervisor - controller")
		c.Ctx.Output.SetStatus(400)
		resp.Error = errors.New("Your total leave is " + strTotal + " day and your " + resGet.TypeName + " balance is " + strBalance + " day left").Error()

	} else {
		errAddLeave := logicLeave.CreateLeaveRequestSupervisor(
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
		helpers.CheckErr("Failed giving output @PostLeaveRequestSupervisor - controller", err)
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
		helpers.CheckErr("Failed unmarshall req body @UpdateRequest - controller", err)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @UpdateRequest - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
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
		helpers.CheckErr("Failed giving output @UpdateRequest - controller", err)
	}
}

// DeleteRequest ...
func (c *LeaveController) DeleteRequest() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @DeleteRequest - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	if err := logicLeave.DeleteRequest(id); err == nil {
		resp.Body = "Delete request success"
	} else {
		resp.Error = err.Error()
		c.Ctx.Output.SetStatus(400)
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @DeleteRequest - controller", err)
	}
}

// GetDownloadReportCSV ...
func (c *LeaveController) GetDownloadReportCSV() {
	var reqDt = structAPI.RequestReport{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
	}

	dt := time.Now()
	fileName := "report_leave_request_" + dt.Format("20060102")
	path := constant.GOPATH + "/src/" + constant.GOAPP + "/storages/" + fileName + ".csv"

	errGet := logicLeave.DownloadReportCSV(&reqDt, path)
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
	var resp structAPI.RespData

	var reqDt = structAPI.RequestReport{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
	}

	resGet, errGetAcc := logicLeave.ReportLeaveRequest(&reqDt)
	if errGetAcc != nil {
		resp.Error = errGetAcc.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetReportLeaveRequest - controller", err)
	}
}

// GetReportLeaveRequestTypeLeave ...
func (c *LeaveController) GetReportLeaveRequestTypeLeave() {
	var resp structAPI.RespData

	var reqDt = structAPI.RequestReportTypeLeave{
		FromDate:    c.Ctx.Input.Query("fromDate"),
		ToDate:      c.Ctx.Input.Query("toDate"),
		TypeLeaveID: c.Ctx.Input.Query("typeID"),
	}

	resGet, errGetAcc := logicLeave.ReportLeaveRequestTypeLeave(&reqDt)
	if errGetAcc != nil {
		resp.Error = errGetAcc.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetReportLeaveRequest - controller", err)
	}
}
