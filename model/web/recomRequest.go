package web

type RecommendationRequest struct {
	Genre   string `json:"genre"`
	Author  string `json:"author"`
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}