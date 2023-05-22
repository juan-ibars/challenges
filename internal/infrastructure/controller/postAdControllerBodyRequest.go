package controller

type PostAdControllerBodyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
}
