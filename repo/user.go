package repo

import "github.com/jmoiron/sqlx"

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {

	query := `INSERT INTO users (first_name, last_name, email, password, is_shop_owner) 
			 VALUES (
			 :first_name, 
			 :last_name, 
			 :email, 
			 :password, 
			 :is_shop_owner
			 ) 
			 RETURNING id`

	// Execute query with named parameters
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get the generated ID
	if rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner
			 FROM users 
			 WHERE email = $1 AND password = $2`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
