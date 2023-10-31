package productservice

import (
	"fmt"
	"shopstoretest/param"
)

func (s Service) ShowByCategory(req param.ShowByCategoryRequest) (param.ShowByCategoryResponse, error) {
	exist, eErr := s.Repository.CheckExistCategory(req.CategoryStr)
	if eErr != nil {

		return param.ShowByCategoryResponse{}, fmt.Errorf("cant check existence %w", eErr)
	}

	if !exist {

		return param.ShowByCategoryResponse{}, fmt.Errorf("category undifined!")
	}

	products, gErr := s.Repository.GetProductByCategory(req.CategoryStr)
	if gErr != nil {

		return param.ShowByCategoryResponse{}, fmt.Errorf("cant get products %w", products)
	}

	res := param.ShowByCategoryResponse{Products: products}

	return res, nil
}

/*func (s Service) ShowByCategory(req param.ShowByCategoryRequest) (param.ShowByCategoryResponse, error) {
	exist, eErr := s.Repository.CheckExistCategory(req.CategoryStr)
	if eErr != nil {

		return param.ShowByCategoryResponse{}, fmt.Errorf("cant check existence %w", eErr)
	}

	if !exist {

		return param.ShowByCategoryResponse{}, fmt.Errorf("category !")
	}
}*/
