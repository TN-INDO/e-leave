package logic

// GetDirectorID ...
type GetDirectorID struct {
	DirectorID int64 `json:"director_id" orm:"column(director_id)"`
}

// GetDirector ...
type GetDirector struct {
	Name  string `json:"name" orm:"column(name)"`
	Email string `json:"email" orm:"column(email)"`
}
