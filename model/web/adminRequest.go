package web

type AdminCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

type AdminUpdateRequest struct {
	Name     string `json:"name" validate:"min=1,max=255"`
	Email    string `json:"email" validate:"email,min=1,max=255"`
	Password string `json:"password" validate:"min=8,max=255"`
}