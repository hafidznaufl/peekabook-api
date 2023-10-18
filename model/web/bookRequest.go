package web

type BookCreateRequest struct {
	Title     string `json:"title" validate:"required"`
	AuthorID  uint   `json:"authorId" validate:"required"`
	Page      int    `json:"page" validate:"required"`
	Years     int    `json:"years" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

type BookUpdateRequest struct {
	Title     string `json:"title" validate:"required"`
	AuthorID  uint   `json:"authorId" validate:"required"`
	Page      int    `json:"page" validate:"required"`
	Years     int    `json:"years" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
