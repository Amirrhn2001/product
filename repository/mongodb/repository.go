package mongodb

import (
	"context"
	"fmt"
	"os"
	"product/domain"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewMongoRepository() (domain.Repository, error) {
	url := os.Getenv("mongoURL")
	timeout, err := strconv.Atoi(os.Getenv("timeout"))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	repo := &mongoRepository{
		client: client,
		database: os.Getenv("database"),
		timeout: time.Duration(timeout) * time.Second,
	}
	return repo, nil
}

func (r mongoRepository) CreateProductRepository(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.ProductsCollection)
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (r mongoRepository) GetOneProductRepository(filters []domain.Filter) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.ProductsCollection)
	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return nil, err
	}

	product := &domain.Product{}
	err = collection.FindOne(ctx, filter).Decode(product)
	if err != nil {
		return nil, err
	}
	if product == (&domain.Product{}) {
		return nil, fmt.Errorf("Not found")
	}

	return product, nil
}

func (r mongoRepository) GetProductsRepository(filters []domain.Filter, pagination *domain.Pagination) ([]domain.Product, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.ProductsCollection)

	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return nil, 0, err
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	options := options.Find()
	options.SetLimit(pagination.Limit)
	options.SetSkip((pagination.Skip-1)*pagination.Limit)

	results, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var products []domain.Product
	err = results.All(ctx, &products)
	if err != nil {
		return nil, 0, err
	}

	err = results.Close(ctx)
	if err != nil {
		return nil, 0, err
	}

	if len(products) == 0 {
		return nil, 0, fmt.Errorf("Not found")
	}

	return products, count, nil
}

func (r mongoRepository) UpdateOneProductRepository(filters []domain.Filter, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.ProductsCollection)

	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return err
	}

	return nil
}

func (r mongoRepository) DeleteOneProductRepository(filters []domain.Filter) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(domain.ProductsCollection)

	filter, err := domain.CreteMongoFilter(filters)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"status": -1}})
	if err != nil {
		return err
	}

	return nil
}