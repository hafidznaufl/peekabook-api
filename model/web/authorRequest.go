package web

type AuthorCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type AuthorUpdateRequest struct {
	Name     string `json:"name" validate:"required"`
}