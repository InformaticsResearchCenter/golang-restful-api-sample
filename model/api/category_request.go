package api

type CategoryCreateRequest struct {
	Name string `validate:"required,max=255,min=100"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=255,min=1"`
}
