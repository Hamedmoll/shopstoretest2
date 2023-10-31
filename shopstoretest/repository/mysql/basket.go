package mysql

import (
	"fmt"
	"shopstoretest/entity"
)

func (mysql MySQLDB) AddBasket(basket entity.Basket) (entity.Basket, error) {
	res, exErr := mysql.DB.Exec("insert into basket(user_id, product_id, price) values(?, ?, ?)",
		basket.UserID, basket.ProductID, basket.ProductID)

	if exErr != nil {

		return entity.Basket{}, exErr
	}

	id, lErr := res.LastInsertId()
	if lErr != nil {

		return entity.Basket{}, lErr
	}

	basket.ID = uint(id)

	return basket, nil
}

func (mysql MySQLDB) GetBasketsByID(id uint) ([]entity.Basket, error) {
	rows, qErr := mysql.DB.Query("select product_id, user_id, price from basket where id = ?", id)
	if qErr != nil {

		return nil, fmt.Errorf("unexpected error %w", qErr)
	}

	var baskets = make([]entity.Basket, 0)
	var tmpBasket entity.Basket

	for rows.Next() {
		sErr := rows.Scan(&tmpBasket.ProductID, &tmpBasket.UserID, &tmpBasket.Price)
		if sErr != nil {

			return nil, fmt.Errorf("cant scan %w", sErr)
		}

		tmpBasket.ID = id
		baskets = append(baskets, tmpBasket)
	}

	return baskets, nil
}
