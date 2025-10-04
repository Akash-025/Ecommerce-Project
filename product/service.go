package product

import "practice/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.productRepo.Create(p)
}
func (svc *service) Get(pid int) (*domain.Product, error){
	return svc.productRepo.Get(pid)
}
func (svc *service) List() ([]*domain.Product, error){
	return svc.productRepo.List()
}
func (svc *service) Update(product domain.Product) (*domain.Product, error){
	return svc.productRepo.Update(product)
}
func (svc *service) Delete(productId int) error{
	return svc.productRepo.Delete(productId)
}
