package domain

type Repository interface {
	CreateProductRepository(product *Product) error
	GetOneProductRepository(filters []Filter) (*Product, error)
	GetProductsRepository(filters []Filter, pagination *Pagination) ([]Product, int64, error)
	UpdateOneProductRepository(filters []Filter, product *Product) error
	DeleteOneProductRepository(filters []Filter) error
}