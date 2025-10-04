package repo

import (
	"database/sql"
	"fmt"
	"practice/domain"
	"practice/product"

	"github.com/jmoiron/sqlx"
)


type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// Constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
        INSERT INTO products (title, description, price, image_url)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	err := r.db.QueryRow(query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageUrl,
	).Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(pid int) (*domain.Product, error) {
	var p domain.Product
	query := `
        SELECT id, title, description, price, image_url
        FROM products
        WHERE id = $1;
    `
	err := r.db.Get(&p, query, pid)

	if err == sql.ErrNoRows {
		return nil, nil // not found
	} else if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	var p []*domain.Product
	query := `
        SELECT id, title, description, price, image_url
        FROM products
    `
	err := r.db.Select(&p, query)
	fmt.Println("P:",p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
        UPDATE products
        SET title = $1,
            description = $2,
            price = $3,
            image_url = $4,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $5
       
    `

	err := r.db.QueryRow(query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageUrl,
		p.ID,
	).Err()

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}
