package service

import (
	"errors"

	"github.com/rjandonirahmana/micro-olshop1/model"
	"github.com/rjandonirahmana/micro-olshop1/repository"
)

type serviceProduct struct {
	repository repository.RepoProduct
}

func NewUsecaseProduct(repo repository.RepoProduct) *serviceProduct {
	return &serviceProduct{repository: repo}
}

type ServiceProductInt interface {
	GetProductCategory(id uint) ([]model.Product, error)
	GetProductByid(id uint) (model.Products, error)
	SearchByCategoryByOrder(keyword string, category, order uint) ([]model.Product, error)
	InsertNewProduct(name, desc string, price uint32, qty uint, seller_id uint, category_id uint) (model.Product, error)
}

func (s *serviceProduct) GetProductCategory(id uint) ([]model.Product, error) {

	products, err := s.repository.GetByCategoryID(id)
	if err != nil {
		return products, err
	}

	if len(products) == 0 {
		return products, errors.New("products is not found")
	}

	return products, nil
}

func (s *serviceProduct) GetProductByid(id uint) (model.Products, error) {
	product, err := s.repository.GetProductByID(id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *serviceProduct) SearchByCategoryByOrder(keyword string, category, order uint) ([]model.Product, error) {
	products, err := s.repository.SearchAndByorder(keyword, category, order)
	if err != nil {
		return products, err
	}

	if len(products) == 0 {
		return products, errors.New("there's no product found")
	}

	return products, nil
}

func (s *serviceProduct) InsertNewProduct(name, desc string, price uint32, qty uint, seller_id uint, category_id uint) (model.Product, error) {
	product := model.Product{
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
