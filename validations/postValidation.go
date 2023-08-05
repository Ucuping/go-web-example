package validations

type PostValidation struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}
