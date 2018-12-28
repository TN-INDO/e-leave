package director

import (
	dbInterfaceDirector "server/models/db/interfaces/director"
	dbInterfaceLeave "server/models/db/interfaces/leave_request"
	dbInterfaceUser "server/models/db/interfaces/user"
	dbLayerDirector "server/models/db/pgsql/director"
	dbLayerLeave "server/models/db/pgsql/leave_request"
	dbLayerUser "server/models/db/pgsql/user"
)

// constant var
var (
	DBDirector dbInterfaceDirector.IBaseDirector
	DBUser     dbInterfaceUser.IBaseUser
	DBLeave    dbInterfaceLeave.IBaseLeaveRequest
)

func init() {
	DBDirector = new(dbLayerDirector.Director)
	DBUser = new(dbLayerUser.User)
	DBLeave = new(dbLayerLeave.LeaveRequest)
}
