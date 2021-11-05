package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	InsertProduct(ctx context.Context, product model.Product) error
}

type Storage struct {
	Source interface{} `json:"_source"`
}

func (r *repository) InsertProduct(ctx context.Context, product model.Product) error {
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

func (r *repository) GetProductByName(ctx context.Context) ([]model.Product, error) {

	query := `{"query": {"match_all" : {}},"size": 2}`
	var b strings.Builder
	b.WriteString(query)
	// convert string to reader
	reqBody := strings.NewReader(b.String())

	search := r.elastic.client
	res, err := search.Search(
		search.Search.WithContext(ctx),
		search.Search.WithIndex(r.elastic.indexName),
		search.Search.WithBody(reqBody),
		search.Search.WithTrackTotalHits(true),
		search.Search.WithPretty(),
	)
	if err != nil {
		return []model.Product{}, err
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		return []model.Product{}, fmt.Errorf("not found")
	}

	var (
		storage []model.Product
		body    Storage
	)

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []model.Product{}, err
	}

	fmt.Println(string(bodyRes))

	body.Source = &storage

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return []model.Product{}, fmt.Errorf("find one: decode: %w", err)
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
