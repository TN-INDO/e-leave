package leave

import (
	dbInterfaceLeave "server/models/db/interfaces/leave_request"
	dbInterfaceUser "server/models/db/interfaces/user"

	dbLayerLeave "server/models/db/pgsql/leave_request"
	dbLayerUser "server/models/db/pgsql/user"
)

// constant var
var (
	DBUser  dbInterfaceUser.IBaseUser
	DBLeave dbInterfaceLeave.IBaseLeaveRequest
)

func init() {
	DBUser = new(dbLayerUser.User)
	DBLeave = new(dbLayerLeave.LeaveRequest)

}
