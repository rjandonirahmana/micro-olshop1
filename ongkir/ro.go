package ongkir

import (
	"net/http"
	"time"

	ro "github.com/Bhinneka/go-rajaongkir"
	"github.com/gin-gonic/gin"
	"github.com/rjandonirahmana/micro-olshop1/handler"
)

type Ongkir struct {
	Key       string
	timeLimit time.Duration
}

func NewOngkir(key string, limit time.Duration) *Ongkir {

	return &Ongkir{Key: key, timeLimit: limit}
}

type RequestOngkir struct {
	CityID      string `json:"city_id"`
	ProvinceID  string `json:"province_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

func (o *Ongkir) CekOngkir(c *gin.Context) {
	raja := ro.New(o.Key, o.timeLimit)

	var req *RequestOngkir
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	querry := ro.QueryRequest{Origin: req.Origin, Destination: req.Destination, Weight: req.Weight, Courier: req.Courier}
	result := raja.GetCost(querry)

	if result.Error != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, result.Error.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	cost, ok := result.Result.(ro.Cost)
	if !ok {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, "result is not cost", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("success", http.StatusOK, "success get result", cost)
	c.JSON(http.StatusOK, response)

}

func (o *Ongkir) GetCity(c *gin.Context) {
	province := c.Query("province")

	raja := ro.New(o.Key, o.timeLimit)

	q := ro.QueryRequest{ProvinceID: province}
	result := raja.GetCity(q)

	response := handler.APIResponse("success", http.StatusOK, "success get result", result)
	c.JSON(http.StatusOK, response)
}
