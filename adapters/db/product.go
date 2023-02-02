package db

import (
	"database/sql"
	"fmt"

	"github.com/infezek/product_ports_adapters_golang/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (p *ProductDb) Migration() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		status TEXT NOT NULL
	);
	`
	_, err := p.db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stml, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stml.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows string
	p.db.QueryRow(`select id from products where id=?`, product.GetID()).Scan(&rows)
	fmt.Println(rows)
	fmt.Println(product.GetID())
	if rows == "" {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	query := `insert into products(id, name, price,status) values (?,?,?,?)`
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	query := `update products set name = ?, price = ?, status = ? where id = ?`
	_, err := p.db.Exec(query,
		product.GetName(),
		product.GetPrice(),
		product.GetPrice(),
		product.GetID(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
