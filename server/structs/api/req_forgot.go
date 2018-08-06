package api

// ReqForgot ...
type ReqForgot struct {
	Email string `json:"email" orm:"column(email)"`
}
