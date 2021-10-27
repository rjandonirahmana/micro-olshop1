package product

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func initDB() *sqlx.DB {
	err := godotenv.Load("../.env")
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

	return db
}

var (
	db1  = initDB()
	repo = NewRepoProduct(db1)
)

func TestInsertProduct(t *testing.T) {
	testcase := []struct {
		testname string
		product  Product
		err      error
	}{
		{
			testname: "berhasil",
			product:  Product{Name: "coba", Quantity: 1, Price: 1000, Description: "coba", SellerID: 35, CategoryID: 1},
			err:      nil,
		},
	}

	for _, test := range testcase {
		p, err := repo.InsertNewProduct(test.product)
		assert.Equal(t, test.err, err)

		fmt.Println(p)
	}
}
