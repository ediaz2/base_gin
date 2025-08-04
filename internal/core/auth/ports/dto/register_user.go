package dto

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
