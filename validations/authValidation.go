package validations

type LoginValidation struct {
	Username string `json:"username" form:"username" tex:"username" plain:"username" validate:"required"`
	Password string `json:"password" form:"password" tex:"username" plain:"username" validate:"required"`
}

type RegisterValidation struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email,required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
