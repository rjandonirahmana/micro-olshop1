package model

type Product struct {
	ID            uint            `db:"id" json:"id"`
	Name          string          `db:"name" json:"name"`
	Price         uint32          `db:"price" json:"price"`
	Quantity      uint            `db:"quantity" json:"quantity"`
	Description   string          `db:"description" json:"description"`
	Rating        float32         `json:"rating"`
	SellerID      uint            `db:"seller_id" json:"seller_id"`
	CategoryID    uint            `db:"category_id" json:"category_id"`
	Category      ProductCategory `db:"product_category" json:"category,omitempty"`
	ProductImages ProductImage    `db:"product_images" json:"product_images,omitempty"`
}

type ProductImage struct {
	ProductID *uint   `db:"product_id" json:"product_id,omitempty"`
	IsPrimary *uint   `db:"is_primary" json:"is_primary,omitempty"`
	Name      *string `db:"name" json:"name,omitempty"`
}

type ProductCategory struct {
	ID   uint   `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
}

type Products struct {
	Name          string          `db:"name" json:"name"`
	ID            uint            `db:"id" json:"id"`
	Price         uint32          `db:"price" json:"price"`
	Quantity      uint            `db:"quantity" json:"quantity"`
	Rating        float32         `json:"rating"`
	CategoryID    uint            `db:"category_id" json:"-"`
	Category      ProductCategory `db:"product_category" json:"category"`
	Description   string          `db:"description" json:"description,omitempty"`
	SellerID      uint            `db:"seller_id" json:"-"`
	ProductImages []ProductImage  `db:"product_images" json:"product_images,omitempty"`
}

type InputNewPoduct struct {
	Name        string `json:"name" validate:"required" form:"name"`
	Price       uint32 `json:"price" validate:"required" form:"price"`
	Quantity    uint   `json:"qty" validate:"required" form:"qty"`
	Category_id uint   `json:"category_id" validate:"required" form:"category_id"`
	Description string `json:"desc" validate:"required" form:"desc"`
	Seller_id   uint   `json:"seller_id" validate:"required" form:"seller_id"`
}
