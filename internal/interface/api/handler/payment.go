package handler

type Payment interface{}

type paymentHandler struct{}

func NewPayment() Payment { return &paymentHandler{} }
