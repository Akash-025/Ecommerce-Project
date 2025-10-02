package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	ImageUrl    string `json:"image_url" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(pid int) (*Product, error)
	List() ([]*Product, error)
	Update(product Product) (*Product, error)
	Delete(productId int) error
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(pid int) (*Product, error) {
	var p Product
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

func (r *productRepo) List() ([]*Product, error) {
	var p []*Product
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

func (r *productRepo) Update(p Product) (*Product, error) {
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
