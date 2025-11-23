package http

import (
	"net/http"

	"backend/internal/application"
	_ "backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GinCartHandler struct {
	cartApp               application.Cart
	ErrRequiredCartID     string
	ErrInvalidCartID      string
	ErrRequiredCartItemID string
	ErrInvalidCartItemID  string
}

var _ CartHandler = &GinCartHandler{}

func ProvideCartHandler(cartApp application.Cart) *GinCartHandler {
	return &GinCartHandler{
		cartApp:               cartApp,
		ErrRequiredCartID:     "cart_id is required",
		ErrInvalidCartID:      "invalid cart_id",
		ErrRequiredCartItemID: "cart_item_id is required",
		ErrInvalidCartItemID:  "invalid cart_item_id",
	}
}

// GetCart godoc
//
//	@Summary		Get cart
//	@Description	Get cart by user ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	false	"Page for pagination"	default(1)
//	@Param			limit	query		int	false	"Limit"					default(20)
//	@Success		200		{object}	domain.Cart
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id} [get]
func (h *GinCartHandler) Get(ctx *gin.Context) {
	cartIDString := ctx.Param("cart_id")
	if cartIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	cartID, err := uuid.Parse(cartIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}
	cart, err := h.cartApp.Get(ctx, application.GetCartParam{
		CartID: cartID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// AddCartItem godoc
//
//	@Summary		Add item to cart
//	@Description	Add a product item to the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			item	body		application.CreateCartItemData	true	"Cart item request"
//	@Success		201		{object}	domain.CartItem
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id}/item  [post]
func (h *GinCartHandler) CreateItem(ctx *gin.Context) {
	cartIDString := ctx.Param("cart_id")
	if cartIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	cartID, err := uuid.Parse(cartIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	var data application.CreateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	// TODO: Get userID from auth context
	userID := uuid.New() // Placeholder

	cartItem, err := h.cartApp.CreateItem(ctx, application.CreateCartItemParam{
		UserID: userID,
		CartID: cartID,
		Data:   data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, cartItem)
}

// UpdateCartItem godoc
//
//	@Summary		Update cart item
//	@Description	Update quantity of a cart item
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart_item_id	path		int								true	"Cart Item ID"
//	@Param			item			body		application.UpdateCartItemData	true	"Update cart item request"
//	@Success		200				{object}	domain.CartItem
//	@Failure		400				{object}	Error
//	@Failure		404				{object}	Error
//	@Failure		500				{object}	Error
//	@Router			/carts/{cart_id}/item [patch]
func (h *GinCartHandler) UpdateItem(ctx *gin.Context) {
	cartIDString := ctx.Param("cart_id")
	if cartIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	cartID, err := uuid.Parse(cartIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	itemIDString := ctx.Param("cart_item_id")
	if itemIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	itemID, err := uuid.Parse(itemIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	var data application.UpdateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	// TODO: Get userID from auth context
	userID := uuid.New() // Placeholder

	cartItem, err := h.cartApp.UpdateItem(ctx, application.UpdateCartItemParam{
		UserID: userID,
		CartID: cartID,
		ItemID: itemID,
		Data:   data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cartItem)
}

// RemoveCartItem godoc
//
//	@Summary		Remove cart item
//	@Description	Remove an item from the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			item_id	path	int	true	"Cart Item ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/carts/{cart_id}/item [delete]
func (h *GinCartHandler) RemoveItem(ctx *gin.Context) {
	cartIDString := ctx.Param("cart_id")
	if cartIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	cartID, err := uuid.Parse(cartIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	itemIDString := ctx.Param("item_id")
	if itemIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	itemID, err := uuid.Parse(itemIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	// TODO: Get userID from auth context
	userID := uuid.New() // Placeholder

	err = h.cartApp.DeleteItem(ctx, application.DeleteCartItemParam{
		UserID: userID,
		CartID: cartID,
		ItemID: itemID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
