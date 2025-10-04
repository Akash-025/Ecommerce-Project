package product

import (
	"practice/domain"
	prdctHndlr "practice/rest/handlers/product"
)

type Service interface{
	prdctHndlr.Service
}


type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(pid int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(productId int) error
}