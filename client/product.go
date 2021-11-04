package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rjandonirahmana/micro-olshop1/model"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// type responseProducts struct {
// 	Meta Meta              `json:"meta"`
// 	Data []product.Product `json:"data"`
// }

type Client struct {
	host    string
	timeout time.Duration
}

func NewClientProduct(host string, timeout time.Duration) *Client {
	return &Client{host: host, timeout: timeout}
}

type ProductInt interface {
	GetProductByid(id int) model.Products
	InsertProduct(input model.InputNewPoduct) (model.Product, error)
	SearchProduct(keyword, category, order string) ([]model.Product, error)
}

func (c *Client) GetProductByid(id int) model.Products {
	cl := http.Client{
		Timeout: c.timeout,
	}

	reqHeader := fmt.Sprintf("%s/api/v1/product/%d", c.host, id)
	req, err := http.NewRequest("GET", reqHeader, nil)
	if err != nil {
		fmt.Println(err)
		return model.Products{}
	}

	res, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		return model.Products{}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(err)
		return model.Products{}
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return model.Products{}
	}

	var response Response
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return model.Products{}
	}

	return response.Data.(model.Products)

}

func (c *Client) InsertProduct(input model.InputNewPoduct) (model.Product, error) {
	client := http.Client{
		Timeout: c.timeout,
	}

	reqBodyProduct, err := json.Marshal(input)
	if err != nil {
		return model.Product{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/newproduct", c.host), bytes.NewBuffer(reqBodyProduct))
	if err != nil {
		return model.Product{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return model.Product{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return model.Product{}, errors.New("not http status ok")
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Product{}, err
	}

	var response Response
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return model.Product{}, err
	}

	return response.Data.(model.Product), nil
}

func (c *Client) SearchProduct(keyword, category, order string) ([]model.Product, error) {
	client := http.Client{
		Timeout: c.timeout,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/search/%s?category=%s&order=%s", c.host, keyword, category, order), nil)
	if err != nil {
		return []model.Product{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []model.Product{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return []model.Product{}, err
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []model.Product{}, err
	}

	var response Response
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return []model.Product{}, err
	}

	return response.Data.([]model.Product), nil

}
