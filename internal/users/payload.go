package users

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Age      int    `json:"age" validate:"required" example:"25"`
	Password string `json:"password" validate:"required" example:"secretpassword"`
}

type UserCreateResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john.doe@example.com"`
	Age   int    `json:"age" example:"25"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Age      int    `json:"age" validate:"required" example:"25"`
	Password string `json:"password" example:"secretpassword"`
}

type UserUpdateResponse struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john.doe@example.com"`
	Age   int    `json:"age" example:"25"`
}
