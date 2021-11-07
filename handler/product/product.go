package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rjandonirahmana/micro-olshop1/handler"
	"github.com/rjandonirahmana/micro-olshop1/model"
	"github.com/rjandonirahmana/micro-olshop1/service"
)

type HandlerProduct struct {
	service service.ServiceProductInt
}

func NewProductHandler(service service.ServiceProductInt) *HandlerProduct {
	return &HandlerProduct{service: service}
}

func (h *HandlerProduct) GetProductByCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response := handler.APIResponse(err.Error(), http.StatusInternalServerError, "failed", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	products, err := h.service.GetProductCategory(uint(id))
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
	product, err := h.service.GetProductByid(uint(product_id))
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

	product, err := h.service.SearchByCategoryByOrder(keyword, uint(category), uint(order))
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("success", http.StatusOK, "success get product", product)
	c.JSON(http.StatusOK, response)
}

func (h *HandlerProduct) InsertNewProduct(c *gin.Context) {
	var input model.InputNewPoduct

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	product, err := h.service.InsertNewProduct(input.Name, input.Description, input.Price, input.Quantity, input.Seller_id, input.Category_id)
	if err != nil {
		response := handler.APIResponse("failed to insert prouct", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("success insert product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *HandlerProduct) UpdateProduct(c *gin.Context) {

	name := c.Request.FormValue("name")
	desc := c.Request.FormValue("desc")
	productID, _ := strconv.Atoi(c.Request.FormValue("product_id"))
	sellerID, _ := strconv.Atoi(c.Request.FormValue("seller_id"))
	price, _ := strconv.Atoi(c.Request.FormValue("price"))
	qty, _ := strconv.Atoi(c.Request.FormValue("qty"))
	categoryID, _ := strconv.Atoi(c.Request.FormValue("category_id"))

	product, err := h.service.UpdateProduct(uint(productID), uint(sellerID), name, desc, uint32(price), uint(qty), uint(categoryID))
	if err != nil {
		response := handler.APIResponse("failed to insert prouct", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("successfully update product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)

}

func (h *HandlerProduct) GetProductsByname(c *gin.Context) {
	keyword := c.Request.FormValue("keyword")
	categoryID, _ := strconv.Atoi(c.Request.FormValue("category"))
	category := uint(categoryID)

	products, err := h.service.GetProductByName(&keyword, &category)
	if err != nil {
		response := handler.APIResponse("failed to insert prouct", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("successfully get products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)

}
