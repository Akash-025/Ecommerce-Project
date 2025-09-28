package repo

type Product struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int `json:"price"`
	ImgaeUrl    string `json:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (Product, error)
	Get(pid int) (*Product, error)
	List() ([]*Product, error)
	Update(product Product) (*Product, error)
	Delete(productId int) error
}

type productRepo struct {
	productList []*Product
}

// Constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generate(repo)
	return repo 
}


func (r *productRepo) Create(p Product) (Product, error) {
	p.ID = len(r.productList)+1
	r.productList = append(r.productList, &p)
	return p, nil
}
func (r *productRepo) Get(pid int) (*Product, error) {
	for _, product := range r.productList{
		if product.ID == pid {
			return product,nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList{
		if p.ID == product.ID{
			r.productList[idx] = &product
		}
	}
	return &product, nil
}
func (r *productRepo) Delete(productId int) error {
	var tempList []*Product

	for _, product := range r.productList{
		if product.ID != productId{
		   tempList = append(tempList, product)
		}
	}
	r.productList = tempList
	return nil
}

func generate(r *productRepo) {
	productList1 := &Product{
		Title:       "Mango",
		Description: "Favorite",
		Price:       50,
		ImgaeUrl:    "oewii",
		ID:          1,
	}

	productList2 := &Product{
		Title:       "Banana",
		Description: "Little Favorite",
		Price:       10,
		ImgaeUrl:    "oewii",
		ID:          2,
	}

	productList3 := &Product{
		Title:       "Jackfrout",
		Description: "Not Favorite",
		Price:       80,
		ImgaeUrl:    "oewii",
		ID:          3,
	}
	productList4 := &Product{
		Title:       "Jackfrout",
		Description: "Not Favorite",
		Price:       80,
		ImgaeUrl:    "oewii",
		ID:          4,
	}

	r.productList = append(r.productList, productList1)
	r.productList = append(r.productList, productList2)
	r.productList = append(r.productList, productList3)
	r.productList = append(r.productList, productList4)

}
