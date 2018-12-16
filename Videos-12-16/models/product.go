package models

import (
	"fmt"
)

type Product struct {
	Id          int
	Description string
	Quantity    int
	Price       float64
	Amount      float64
	CreatedAt   string
}

func GetAll() ([]Product, error) {

	con := Connect()

	sql := "select * from product"

	rs, erro := con.Query( sql )

	if erro != nil {

		return nil, erro

	}

	var products []Product

	for rs.Next() {

		var product Product

		erro := rs.Scan(
			&product.Id,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.Amount,
			&product.CreatedAt)

		if erro != nil {

			return nil, erro 
		
		}

		products = append(products, product)

	}

	defer rs.Close()
	defer con.Close()

	return products, nil

}

func NewProduct(product Product) (bool, error) {

	con := Connect()

	sql := "insert into product (description, quantity, price, amount) values (?, ?, ?, ?)"

	stmt, erro := con.Prepare( sql )

	if erro != nil {

		return false, erro

	}

	_, erro = stmt.Exec(product.Description, product.Quantity, product.Price, product.Amount)

	if erro != nil {

		return false, erro

	}

	defer stmt.Close()
	defer con.Close()

	return true, nil

}

func UpdateProduct(product Product) (bool, error) {

	con := Connect()

	sql := "update product set description = ?, quantity = ?, price = ?, amount = ? where id = ?"

	stmt, erro := con.Prepare( sql )

	if erro != nil {

		return false, erro

	}

	_, erro = stmt.Exec(product.Description, product.Quantity, product.Price, product.Amount, product.Id)

	if erro != nil {

		return false, erro

	}

	defer stmt.Close()
	defer con.Close()

	return true, nil

}

func DeleteProduct(field, value string) (int64, error) {

	con := Connect()

	sql := fmt.Sprintf("delete from product where %s = ?", field)

	stmt, erro := con.Prepare( sql )

	if erro != nil {

		return -1, erro

	}

	rs, erro := stmt.Exec( value )

	if erro != nil {

		return -1, erro

	}

	rows, erro := rs.RowsAffected()

	if erro != nil {

		return -1, erro

	}

	defer stmt.Close()
	defer con.Close()

	return rows, nil

}

func Find(id int) (Product, error) {

	con := Connect()

	sql := "select * from product where id = ?"

	rs, erro := con.Query(sql, id)

	if erro != nil {

		return Product{}, nil

	}

	var product Product

	if rs.Next() {

		erro := rs.Scan(&product.Id, &product.Description, &product.Quantity, &product.Price, &product.Amount, &product.CreatedAt)

		if erro != nil {

			return Product{}, nil

		}
	}

	defer con.Close()
	defer rs.Close()

	return product, nil

}