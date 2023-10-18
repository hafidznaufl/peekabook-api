package web

type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	AuthorID  uint   `json:"authorId"`
	Page      int    `json:"page"`
	Years     int    `json:"years"`
	Publisher string `json:"publisher"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}
