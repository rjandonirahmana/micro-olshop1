package elastic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/rjandonirahmana/micro-olshop1/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	testcases := []struct {
		nametest string
		product  *model.Product
		err      error
	}{
		{
			nametest: "1",
			product:  &model.Product{Name: "cobaganti", ID: "47", Price: 10000},
			err:      nil,
		},
	}

	for _, test := range testcases {
		tableProduct, err := NewCreateIndex([]string{"http://localhost:9200"})
		assert.NoError(t, err)
		err = tableProduct.CreateIndex("product")
		assert.Nil(t, err)

		repoProduct := NewElasticRepo(*tableProduct, time.Second*10)
		hasil, err := repoProduct.UpdateProduct(context.Background(), test.product)

		fmt.Println(err)

		fmt.Println(hasil)
	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		nametest  string
		idProduct string
		err       error
	}{
		{
			nametest:  "1",
			idProduct: "47",
			err:       nil,
		},
	}

	for _, test := range testcases {
		tableProduct, err := NewCreateIndex([]string{"http://localhost:9200"})
		assert.NoError(t, err)
		err = tableProduct.CreateIndex("product")
		assert.NoError(t, err)
		repoProduct := NewElasticRepo(*tableProduct, time.Second*10)
		hasil, err := repoProduct.GetProductByID(context.Background(), test.idProduct)

		fmt.Println(err)

		fmt.Println(hasil)
	}
}
