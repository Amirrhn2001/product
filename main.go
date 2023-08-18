package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"product/delivery/api"
	"product/domain"
	"product/repository/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnvFile()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	repository, err := mongodb.NewMongoRepository()
	if err != nil {
		log.Fatal("Error connecting repository:", err)
	}
	service := domain.NewDomainService(repository)
	handler := api.NewHandler(service)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// product CRUD
	r.POST("/api/product", handler.CreateProduct)
	r.GET("/api/product/uuid/:uuid", handler.GetOneProductById)
	r.GET("/api/products", handler.GetProductsByFilter)
	r.PUT("/api/product/uuid/:uuid", handler.UpdateProductById)
	r.DELETE("/api/product/uuid/:uuid", handler.DeleteProductById)

	Err := make(chan error, 1)
	go func() {
		fmt.Println("Listening on port", choosePort())
		Err <- http.ListenAndServe(choosePort(), r)
	}()

	fmt.Printf("Terminated %s", <-Err)
}

func loadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func choosePort() string {
	port := os.Getenv("port")
	if port == "" {
		return ":8080"
	}
	return fmt.Sprintf(":%s", port)
}