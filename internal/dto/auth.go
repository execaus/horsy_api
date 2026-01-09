package dto

type SignUpRequest struct {
	Email string `json:"email" validate:"required,email,min=4,max=255"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,min=4,max=255"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
