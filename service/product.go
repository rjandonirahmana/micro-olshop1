package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rjandonirahmana/micro-olshop1/model"
	"github.com/rjandonirahmana/micro-olshop1/repository"
)

type serviceProduct struct {
	repository repository.RepoProduct
	// elasticRepo elastic.RepoProduct
}

func NewUsecaseProduct(repo repository.RepoProduct) *serviceProduct {
	return &serviceProduct{repository: repo}
}

type ServiceProductInt interface {
	GetProductCategory(id uint) ([]*model.Product, error)
	GetProductByid(id *string) (*model.Products, error)
	SearchByCategoryByOrder(keyword *string, category, order *uint) ([]*model.Product, error)
	InsertNewProduct(ctx context.Context, name, desc string, price uint32, qty uint, seller_id string, category_id uint) (*model.Product, error)
	UpdateProduct(productID *string, sellerID *string, name, desc *string, price *uint32, qty, categoryID *uint) (*model.Product, error)
	// GetProductByName(product *string, categoryID *uint) ([]*model.Product, error)
}

func (s *serviceProduct) GetProductCategory(id uint) ([]*model.Product, error) {

	products, err := s.repository.GetByCategoryID(id)
	if err != nil {
		return products, err
	}

	if len(products) == 0 {
		return products, errors.New("products is not found")
	}

	return products, nil
}

func (s *serviceProduct) GetProductByid(id *string) (*model.Products, error) {
	product, err := s.repository.GetProductByID(id)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *serviceProduct) SearchByCategoryByOrder(keyword *string, category, order *uint) ([]*model.Product, error) {
	products, err := s.repository.SearchAndByorder(keyword, category, order)
	if err != nil {
		return products, err
	}

	if len(products) == 0 {
		return products, errors.New("there's no product found")
	}

	return products, nil
}

func (s *serviceProduct) InsertNewProduct(ctx context.Context, name, desc string, price uint32, qty uint, seller_id string, category_id uint) (*model.Product, error) {
	product := &model.Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Quantity:    qty,
		SellerID:    seller_id,
		CategoryID:  category_id,
		ID:          uuid.New().String(),
	}
	product, err := s.repository.InsertNewProduct(ctx, product)
	if err != nil {
		return product, err
	}

	// err = s.elasticRepo.InsertProduct(context.Background(), product)
	// if err != nil {
	// 	s.repository.DeleteByID(product.ID)
	// 	return product, fmt.Errorf("error : %v and cannot insert to elastic please try insert again", err)
	// }

	return product, nil
}

func (s *serviceProduct) UpdateProduct(productID *string, sellerID *string, name, desc *string, price *uint32, qty, categoryID *uint) (*model.Product, error) {
	product, err := s.repository.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	if product.SellerID != *sellerID {
		return nil, fmt.Errorf("can not update info product this is not your product")
	}

	if *name == "" {
		*name = product.Name
	}
	if *desc == "" {
		*desc = product.Description
	}
	if *price == 0 {
		*price = product.Price
	}
	if *qty == 0 {
		*qty = product.Quantity
	}
	if *categoryID == 0 {
		*categoryID = product.CategoryID
	}

	updatedProduct := &model.Product{
		ID:          product.ID,
		Name:        *name,
		Price:       *price,
		Description: *desc,
		Quantity:    *qty,
		Rating:      product.Rating,
		SellerID:    product.SellerID,
		CategoryID:  *categoryID,
	}

	updatedProduct, err = s.repository.UpdateProduct(updatedProduct)
	if err != nil {
		return updatedProduct, err
	}

	// updatedProduct, err = s.elasticRepo.UpdateProduct(context.Background(), updatedProduct)
	// if err != nil {
	// 	return nil, err
	// }

	return updatedProduct, nil
}

// func (s *serviceProduct) GetProductByName(product *string, categoryID *uint) ([]*model.Product, error) {
// 	// ctx := context.Background()

// 	// // products, err := s.elasticRepo.GetProductByName(ctx, product, categoryID)
// 	// // if err != nil {
// 	// // 	return products, err
// 	// // }

// 	// return products, nil
// }
