package helpers

import (
	"bytes"
	"html/template"
	"path/filepath"
	"server/helpers/constant"

	gomail "gopkg.in/gomail.v2"
)

type employeeMail struct {
	LeaveID        string
	EmployeeName   string
	SupervisorName string
}

type supervisorMail struct {
	// LeaveID        string
	EmployeeName   string
	SupervisorName string
}

type supervisorMailReject struct {
	LeaveID        string
	EmployeeName   string
	SupervisorName string
	Reason         string
}

type directorMail struct {
	LeaveID        string
	EmployeeName   string
	SupervisorName string
	DirectorName   string
}
type directorMailAccept struct {
	LeaveID      string
	EmployeeName string
	DirectorName string
}

type directorMailReject struct {
	LeaveID      string
	EmployeeName string
	DirectorName string
	Reason       string
}

type directorMailCancel struct {
	LeaveID      string
	EmployeeName string
	DirectorName string
}

type employeeMailCancel struct {
	LeaveID      string
	EmployeeName string
}

type sendPassword struct {
	EmployeeName string
	Password     string
}

type newUser struct {
	Password string
}

// GoMailEmployee ...
func GoMailEmployee(mailTo string, leaveID string, employeeName string, supervisorName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("employee.html")
	infoHTML := employeeMail{leaveID, employeeName, supervisorName}
	t, errParse = t.ParseFiles(filePrefix + "/employee.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Accepted Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailSupervisor ...
func GoMailSupervisor(mailTo string, employeeName string, supervisorName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("supervisor.html")
	infoHTML := supervisorMail{employeeName, supervisorName}
	t, errParse = t.ParseFiles(filePrefix + "/supervisor.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Request Leave")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailSupervisorReject ...
func GoMailSupervisorReject(mailTo string, leaveID string, employeeName string, supervisorName string, reason string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("supervisor_reject.html")
	infoHTML := supervisorMailReject{leaveID, employeeName, supervisorName, reason}
	t, errParse = t.ParseFiles(filePrefix + "/supervisor_reject.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Reject Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailDirector ...
func GoMailDirector(mailTo string, leaveID string, employeeName string, supervisorName string, DirectorName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("director.html")
	infoHTML := directorMail{leaveID, employeeName, supervisorName, DirectorName}
	t, errParse = t.ParseFiles(filePrefix + "/director.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Request Leave")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailDirectorAccept ...
func GoMailDirectorAccept(mailTo string, leaveID string, employeeName string, directorName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("director_accept.html")
	infoHTML := directorMailAccept{leaveID, employeeName, directorName}
	t, errParse = t.ParseFiles(filePrefix + "/director_accept.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Accept Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)
	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailDirectorReject ...
func GoMailDirectorReject(mailTo string, leaveID string, employeeName string, directorName string, reason string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("director_reject.html")
	infoHTML := directorMailReject{leaveID, employeeName, directorName, reason}
	t, errParse = t.ParseFiles(filePrefix + "/director_reject.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Reject Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailDirectorCancel ...
func GoMailDirectorCancel(mailTo string, leaveID string, employeeName string, directorName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("cancel_request_director.html")
	infoHTML := directorMailCancel{leaveID, employeeName, directorName}
	t, errParse = t.ParseFiles(filePrefix + "/cancel_request_director.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Cancel Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailEmployeeCancel ...
func GoMailEmployeeCancel(mailTo string, leaveID string, employeeName string) {

	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("cancel_request_employee.html")
	infoHTML := employeeMailCancel{leaveID, employeeName}
	t, errParse = t.ParseFiles(filePrefix + "/cancel_request_employee.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Cancel Leave Request")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailForgotPassword ...
func GoMailForgotPassword(mailTo string, employeeName string) {

	var errParse error
	resetPassword := constant.GOPWDRESET
	filePrefix, _ := filepath.Abs("./views")
	t := template.New("forgot_password.html")
	infoHTML := sendPassword{employeeName, resetPassword}
	t, errParse = t.ParseFiles(filePrefix + "/forgot_password.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Forgot Password")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailRegisterPassword ...
func GoMailRegisterPassword(mailTo string, password string) {

	var errParse error
	filePrefix, _ := filepath.Abs("./views")
	t := template.New("register_password.html")
	infoHTML := newUser{password}
	t, errParse = t.ParseFiles(filePrefix + "/register_password.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Register Password")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}

// GoMailDirectorFromSupervisor ...
func GoMailDirectorFromSupervisor(mailTo string, employeeName string, directorName string) {
	var errParse error

	filePrefix, _ := filepath.Abs("./views")
	t := template.New("supervisor_leave.html")
	infoHTML := supervisorMail{employeeName, directorName}
	t, errParse = t.ParseFiles(filePrefix + "/supervisor_leave.html")
	if errParse != nil {
		CheckErr("errParse ", errParse)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, infoHTML); err != nil {
		CheckErr("err ", err)
	}
	mailHTML := tpl.String()

	authEmail := constant.EmailNoRepply
	authPassword := constant.GOPWD
	authHost := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", authEmail)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", "Request Leave")
	m.Embed(filePrefix + "/tnis.png")
	m.SetBody("text/html", mailHTML)

	d := gomail.NewDialer(authHost, port, authEmail, authPassword)

	if err := d.DialAndSend(m); err != nil {
		CheckErr("error email", err)
	}
}
