package mysql

import (
	"fmt"
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (mysql MySQLDB) AddProduct(product param.AddProductRequest) (entity.Product, error) {
	existCategory, eErr := mysql.CheckExistCategory(product.Category)
	if eErr != nil {

		return entity.Product{}, fmt.Errorf("unexpected error %w", eErr)
	}

	if !existCategory {
		newCategory := entity.Category{Name: product.Category}
		_, aErr := mysql.AddCategory(newCategory)
		if aErr != nil {

			return entity.Product{}, fmt.Errorf("unexpected error %w", aErr)
		}
	}

	category, gErr := mysql.GetCategoryByName(product.Category)
	if gErr != nil {
		fmt.Println()
		return entity.Product{}, fmt.Errorf("unexpected error %w", gErr)
	}

	res, exErr := mysql.DB.Exec("insert into products(name, price, description, category_id, count) values(?, ?, ?, ?, ?)",
		product.Name, product.Price, product.Description, category.ID, product.Count)
	if exErr != nil {

		return entity.Product{}, exErr
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.Product{}, fmt.Errorf("cant get last id after save new product in table %w", iErr)
	}

	createdProduct := entity.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  category.ID,
		Count:       product.Count,
	}

	createdProduct.ID = uint(id)

	return createdProduct, nil
}

func (mysql MySQLDB) GetProductByCategory(name string) ([]param.ProductInfo, error) {
	category, gErr := mysql.GetCategoryByName(name)
	if gErr != nil {

		return nil, fmt.Errorf("cant get category %w", gErr)

	}

	rows, qErr := mysql.DB.Query("select name, count, price, Description from products where category_id = ?", category.ID)
	if qErr != nil {

		return nil, fmt.Errorf("unexpected error %w", qErr)
	}

	var tmpProduct param.ProductInfo
	var products []param.ProductInfo

	for rows.Next() {
		sErr := rows.Scan(&tmpProduct.Name, &tmpProduct.Count, &tmpProduct.Price, &tmpProduct.Description)
		if sErr != nil {

			return nil, fmt.Errorf("cant scan %w", sErr)
		}

		tmpProduct.Category = name
		products = append(products, tmpProduct)
	}

	return products, nil
}

func (mysql MySQLDB) GetProductByID(id uint) (entity.Product, error) {
	row := mysql.DB.QueryRow("select * from products where id = ?", id)
	product := entity.Product{}
	var createdAt []uint8

	err := row.Scan(&product.ID, &product.Name, &product.Count, &product.Price,
		&product.Description, &product.CategoryID, &createdAt)
	if err != nil {

		return entity.Product{}, err
	}

	return product, nil
}
