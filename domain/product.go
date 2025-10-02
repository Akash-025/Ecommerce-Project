package domain

type Product struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	ImageUrl    string `json:"image_url" db:"image_url"`
}