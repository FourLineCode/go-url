package models

type User struct {
	BaseModel

	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func NewUser(email, username, passwordHash string) User {
	return User{
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}
}
