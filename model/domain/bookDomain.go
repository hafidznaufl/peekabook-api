package domain

type Book struct {
	ID        uint
	Title     string
	AuthorID  uint
	Page      int
	Years     int
	Publisher string
	Type      string
	Quantity  int
	Status    string
}