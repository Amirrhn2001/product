package api

import (
	"encoding/json"
	"io"
	"net/http"
	"product/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler interface {
	CreateProduct(c *gin.Context)
	GetOneProductById(c *gin.Context)
	GetProductsByFilter(c *gin.Context)
	UpdateProductById(c *gin.Context)
	DeleteProductById(c *gin.Context)
}

type handler struct {
	Service domain.Service
}

func NewHandler(service domain.Service) Handler {
	return &handler{Service: service}
}

func (h handler) CreateProduct(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	product := &domain.Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, Err := h.Service.CreateProductService(product)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h handler) GetOneProductById(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	id := c.Param("uuid")
	data, Err := h.Service.GetOneProductByIdService(id)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) GetProductsByFilter(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	filter := []domain.Filter{}
	err := json.Unmarshal([]byte(c.Query("filter")), &filter)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	pagination := &domain.Pagination{}
	err = json.Unmarshal([]byte(c.Query("pagination")), pagination)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	data, _, Err := h.Service.GetProductsByFilterService(filter, pagination)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) UpdateProductById(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	product := &domain.Product{}
	err = json.Unmarshal(body, product)
	if err != nil {
		Err := domain.CreateError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, bson.M{"errors": Err})
		return
	}

	id := c.Param("uuid")
	data, Err := h.Service.UpdateProductByIdService(product, id)
	response := domain.CreteResponse(data, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	c.JSON(http.StatusOK, response)
}

func (h handler) DeleteProductById(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValid := domain.IsTokenValid(token)
	if !isValid {
		Err := domain.CreateError(http.StatusUnauthorized, "Access Error")
		c.JSON(http.StatusUnauthorized, bson.M{"errors": Err})
		return
	}

	id := c.Param("uuid")
	Err := h.Service.DeleteProductByIdService(id)
	response := domain.CreteResponse(nil, Err, nil)
	if Err != nil {
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	c.JSON(http.StatusOK, response)
}