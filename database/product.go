package database

type Product struct {
	ID          int
	Title       string
	Description string
	Price       int
	ImgaeUrl    string
}
var productList []Product

func Stor(p Product) Product{
	p.ID = len(productList)+1
	productList = append(productList, p)
	return p
}

func List() []Product{
	return productList
}

func Get(pid int) *Product{
	for _, product := range productList{
		if product.ID == pid {
			return &product
		}
	}
	return nil
}

func Update(product Product){
	for idx, p := range productList{
		if p.ID == product.ID{
			productList[idx] = product
		}
	}
}

func Delete(productId int){
	var tempList []Product

	for _, product := range productList{
		if product.ID != productId{
		   tempList = append(tempList, product)
		}
	}
	productList = tempList
}

func init() {
	productList1 := Product{
		Title:       "Mango",
		Description: "Favorite",
		Price:       50,
		ImgaeUrl:    "oewii",
		ID:          1,
	}

	productList2 := Product{
		Title:       "Banana",
		Description: "Little Favorite",
		Price:       10,
		ImgaeUrl:    "oewii",
		ID:          2,
	}

	productList3 := Product{
		Title:       "Jackfrout",
		Description: "Not Favorite",
		Price:       80,
		ImgaeUrl:    "oewii",
		ID:          3,
	}
	productList4 := Product{
		Title:       "Jackfrout",
		Description: "Not Favorite",
		Price:       80,
		ImgaeUrl:    "oewii",
		ID:          4,
	}

	productList = append(productList, productList1)
	productList = append(productList, productList2)
	productList = append(productList, productList3)
	productList = append(productList, productList4)

}