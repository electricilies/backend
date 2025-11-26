package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandlerImpl struct {
	cartApp               CartApplication
	ErrRequiredCartID     string
	ErrInvalidCartID      string
	ErrRequiredCartItemID string
	ErrInvalidCartItemID  string
	ErrInvalidUserID      string
}

var _ CartHandler = &CartHandlerImpl{}

func ProvideCartHandler(cartApp CartApplication) *CartHandlerImpl {
	return &CartHandlerImpl{
		cartApp:               cartApp,
		ErrRequiredCartID:     "cart_id is required",
		ErrInvalidCartID:      "invalid cart_id",
		ErrRequiredCartItemID: "cart_item_id is required",
		ErrInvalidCartItemID:  "invalid cart_item_id",
		ErrInvalidUserID:      "invalid user_id",
	}
}

// GetCart godoc
//
//	@Summary		Get cart
//	@Description	Get cart by ID
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart_id	path		string	true	"Cart ID"	format(uuid)
//	@Success		200		{object}	CartResponseDto
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) Get(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "cart_id")
	if *cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}
	cart, err := h.cartApp.Get(ctx, GetCartRequestDto{
		CartID: *cartID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// GetCartByUser godoc
//
//	@Summary		Get cart by user ID
//	@Description	Get cart for the authenticated user
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		string	true	"User ID"	format(uuid)
//	@Success		200		{object}	CartResponseDto
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/users/{user_id} [get]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) GetByUser(ctx *gin.Context) {
	userID, ok := pathToUUID(ctx, "user_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}
	cart, err := h.cartApp.GetByUser(ctx, GetCartByUserRequestDto{
		UserID: *userID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// CreateCart godoc
//
//	@Summary		Create cart
//	@Description	Create a new cart for the user
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart	body		CreateCartData	true	"Create cart request"
//	@Success		201		{object}	CartResponseDto
//	@Failure		500		{object}	Error
//	@Router			/carts  [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) Create(ctx *gin.Context) {
	var data CreateCartData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	cart, err := h.cartApp.Create(ctx, CreateCartRequestDto{
		Data: data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, cart)
}

// AddCartItem godoc
//
//	@Summary		Add item to cart
//	@Description	Add a product item to the cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			cart_id	path		string				true	"Cart ID"	format(uuid)
//	@Param			item	body		CreateCartItemData	true	"Cart item request"
//	@Success		201		{object}	CartItemResponseDto
//	@Failure		400		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id}/item  [post]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) CreateItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "cart_id")
	if *cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	var data CreateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	userID, ok := ctxValueToUUID(ctx, "userID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}

	cartItem, err := h.cartApp.CreateItem(ctx, CreateCartItemRequestDto{
		UserID: *userID,
		CartID: *cartID,
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
//	@Param			cart_id	path		string				true	"Cart ID"		format(uuid)
//	@Param			item_id	path		string				true	"Cart Item ID"	format(uuid)
//	@Param			item	body		UpdateCartItemData	true	"Update cart item request"
//	@Success		200		{object}	CartItemResponseDto
//	@Failure		400		{object}	Error
//	@Failure		404		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/carts/{cart_id}/item/{item_id} [patch]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) UpdateItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "cart_id")
	if *cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	itemID, ok := pathToUUID(ctx, "cart_item_id")
	if *itemID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	var data UpdateCartItemData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	userID, ok := ctxValueToUUID(ctx, "userID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}

	cartItem, err := h.cartApp.UpdateItem(ctx, UpdateCartItemRequestDto{
		UserID: *userID,
		CartID: *cartID,
		ItemID: *itemID,
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
//	@Param			cart_id	path	string	true	"Cart ID"		format(uuid)
//	@Param			item_id	path	string	true	"Cart Item ID"	format(uuid)
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/carts/{cart_id}/item/{item_id} [delete]
//	@Security		OAuth2AccessCode
//	@Security		OAuth2Password
func (h *CartHandlerImpl) RemoveItem(ctx *gin.Context) {
	cartID, ok := pathToUUID(ctx, "cart_id")
	if *cartID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartID))
		return
	}

	itemID, ok := pathToUUID(ctx, "item_id")
	if *itemID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredCartItemID))
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidCartItemID))
		return
	}

	userID, ok := ctxValueToUUID(ctx, "userID")
	if !ok {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidUserID))
		return
	}

	err := h.cartApp.DeleteItem(ctx, DeleteCartItemRequestDto{
		UserID: *userID,
		CartID: *cartID,
		ItemID: *itemID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
