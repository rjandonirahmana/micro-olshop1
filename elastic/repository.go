package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/rjandonirahmana/micro-olshop1/model"
)

type repository struct {
	elastic Index
	timeout time.Duration
}

func NewElasticRepo(elastic Index, timeduration time.Duration) *repository {
	return &repository{elastic: elastic, timeout: timeduration}
}

type RepoProduct interface {
	InsertProduct(ctx context.Context, product *model.Product) error
	GetProductByID(ctx context.Context, id string) (model.Product, error)
	GetProductByName(ctx context.Context, product *string, categoryID *uint) ([]*model.Product, error)
	UpdateProduct(ctx context.Context, product model.Product) (model.Product, error)
}

type Storage struct {
	Source interface{} `json:"_source"`
}

func (r *repository) InsertProduct(ctx context.Context, product *model.Product) error {
	reqBody, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req := esapi.CreateRequest{
		Index:      r.elastic.indexName,
		DocumentID: fmt.Sprint(product.ID),
		Body:       bytes.NewBuffer(reqBody),
	}

	ctx1, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := req.Do(ctx1, r.elastic.client)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 400 {
		return fmt.Errorf("status code is not 400")
	}

	if res.IsError() {
		return fmt.Errorf("insert: response: %s", res.String())
	}

	fmt.Println(res.Body)
	return nil

}

func (r *repository) GetProductByID(ctx context.Context, id string) (model.Product, error) {
	// p.Elastic.client.Get()

	req := esapi.GetRequest{
		Index:      r.elastic.indexName,
		DocumentID: id,
	}

	ctxTime, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res, err := req.Do(ctxTime, r.elastic.client)
	if err != nil {
		return model.Product{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return model.Product{}, err
	}
	if res.StatusCode == 404 {
		return model.Product{}, fmt.Errorf("cannot find")
	}

	var (
		storage model.Product
		body    Storage
	)

	body.Source = &storage

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return model.Product{}, fmt.Errorf("find one: decode: %w", err)
	}

	return storage, nil

}

func (r *repository) UpdateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	reqBody, err := json.Marshal(product)
	if err != nil {
		return product, err
	}

	req := esapi.UpdateRequest{
		Index:      r.elastic.indexName,
		DocumentID: fmt.Sprint(product.ID),
		Body:       bytes.NewBuffer(reqBody),
	}

	res, err := req.Do(ctx, r.elastic.client)
	if err != nil {
		return product, err
	}

	defer res.Body.Close()
	if res.StatusCode == 404 {
		return product, errors.New("value not found")
	}

	if res.IsError() {
		return product, fmt.Errorf("update: response: %s", res.String())
	}

	var (
		body    model.Product
		storage Storage
	)

	storage.Source = &body

	if err := json.NewDecoder(res.Body).Decode(&storage); err != nil {
		return model.Product{}, fmt.Errorf("find one: decode: %w", err)
	}

	return body, nil
}

func (r *repository) GetProductByName(ctx context.Context, product *string, categoryID *uint) ([]*model.Product, error) {

	var query string
	var buf bytes.Buffer
	if *categoryID != 0 {
		query = `{
		"query": {
		  "bool": {
			"must": [
			{ "term": { "category_id": %d}}
			],
			"should": [
			{ "term": { "name": "%s"}},
			{ "prefix": { "name": ""}}
			]
		  }
		}
	  }`

		buf.WriteString(fmt.Sprintf(query, *categoryID, *product))
	} else {
		query = `{
			"query": {
			  "bool": {
				"should": [
				{ "term": { "name": "%s"}},
				{ "prefix": { "name": ""}}
				]
			  }
			}
		  }`

		buf.WriteString(fmt.Sprintf(query, *product))
	}

	fmt.Println(buf.String())
	req := esapi.SearchRequest{
		Index: []string{r.elastic.indexName},
		Body:  strings.NewReader(buf.String()),
	}

	resp, err := req.Do(ctx, r.elastic.client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, err
	}

	var hits struct {
		Hits struct {
			Hits []struct {
				model.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&hits); err != nil {
		fmt.Println("Error here", err)
		return nil, fmt.Errorf("%v", err)
	}

	res := make([]*model.Product, len(hits.Hits.Hits))

	for _, hit := range hits.Hits.Hits {
		pr := &model.Product{
			ID:          hit.ID,
			Name:        hit.Name,
			Price:       hit.Price,
			Quantity:    hit.Quantity,
			Description: hit.Description,
			Rating:      hit.Rating,
			SellerID:    hit.SellerID,
			CategoryID:  hit.CategoryID,
		}

		res = append(res, pr)
		// res[i].ID = hit.ID
		// res[i].Name = hit.Name
		// res[i].Price = hit.Price
		// res[i].Quantity = hit.Quantity
		// res[i].Description = hit.Description
		// res[i].Rating = hit.Rating
		// res[i].SellerID = hit.SellerID
		// res[i].CategoryID = hit.CategoryID

		// res[i].Priority = internal.Priority(hit.Source.Priority)
		// res[i].Dates.Due = time.Unix(0, hit.Source.DateDue).UTC()
		// res[i].Dates.Start = time.Unix(0, hit.Source.DateStart).UTC()
	}

	return res, nil
}

// func (r *repository) SearchProducts(name *string, )
