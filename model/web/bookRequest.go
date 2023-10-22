package web

type BookCreateRequest struct {
	Title     string `json:"title" form:"title" validate:"required"`
	AuthorID  uint   `json:"authorId" form:"authorId" validate:"required"`
	Page      int    `json:"page" form:"page" validate:"required"`
	Years     int    `json:"years" form:"years" validate:"required"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
	Type      string `json:"type" form:"type" validate:"required"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
}

type BookUpdateRequest struct {
	Title     string `json:"title" form:"title" validate:"required"`
	AuthorID  uint   `json:"authorId" form:"authorId" validate:"required"`
	Page      int    `json:"page" form:"page" validate:"required"`
	Years     int    `json:"years" form:"years" validate:"required"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
	Type      string `json:"type" form:"type" validate:"required"`
	Quantity  int    `json:"quantity" form:"quantity" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
}
