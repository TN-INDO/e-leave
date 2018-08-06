package db

// TypeLeave ...
type TypeLeave struct {
	ID       int64  `json:"id" orm:"column(id);pk"`
	TypeName string `json:"type_name" orm:"column(type_name)"`
}

// TableName ...
func (u *TypeLeave) TableName() string {
	return "type_leave"
}
