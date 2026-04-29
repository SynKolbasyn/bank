package model

type SignRequest struct {
	Email    string `json:"email" validate:"required,email,max=256"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
