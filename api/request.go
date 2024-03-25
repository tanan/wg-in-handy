package api

type CreateUserRequest struct {
	Email string `json:"email" binding:"required"`
}
