package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjandonirahmana/micro-olshop1/handler"
	"github.com/rjandonirahmana/micro-olshop1/product"
)

type HandlerProduct struct {
	service product.ServiceProductInt
}

func NewProductHandler(service product.ServiceProductInt) *HandlerProduct {
	return &HandlerProduct{service: service}
}

func (h *HandlerProduct) GetProductByCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response := handler.APIResponse(err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	products, err := h.service.GetProductCategory(id)
	if err != nil {
		response := handler.APIResponse(err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("success", http.StatusOK, "products", products)
	c.JSON(http.StatusOK, response)

}

func (h *HandlerProduct) GetProductByID(c *gin.Context) {
	product_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	product, err := h.service.GetProductByid(product_id)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if product.ID == 0 {
		response := handler.APIResponse("failed", http.StatusBadRequest, "product not found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := handler.APIResponse("success", http.StatusOK, "success get product", product)
	c.JSON(http.StatusOK, response)

}

func (h *HandlerProduct) SearchProduct(c *gin.Context) {
	keyword := c.Param("keyword")
	category, _ := strconv.Atoi(c.Query("category"))
	order, _ := strconv.Atoi(c.Query("order"))

	fmt.Println(keyword)

	product, err := h.service.SearchByCategoryByOrder(keyword, category, order)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("success", http.StatusOK, "success get product", product)
	c.JSON(http.StatusOK, response)
}
