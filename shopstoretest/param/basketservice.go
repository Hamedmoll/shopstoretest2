package param

import "shopstoretest/entity"

type AddToBasketRequest struct {
	ItemID uint `json:"item_id"`
	UserID uint `json:"user_id"`
}

type AddToBasketResponse struct {
	Basket entity.Basket `json:"basket"`
}

type ShowBasketRequest struct {
	ID uint
}

type ShowBasketResponse struct {
	Baskets []entity.Basket
}
