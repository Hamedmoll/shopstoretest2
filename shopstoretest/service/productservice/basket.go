package productservice

import (
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (s Service) AddBasket(req param.AddToBasketRequest) (param.AddToBasketResponse, error) {
	user, guErr := s.Repository.GetUserByID(req.UserID)
	if guErr != nil {

		return param.AddToBasketResponse{}, guErr
	}

	product, gpErr := s.Repository.GetProductByID(req.ItemID)
	if gpErr != nil {

		return param.AddToBasketResponse{}, gpErr
	}

	newBasket := entity.Basket{
		ID:        0,
		UserID:    user.ID,
		ProductID: product.ID,
		Price:     product.Price,
	}

	createdBasket, aErr := s.Repository.AddBasket(newBasket)

	if aErr != nil {

		return param.AddToBasketResponse{}, aErr
	}

	res := param.AddToBasketResponse{Basket: createdBasket}

	return res, nil
}

func (s Service) ShowBasket(req param.ShowBasketRequest) (param.ShowBasketResponse, error) {
	baskets, gErr := s.Repository.GetBasketsByID(req.ID)
	if gErr != nil {

		return param.ShowBasketResponse{}, gErr
	}

	res := param.ShowBasketResponse{
		Baskets: baskets,
	}

	return res, nil
}
