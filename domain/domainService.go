package domain

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type domain struct {
	Repo Repository
}

func NewDomainService(repo Repository) Service {
	return &domain{Repo: repo}
}

func (d domain) CreateProductService(product *Product) (*Product, *Error) {
	product.ID = uuid.NewString()
	product.CreatedAt = time.Now().Unix()
	product.Status = 1

	err := d.Repo.CreateProductRepository(product)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return product, nil
}

func (d domain) GetOneProductByIdService(id string) (*Product, *Error) {
	filters := []Filter{
		{Field: "_id", Value: id, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}

	product, err := d.Repo.GetOneProductRepository(filters)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return product, nil
}

func (d domain) GetProductsByFilterService(filters []Filter, pagination *Pagination) ([]Product, int64, *Error) {
	filters = append(filters, Filter{Field: "status", Value: -1, Operator: "$ne"})

	products, count, err := d.Repo.GetProductsRepository(filters, pagination)
	if err != nil {
		return nil, 0, CreateError(http.StatusBadRequest, err.Error())
	}

	return products, count, nil
}

func (d domain) UpdateProductByIdService(product *Product, id string) (*Product, *Error) {
	filters := []Filter{
		{Field: "_id", Value: id, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}

	err := d.Repo.UpdateOneProductRepository(filters, product)
	if err != nil {
		return nil, CreateError(http.StatusBadRequest, err.Error())
	}

	return product, nil
}

func (d domain) DeleteProductByIdService(id string) *Error {
	filters := []Filter{
		{Field: "_id", Value: id, Operator: "$eq"},
		{Field: "status", Value: -1, Operator: "$ne"},
	}

	err := d.Repo.DeleteOneProductRepository(filters)
	if err != nil {
		return CreateError(http.StatusBadRequest, err.Error())
	}

	return nil
}