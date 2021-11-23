package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

// type responseProduct struct {
// 	Meta Meta          `json:"meta"`
// 	Data model.Product `json:"data"`
// }

type Client struct {
	host    string
	timeout time.Duration
}

type Storage struct {
	Source interface{} `json:"data"`
}

func NewClientProduct(host string, timeout time.Duration) *Client {
	return &Client{host: host, timeout: timeout}
}

type ProductInt interface {
	GetProductByid(id uint) (*model.Products, error)
	InsertProduct(input model.InputNewPoduct) (*model.Product, error)
	SearchProduct(keyword, category, order string) ([]*model.Product, error)
}

func (c *Client) GetProductByid(id uint) (*model.Products, error) {
	cl := &http.Client{
		Timeout: c.timeout,
	}

	reqHeader := fmt.Sprintf("%s/api/v1/product/%d", c.host, id)
	req, err := http.NewRequest("GET", reqHeader, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(err)
		return nil, err
	}

	var (
		response model.Products
		resource Storage
	)

	resource.Source = &response

	err = json.NewDecoder(res.Body).Decode(&resource)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (c *Client) InsertProduct(input model.InputNewPoduct) (*model.Product, error) {
	client := &http.Client{
		Timeout: c.timeout,
	}

	reqBodyProduct, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/newproduct", c.host), bytes.NewBuffer(reqBodyProduct))
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Body)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("not http status ok")
	}

	var (
		response model.Product
		storage  Storage
	)

	storage.Source = &response

	err = json.NewDecoder(res.Body).Decode(&storage)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SearchProduct(keyword, category, order string) ([]*model.Product, error) {
	client := &http.Client{
		Timeout: c.timeout,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/search/%s?category=%s&order=%s", c.host, keyword, category, order), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("not http status ok")
	}

	var (
		Data   []*model.Product
		Source Storage
	)

	Source.Source = &Data

	err = json.NewDecoder(res.Body).Decode(&Source)
	if err != nil {
		return nil, err
	}
	return Data, nil

}
