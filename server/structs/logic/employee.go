package logic

// GetEmployee ...
type GetEmployee struct {
	Name  string `json:"name" orm:"column(name)"`
	Email string `json:"email" orm:"column(email)"`
}

// NewPassword ...
type NewPassword struct {
	OldPassword     string `json:"old_password" orm:"column(old_password)"`
	NewPassword     string `json:"new_password" orm:"column(new_password)"`
	ConfirmPassword string `json:"confirm_password" orm:"column(confirm_password)"`
}

// PasswordReset ...
type PasswordReset struct {
	Email string `json:"email" orm:"column(email)"`
}
