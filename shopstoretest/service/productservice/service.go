package productservice

import (
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"shopstoretest/param"
	"shopstoretest/repository/mysql"
	"shopstoretest/service/categoryservice"
)

type Service struct {
	Repository      Repository
	CategoryService CategoryService
}

type CategoryService interface {
	AddCategory(req param.AddCategoryRequest) (param.AddCategoryResponse, error)
}

type Repository interface {
	AddBasket(basket entity.Basket) (entity.Basket, error)
	CheckExistCategory(name string) (bool, error)
	AddProduct(product param.AddProductRequest) (entity.Product, error)
	GetCategoryByName(name string) (entity.Category, error)
	GetProductByCategory(name string) ([]param.ProductInfo, error)
	GetUserByID(id uint) (entity.User, error)
	GetProductByID(id uint) (entity.Product, error)
	GetBasketsByID(id uint) ([]entity.Basket, error)
}

func New(cfg cfg.Cfg) Service {
	myRepo := mysql.New(cfg.DataBaseCfg)
	myCategory := categoryservice.New(cfg)
	newService := Service{
		Repository:      myRepo,
		CategoryService: myCategory,
	}
	return newService
}
