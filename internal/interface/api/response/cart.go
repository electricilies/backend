package response

type CartItem struct {
	ID       int     `json:"id"`
	Product  Product `json:"product_id"`
	Quantity int     `json:"quantity"`
}

type Cart struct {
	ID    int        `json:"id"`
	Items []CartItem `json:"items"`
}
