package elastic

//put localhost/9200/olshopala_products
var products string = ` {
	"settings":{
		"number_of_shards": 3,
		"number_of_replicas": 1
	},
	"mappings":{
		"records":{
			"properties":{
				"id": {
					"type" : "interger"
				},
				"name":{
					"type":"keyword"
				},
				"price":{
					"type":"interger"
				},
				"rating":{
					"type":"float"
				},
				"seller_id":{
					"type":"interger"
				},
				"category_id":{
					"type":"date"
				},
				"product_images" : {
					"type" : "nested",
					"properties" : {
						"name" : {
							"type" : "keyword"
						},
						"is_primary" : {
							"type" : "bool"
						},
						"product_id" : {
							"type" : "interger"
						}

						
					}
				},
				"description" : {
					"type" : "text"
				}

			}
		}
	}
}`

// Price         uint32          `db:"price" json:"price"`
// 	Quantity      uint            `db:"quantity" json:"quantity"`
// 	Description   string          `db:"description" json:"description"`
// 	Rating        float32         `json:"rating"`
// 	SellerID      uint            `db:"seller_id" json:"seller_id"`
// 	CategoryID    uint            `db:"category_id" json:"category_id"`
// 	Category      ProductCategory `db:"product_category" json:"category,omitempty"`
// 	ProductImages ProductImage    `db:"product_images" json:"product_images,omitempty"`
