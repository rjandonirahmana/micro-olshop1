package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rjandonirahmana/micro-olshop1/model"
)

type repository struct {
	db *sqlx.DB
}

type RepoProduct interface {
	GetProductByID(id uint) (model.Products, error)
	GetByCategoryID(id uint) ([]model.Product, error)
	SearchAndByorder(keyword string, category, order uint) ([]model.Product, error)
	InsertNewProduct(product model.Product) (model.Product, error)
}

func NewRepoProduct(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetProductByID(id uint) (model.Products, error) {
	querry := `SELECT p.*, pc.id as "product_category.id", pc.name as "product_category.name" FROM products p INNER JOIN product_category pc ON p.category_id = pc.id  WHERE p.id = ?`

	var product model.Products
	err := r.db.Get(&product, querry, id)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return product, err
	}
	var images []model.ProductImage
	err = r.db.Select(&images, `SELECT * FROM product_images p WHERE p.product_id = ?`, id)
	if err != sql.ErrNoRows {
		return product, err
	}

	product.ProductImages = append(product.ProductImages, images...)

	return product, nil
}

func (r *repository) GetByCategoryID(id uint) ([]model.Product, error) {
	querry := `SELECT p.*, pi.name as "product_images.name", pi.is_primary as "product_images.is_primary", pi.product_id as "product_images.product_id", pc.id as "product_category.id", pc.name as "product_category.name" FROM products p LEFT JOIN product_category pc ON p.category_id = pc.id LEFT JOIN product_images pi ON p.id = pi.product_id WHERE pc.id = ? GROUP BY p.id`

	products := []model.Product{}

	err := r.db.Select(&products, querry, id)
	if err != nil {
		return products, err
	}

	return products, nil

}

func (r *repository) SearchAndByorder(keyword string, category, order uint) ([]model.Product, error) {

	var err error
	var product []model.Product
	if keyword != "%%" && category != 0 && order != 0 {
		err = r.db.Select(&product, `SELECT DISTINCT p.*, pi.name as "product_images.name", pi.is_primary as "product_images.is_primary", pi.product_id as "product_images.product_id", pc.id as "product_category.id", pc.name as "product_category.name" FROM products p LEFT JOIN product_category pc ON p.category_id = pc.id LEFT JOIN product_images pi ON p.id = pi.product_id WHERE p.name LIKE ? AND pc.id = ? GROUP BY p.id ORDER BY p.price`, "%"+keyword+"%", category)
	} else if keyword != "%%" && category != 0 {
		err = r.db.Select(&product, `SELECT DISTINCT p.*, pi.name as "product_images.name", pi.is_primary as "product_images.is_primary", pi.product_id as "product_images.product_id", pc.id as "product_category.id", pc.name as "product_category.name" FROM products p LEFT JOIN product_category pc ON p.category_id = pc.id LEFT JOIN product_images pi ON p.id = pi.product_id WHERE p.name LIKE ? AND pc.id = ? GROUP BY p.id`, "%"+keyword+"%", category)
	} else if keyword != "%%" && order != 0 {
		err = r.db.Select(&product, `SELECT DISTINCT p.*, pi.name as "product_images.name", pi.is_primary as "product_images.is_primary", pi.product_id as "product_images.product_id", pc.id as "product_category.id", pc.name as "product_category.name" FROM products p LEFT JOIN product_category pc ON p.category_id = pc.id LEFT JOIN product_images pi ON p.id = pi.product_id WHERE p.name LIKE ? GROUP BY p.id ORDER BY p.price`, "%"+keyword+"%")
	} else if keyword != "%%" {
		err = r.db.Select(&product, `SELECT DISTINCT p.*, pi.name as "product_images.name", pi.is_primary as "product_images.is_primary", pi.product_id as "product_images.product_id", pc.id as "product_category.id", pc.name as "product_category.name" FROM products p LEFT JOIN product_category pc ON p.category_id = pc.id LEFT JOIN product_images pi ON p.id = pi.product_id WHERE p.name LIKE ? GROUP BY p.id`, "%"+keyword+"%")
	}

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) InsertNewProduct(product model.Product) (model.Product, error) {
	querry := `INSERT INTO products (name, price, quantity, description, category_id, seller_id) VALUES (?,?,?,?,?,?)`

	_, err := r.db.Exec(querry, product.Name, product.Price, product.Quantity, product.Description, product.CategoryID, product.SellerID)
	if err != nil {
		return product, err
	}
	querry = `SELECT last_insert_id()`
	err = r.db.Get(&product.ID, querry)
	if err != nil {
		return product, err
	}

	return product, nil

}