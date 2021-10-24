package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	p "github.com/rjandonirahmana/micro-olshop1/handler/product"
	"github.com/rjandonirahmana/micro-olshop1/product"
)

func main() {
	err := godotenv.Load(".env")
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

	repoProduct := product.NewRepoProduct(db)
	serviceProduct := product.NewUsecaseProduct(repoProduct)
	HandlerProduct := p.NewProductHandler(serviceProduct)

	c := gin.Default()
	api := c.Group("/api/v1")

	api.GET("/product/:id", HandlerProduct.GetProductByID)
	api.GET("/productcategory", HandlerProduct.GetProductByCategory)
	api.GET("/search/:keyword", HandlerProduct.SearchProduct)

	c.Run(":6060")

}
