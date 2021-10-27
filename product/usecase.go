package product

import "errors"

type service struct {
	repository RepoProduct
}

func NewUsecaseProduct(repo RepoProduct) *service {
	return &service{repository: repo}
}

type ServiceProductInt interface {
	GetProductCategory(id int) ([]Product, error)
	GetProductByid(id int) (Products, error)
	SearchByCategoryByOrder(keyword string, category, order int) ([]Product, error)
	InsertNewProduct(name, desc string, price uint32, qty uint, seller_id uint, category_id uint) (Product, error)
}

func (s *service) GetProductCategory(id int) ([]Product, error) {

	products, err := s.repository.GetByCategoryID(id)
	if err != nil {
		return []Product{}, err
	}

	if len(products) == 0 {
		return []Product{}, errors.New("products is not found")
	}

	return products, nil
}

func (s *service) GetProductByid(id int) (Products, error) {
	product, err := s.repository.GetProductByID(id)

	if err != nil {
		return Products{}, err
	}

	return product, nil
}

func (s *service) SearchByCategoryByOrder(keyword string, category, order int) ([]Product, error) {
	products, err := s.repository.SearchAndByorder(keyword, category, order)
	if err != nil {
		return []Product{}, err
	}

	if len(products) == 0 {
		return products, errors.New("there's no product found")
	}

	return products, nil
}

func (s *service) InsertNewProduct(name, desc string, price uint32, qty uint, seller_id uint, category_id uint) (Product, error) {
	product := Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Quantity:    qty,
		SellerID:    seller_id,
		CategoryID:  category_id,
	}
	product, err := s.repository.InsertNewProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}
