package response

import "time"

type (
	LoginResponse struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name" example:"user name"`
		Email     string    `json:"email" example:"email@email.com"`
		UserLevel uint      `json:"user_level"`
		CreatedAt time.Time `json:"createdAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
		UpdatedAt time.Time `json:"updatedAt,omitempty" example:"2023-02-11T15:01:00+00:00"`
		Token     string    `json:"token,omitempty"`
		Expires   time.Time `json:"expires,omitempty"`
	}
)
