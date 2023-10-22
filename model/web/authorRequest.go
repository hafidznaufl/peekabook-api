package web

type AuthorCreateRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type AuthorUpdateRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
