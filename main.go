package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rjandonirahmana/micro-olshop1/elastic"
	"github.com/rjandonirahmana/micro-olshop1/model"
)

// _ "github.com/go-sql-driver/mysql"

func main() {

	ctx := context.Background()
	tableProduct := elastic.NewCreateIndex([]string{"http://localhost:9200"})
	err := tableProduct.CreateIndex("product")
	if err != nil {
		fmt.Println(err)
	}

	repoProduct := elastic.NewElasticRepo(*tableProduct, time.Second*10)
	err = repoProduct.InsertProduct(ctx, model.Product{
		ID:          10,
		Name:        "ngasal",
		CategoryID:  1,
		Quantity:    10,
		Description: "yahhh coba",
		Rating:      0,
		SellerID:    1,
	})

	if err != nil {
		fmt.Println(err)
	}

	product, err := repoProduct.GetProductByID(ctx, "3")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(product)

	products, err := repoProduct.GetProductByName(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(products)

	updatedProduct, err := repoProduct.UpdateProduct(ctx, model.Product{
		ID:          1,
		Name:        "ngasalahh111",
		CategoryID:  1,
		Quantity:    10,
		Description: "ni product baru ni",
		Rating:      0,
		SellerID:    1,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(updatedProduct)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dbUserName := os.Getenv("DB_username")
	// dbName := os.Getenv("DB_name")
	// dbPass := os.Getenv("DB_password")
	// dbstring := fmt.Sprintf("%s:%s@(localhost:3300)/%s?parseTime=true", dbUserName, dbPass, dbName)

	// db, err := sqlx.Connect("mysql", dbstring)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// repoProduct := repository.NewRepoProduct(db)
	// serviceProduct := service.NewUsecaseProduct(repoProduct)
	// HandlerProduct := p.NewProductHandler(serviceProduct)

	// c := gin.Default()
	// api := c.Group("/api/v1")

	// api.GET("/product/:id", HandlerProduct.GetProductByID)
	// api.GET("/productcategory", HandlerProduct.GetProductByCategory)
	// api.GET("/search/:keyword", HandlerProduct.SearchProduct)
	// api.POST("/newproduct", HandlerProduct.InsertNewProduct)

	// c.Run(":6262")

}
