package api

// ReqLogin ...
type ReqLogin struct {
	Email    string `json:"email" orm:"column(email)"`
	Password string `json:"password" orm:"column(password)"`
}
