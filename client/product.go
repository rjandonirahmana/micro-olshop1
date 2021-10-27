package client

import (
	"bytes"
	"encoding/json"
	"errors"
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

type ResponseProduct struct {
	Meta Meta            `json:"meta"`
	Data product.Product `json:"data"`
}

type responseProducts struct {
	Meta Meta              `json:"meta"`
	Data []product.Product `json:"data"`
}

type Client struct {
	host    string
	timeout time.Duration
}

func NewClientProduct(host string, timeout time.Duration) *Client {
	return &Client{host: host, timeout: timeout}
}

type ProductInt interface {
	GetProductByid(id int) product.Products
	InsertProduct(input product.InputNewPoduct) (product.Product, error)
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

func (c *Client) InsertProduct(input product.InputNewPoduct) (product.Product, error) {
	client := http.Client{
		Timeout: c.timeout,
	}

	reqBodyProduct, err := json.Marshal(input)
	if err != nil {
		return product.Product{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/newproduct", c.host), bytes.NewBuffer(reqBodyProduct))
	if err != nil {
		return product.Product{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return product.Product{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return product.Product{}, errors.New("not http status ok")
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return product.Product{}, err
	}

	var response ResponseProduct
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return product.Product{}, err
	}

	return response.Data, nil
}

func (c *Client) SearchProduct(keyword, category, order string) ([]product.Product, error) {
	client := http.Client{
		Timeout: c.timeout,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/search/%s?category=%s&order=%s", c.host, keyword, category, order), nil)
	if err != nil {
		return []product.Product{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []product.Product{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []product.Product{}, err
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []product.Product{}, err
	}

	var response responseProducts
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return []product.Product{}, err
	}

	return response.Data, nil

}
