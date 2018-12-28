package admin

import (
	dbInterfacAdmin "server/models/db/interfaces/admin"
	dbInterfaceLeave "server/models/db/interfaces/leave_request"
	dbInterfaceUser "server/models/db/interfaces/user"
	dbLayerAdmin "server/models/db/pgsql/admin"
	dbLayerLeave "server/models/db/pgsql/leave_request"
	dbLayerUser "server/models/db/pgsql/user"
)

var (
	DBAdmin dbInterfacAdmin.IBaseAdmin
	DBUser  dbInterfaceUser.IBaseUser
	DBLeave dbInterfaceLeave.IBaseLeaveRequest
)

func init() {
	DBAdmin = new(dbLayerAdmin.Admin)
	DBUser = new(dbLayerUser.User)
	DBLeave = new(dbLayerLeave.LeaveRequest)
}
