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
	product_id := c.Param("id")
	product, err := h.service.GetProductByid(&product_id)
	if err != nil {
		response := handler.APIResponse("failed", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if product.ID == "" {
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
	category_id := uint(category)
	order_id := uint(order)

	product, err := h.service.SearchByCategoryByOrder(&keyword, &category_id, &order_id)
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

	product, err := h.service.InsertNewProduct(c.Request.Context(), input.Name, input.Description, input.Price, input.Quantity, input.Seller_id, input.Category_id)
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
	productID := c.Request.FormValue("product_id")
	sellerID := c.Request.FormValue("seller_id")
	price, _ := strconv.ParseInt((c.Request.FormValue("price")), 10, 32)
	qty, _ := strconv.Atoi(c.Request.FormValue("qty"))
	categoryID, _ := strconv.Atoi(c.Request.FormValue("category_id"))

	price32 := uint32(price)
	quantity := uint(qty)
	category := uint(categoryID)

	product, err := h.service.UpdateProduct(&productID, &sellerID, &name, &desc, &price32, &quantity, &category)
	if err != nil {
		response := handler.APIResponse("failed to insert prouct", http.StatusUnprocessableEntity, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := handler.APIResponse("successfully update product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)

}

// func (h *HandlerProduct) GetProductsByname(c *gin.Context) {
// 	keyword := c.Request.FormValue("keyword")
// 	categoryID, _ := strconv.Atoi(c.Request.FormValue("category"))
// 	category := uint(categoryID)

// 	products, err := h.service.GetProductByName(&keyword, &category)
// 	if err != nil {
// 		response := handler.APIResponse("failed to insert prouct", http.StatusUnprocessableEntity, err.Error(), nil)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	response := handler.APIResponse("successfully get products", http.StatusOK, "success", products)
// 	c.JSON(http.StatusOK, response)

// }
