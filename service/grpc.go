package service

import (
	"context"
	"log"

	"github.com/rjandonirahmana/micro-olshop1/grpc/product"
)

type GRPCProduct struct {
	product.UnimplementedUserServiceServer
	service ServiceProductInt
}

func NewGRPCProduct(service ServiceProductInt) *GRPCProduct {
	return &GRPCProduct{service: service}
}

// func (c *GRPCProduct) GetProductByID(ctx context.Context, req *product.GetProductByIDReq) (*product.GetProductByIDOutput, error) {

// 	id := req.ID
// 	prod, err := c.service.GetProductByid(&id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &product.GetProductByIDOutput{
// 		ID:          req.ID,
// 		Price:       int32(prod.Price),
// 		Quantity:    int32(prod.Quantity),
// 		Rating:      prod.Rating,
// 		CategoryID:  int32(prod.CategoryID),
// 		Description: prod.Description,
// 		SellerID:    prod.SellerID,
// 		Category: &product.ProductCategory{
// 			ID:   int32(prod.Category.ID),
// 			Name: prod.Category.Name,
// 		},
// 	}, nil
// }

func (c *GRPCProduct) InsertNewProduct(ctx context.Context, req *product.InputNewProduct) (*product.GetProductByIDOutput, error) {
	prouctRes, err := c.service.InsertNewProduct(ctx, req.Name, req.Description, req.Price, uint(req.Quantity), req.SellerId, uint(req.CategoryId))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &product.GetProductByIDOutput{
		ID:         prouctRes.ID,
		Price:      int32(prouctRes.Price),
		Quantity:   int32(prouctRes.Quantity),
		Rating:     prouctRes.Rating,
		CategoryID: int32(prouctRes.CategoryID),
		Category: &product.ProductCategory{
			ID:   int32(prouctRes.CategoryID),
			Name: prouctRes.Category.Name,
		},
		SellerID: prouctRes.SellerID,
	}, nil

}
