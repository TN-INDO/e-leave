package user

import (
	dbInterfaceAdmin "server/models/db/interfaces/admin"
	dbInterfaceDirector "server/models/db/interfaces/director"
	dbInterfaceLeave "server/models/db/interfaces/leave_request"
	dbInterfaceSupervisor "server/models/db/interfaces/supervisor"
	dbInterfaceUser "server/models/db/interfaces/user"

	dbLayerAdmin "server/models/db/pgsql/admin"
	dbLayerDirector "server/models/db/pgsql/director"
	dbLayerLeave "server/models/db/pgsql/leave_request"
	dbLayerSupervisor "server/models/db/pgsql/supervisor"
	dbLayerUser "server/models/db/pgsql/user"
)

// constant var
var (
	DBPostAdmin      dbInterfaceAdmin.IBaseAdmin
	DBPostUser       dbInterfaceUser.IBaseUser
	DBPostLeave      dbInterfaceLeave.IBaseLeaveRequest
	DBPostDirector   dbInterfaceDirector.IBaseDirector
	DBPostSupervisor dbInterfaceSupervisor.IBaseSupervisor
)

func init() {
	DBPostUser = new(dbLayerUser.User)
	DBPostLeave = new(dbLayerLeave.LeaveRequest)
	DBPostAdmin = new(dbLayerAdmin.Admin)
	DBPostDirector = new(dbLayerDirector.Director)
	DBPostSupervisor = new(dbLayerSupervisor.Supervisor)
}
