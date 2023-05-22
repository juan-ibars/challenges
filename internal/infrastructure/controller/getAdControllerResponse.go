package controller

type GetAdControllerResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Date        string `json:"date"`
}
