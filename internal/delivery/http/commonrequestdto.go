package http

type PaginationRequestDto struct {
	Page  int `binding:"gte=1,lte=50"`
	Limit int `binding:"gte=1,lte=100"`
}
