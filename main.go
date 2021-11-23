package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/rjandonirahmana/micro-olshop1/elastic"
	"github.com/rjandonirahmana/micro-olshop1/handler/product"
	"github.com/rjandonirahmana/micro-olshop1/repository"
	"github.com/rjandonirahmana/micro-olshop1/service"
)

func main() {

	tableProduct, err := elastic.NewCreateIndex([]string{"http://localhost:9200"})

	if err != nil {
		panic(err)
	}
	err = tableProduct.CreateIndex("product")
	if err != nil {
		fmt.Println(err)
	}

	repoProduct := elastic.NewElasticRepo(*tableProduct, time.Second*10)

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUserName := os.Getenv("DB_username")
	dbName := os.Getenv("DB_name")
	dbPass := os.Getenv("DB_password")
	dbstring := fmt.Sprintf("%s:%s@(localhost:3306)/%s?parseTime=true", dbUserName, dbPass, dbName)

	db, err := sqlx.Connect("mysql", dbstring)
	if err != nil {
		log.Fatalln(err)
	}

	SQLrepoProduct := repository.NewRepoProduct(db)
	serviceProduct := service.NewUsecaseProduct(SQLrepoProduct, repoProduct)
	HandlerProduct := product.NewProductHandler(serviceProduct)

	c := gin.Default()
	api := c.Group("/api/v1")

	api.GET("/product/:id", HandlerProduct.GetProductByID)
	api.GET("/product/category", HandlerProduct.GetProductByCategory)
	api.GET("/search/:keyword", HandlerProduct.SearchProduct)
	api.POST("/newproduct", HandlerProduct.InsertNewProduct)
	api.PUT("/product", HandlerProduct.UpdateProduct)
	api.GET("/product", HandlerProduct.GetProductsByname)

	c.Run(":6262")

}
