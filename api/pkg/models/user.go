package models

type User struct {
	BaseModel

	Email        string `json:"email" gorm:"unique"`
	Username     string `json:"username" gorm:"unique;size=32"`
	PasswordHash string `json:"password_hash"`
}

func NewUser(email, username, passwordHash string) User {
	return User{
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}
}
