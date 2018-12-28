package supervisor

import (
	dbInterfaceLeave "server/models/db/interfaces/leave_request"
	dbInterfaceSupervisor "server/models/db/interfaces/supervisor"
	dbInterfaceUser "server/models/db/interfaces/user"
	dbLayerLeave "server/models/db/pgsql/leave_request"
	dbLayerSupervisor "server/models/db/pgsql/supervisor"
	dbLayerUser "server/models/db/pgsql/user"
)

var (
	DBSupervisor dbInterfaceSupervisor.IBaseSupervisor
	DBUser       dbInterfaceUser.IBaseUser
	DBLeave      dbInterfaceLeave.IBaseLeaveRequest
)

func init() {
	DBSupervisor = new(dbLayerSupervisor.Supervisor)
	DBUser = new(dbLayerUser.User)
	DBLeave = new(dbLayerLeave.LeaveRequest)
}
