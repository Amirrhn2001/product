package domain

type Service interface {
	CreateProductService(product *Product) (*Product, *Error)
	GetOneProductByIdService(id string) (*Product, *Error)
	GetProductsByFilterService(filters []Filter, pagination *Pagination) ([]Product, int64, *Error)
	UpdateProductByIdService(product *Product, id string) (*Product, *Error)
	DeleteProductByIdService(id string) *Error
}