package model

type User struct {
	ID       string `json:"id"`
	Username string `gorm:"unique" json:"username"`
	Hash     string `json:"-"`
}

type RegisterRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
