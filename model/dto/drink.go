package dto

type DrinkRes struct {
	ID    int    `json:"id"`
	Who   string `json:"who"`
	Drink string `json:"drink"`
	Sugar string `json:"sugar"`
	Ice   string `json:"ice"`
}
