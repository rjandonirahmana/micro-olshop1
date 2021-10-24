package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rjandonirahmana/micro-olshop1/product"
)

type Response struct {
	Meta Meta             `json:"meta"`
	Data product.Products `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Client struct {
	host    string
	timeout time.Duration
}

func NewClientProduct(host string, timeout time.Duration) *Client {
	return &Client{host: host, timeout: timeout}
}

func (c *Client) GetProductByid(id int) product.Products {
	cl := http.Client{
		Timeout: c.timeout,
	}

	reqHeader := fmt.Sprintf("%s/api/v1/product/%d", c.host, id)
	req, err := http.NewRequest("GET", reqHeader, nil)
	if err != nil {
		fmt.Println(err)
		return product.Products{}
	}

	res, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		return product.Products{}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(err)
		return product.Products{}
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return product.Products{}
	}

	var response Response
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return product.Products{}
	}

	return response.Data

}
