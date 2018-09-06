package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",

		// ========================= user ========================= //
		// login
		beego.NSRouter("/login",
			&controllers.UserController{},
			"post:Login",
		),
		// password reset
		beego.NSRouter("/user/password-reset",
			&controllers.UserController{},
			"put:PasswordReset",
		),

		// upadate new password
		beego.NSRouter("/user/update/:id:int ",
			&controllers.UserController{},
			"put:UpdateNewPassword",
		),

		// get employee sumarry
		beego.NSRouter("/user/summary/:id:int",
			&controllers.UserController{},
			"get:GetUserSummary",
		),

		// get employee type leave
		beego.NSRouter("/user/type-leave/:id:int",
			&controllers.UserController{},
			"get:GetUserTypeLeave",
		),

		// get type leave
		beego.NSRouter("/user/type-leave",
			&controllers.UserController{},
			"get:GetTypeLeave",
		),

		// get supervisor
		beego.NSRouter("/user/supervisor",
			&controllers.UserController{},
			"get:GetSupervisors",
		),

		// ========================= leave-request ========================= //
		// create leave request for employee, update leave request, delete leave request
		beego.NSRouter("/employee/leave/:id:int ",
			&controllers.LeaveController{},
			"post:PostLeaveRequestEmployee;put:UpdateRequest;delete:DeleteRequest",
		),
		// create leave request for supervisor
		beego.NSRouter("/supervisor/leave/:id:int ",
			&controllers.LeaveController{},
			"post:PostLeaveRequestSupervisor",
		),

		// download report leave request csv
		beego.NSRouter("/leave/report",
			&controllers.LeaveController{},
			"get:GetDownloadReportCSV",
		),
		// get report leave request
		beego.NSRouter("/leave/reports",
			&controllers.LeaveController{},
			"get:GetReportLeaveRequest",
		),
		// get report leave request
		beego.NSRouter("/leave/report/type",
			&controllers.LeaveController{},
			"get:GetReportLeaveRequestTypeLeave",
		),

		// ========================= admin ========================= //
		// register user
		beego.NSRouter("/admin/user/register",
			&controllers.AdminController{},
			"post:CreateUser",
		),
		// get all user
		beego.NSRouter("/admin/user",
			&controllers.AdminController{},
			"get:GetUsers;post:CreateUser",
		),
		// get one user, update one user, delete one user
		beego.NSRouter("/admin/user/:id:int",
			&controllers.AdminController{},
			"get:GetUser;put:UpdateUser;delete:DeleteUser",
		),
		// delete user
		beego.NSRouter("/admin/user/delete/:id:int",
			&controllers.AdminController{},
			"delete:DeleteUser",
		),
		// get leave request pending
		beego.NSRouter("/admin/leave/pending",
			&controllers.AdminController{},
			"get:GetRequestPending",
		),
		// get leave request approve
		beego.NSRouter("/admin/leave/accept",
			&controllers.AdminController{},
			"get:GetRequestAccept",
		),
		// get leave request reject
		beego.NSRouter("/admin/leave/reject",
			&controllers.AdminController{},
			"get:GetRequestReject",
		),

		// canceled leave request by admin
		beego.NSRouter("/admin/leave/cancel/:id:int/:enumber:int",
			&controllers.AdminController{},
			"put:CancelRequestLeave",
		),
		// reset leave balance by admin
		beego.NSRouter("/admin/leave/reset_balance",
			&controllers.AdminController{},
			"put:ResetLeaveBalance",
		),

		// ========================= director ========================= //
		// accept status by director
		beego.NSRouter("/director/accept/:id:int/:enumber:int",
			&controllers.DirectorController{},
			"put:AcceptStatusByDirector",
		),
		// reject status by director
		beego.NSRouter("/director/reject/:id:int/:enumber:int",
			&controllers.DirectorController{},
			"put:RejectStatusByDirector",
		),
		// get status pending in director
		beego.NSRouter("/director/pending/",
			&controllers.DirectorController{},
			"get:GetDirectorPendingLeave",
		),
		// get status accept in director
		beego.NSRouter("/director/accept/",
			&controllers.DirectorController{},
			"get:GetDirectorAcceptLeave",
		),
		// get status reject in director
		beego.NSRouter("/director/reject/",
			&controllers.DirectorController{},
			"get:GetDirectorRejectLeave",
		),

		// ========================= supervisor ========================= //
		// accept status by supervisor
		beego.NSRouter("/supervisor/accept/:id:int/:enumber:int",
			&controllers.SupervisorController{},
			"put:AcceptStatusBySupervisor",
		),
		// reject status by supervisor
		beego.NSRouter("/supervisor/reject/:id:int/:enumber:int",
			&controllers.SupervisorController{},
			"put:RejectStatusBySv",
		),
		// get status pending in supervisor
		beego.NSRouter("/supervisor/pending/:id:int ",
			&controllers.SupervisorController{},
			"get:GetPendingLeave",
		),
		// get status accept in supervisor
		beego.NSRouter("/supervisor/accept/:id:int ",
			&controllers.SupervisorController{},
			"get:GetAcceptLeave",
		),
		// get status reject in supervisor
		beego.NSRouter("/supervisor/reject/:id:int ",
			&controllers.SupervisorController{},
			"get:GetRejectLeave",
		),

		// ========================= employee ========================= //
		// get all pending request
		beego.NSRouter("/employee/pending/:id:int ",
			&controllers.EmployeeController{},
			"get:GetRequestPending",
		),
		// get all accept request
		beego.NSRouter("/employee/accept/:id:int ",
			&controllers.EmployeeController{},
			"get:GetRequestAccept",
		),
		// get all reject request in employee
		beego.NSRouter("/employee/reject/:id:int ",
			&controllers.EmployeeController{},
			"get:GetRequestReject",
		),
	)
	beego.AddNamespace(ns)
}
