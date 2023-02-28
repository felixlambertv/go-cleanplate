package request

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required" example:"email@email.com"`
		Password string `json:"password" binding:"required" example:"password123"`
	}
)
