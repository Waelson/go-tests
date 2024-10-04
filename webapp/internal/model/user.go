package model

// User model info
// @Description User information
// @Description with user id, email, password and username
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
