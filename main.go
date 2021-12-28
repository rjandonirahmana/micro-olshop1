package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	grpcproduct "github.com/rjandonirahmana/micro-olshop1/grpc/product"
	"github.com/rjandonirahmana/micro-olshop1/handler/product"
	"github.com/rjandonirahmana/micro-olshop1/ongkir"
	"github.com/rjandonirahmana/micro-olshop1/repository"
	"github.com/rjandonirahmana/micro-olshop1/service"
	"google.golang.org/grpc"
)

func main() {

	// tableProduct, err := elastic.NewCreateIndex([]string{"http://localhost:9200"})

	// if err != nil {
	// 	panic(err)
	// }
	// err = tableProduct.CreateIndex("product")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// repoProduct := elastic.NewElasticRepo(*tableProduct, time.Second*10)

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

	SQLrepoProduct := repository.NewRepoProduct(db)
	serviceProduct := service.NewUsecaseProduct(SQLrepoProduct)
	HandlerProduct := product.NewProductHandler(serviceProduct)

	handlerOngkir := ongkir.NewOngkir("21a40172jdjdhdbsgs", 3*time.Second)

	go func() {
		listen, err := net.Listen("tcp", ":10010")
		if err != nil {
			log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcproduct.RegisterUserServiceServer(grpcServer, service.NewGRPCProduct(serviceProduct))

		log.Println("gRPC server is running...")
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	c := gin.Default()
	api := c.Group("/api/v1")

	api.GET("/product/:id", HandlerProduct.GetProductByID)
	api.GET("/product/category", HandlerProduct.GetProductByCategory)
	api.GET("/search/:keyword", HandlerProduct.SearchProduct)
	api.POST("/newproduct", HandlerProduct.InsertNewProduct)
	api.PUT("/product", HandlerProduct.UpdateProduct)
	// api.GET("/product", HandlerProduct.GetProductsByname)

	api.GET("/cost", handlerOngkir.CekOngkir)
	api.GET("/city", handlerOngkir.GetCity)

	c.Run(":6262")

}
