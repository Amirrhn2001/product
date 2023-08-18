package domain

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

func IsTokenValid(t string) bool {
	publicKey, err := os.ReadFile(os.Getenv("publicKey"))
	if err != nil {
		fmt.Println("Err to read public key", err)
		return false
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		fmt.Println("Err to parse public key", err)
		return false
	}

	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("Err to parse token", err)
		return false
	}

	fmt.Println(token.Valid)

	return token.Valid
}

func CreateError(status int, message string) *Error {
	err := &Error{
		Code:    status,
		Message: message,
		Details: "",
	}
	return err
}

func CreteResponse(data any, err *Error, meta map[string]any) *Response {
	return &Response{
		Error: err,
		Data:  data,
		Meta:  meta,
	}
}

func CreteMongoFilter(filters []Filter) (map[string]any, error) {
	if len(filters) == 0 {
		return nil, fmt.Errorf("Error empty filter")
	}
	var mongoFilter = make(map[string]any, len(filters))
	for _, filter := range filters {
		switch filter.Operator {
		case "$eq":
			mongoFilter[filter.Field] = bson.M{"$eq": filter.Value}
		case "$ne":
			mongoFilter[filter.Field] = bson.M{"$ne": filter.Value}
		case "$gt":
			mongoFilter[filter.Field] = bson.M{"$gt": filter.Value}
		case "$gte":
			mongoFilter[filter.Field] = bson.M{"$gte": filter.Value}
		case "$lt":
			mongoFilter[filter.Field] = bson.M{"$lt": filter.Value}
		case "$lte":
			mongoFilter[filter.Field] = bson.M{"$lte": filter.Value}
		case "$in":
			mongoFilter[filter.Field] = bson.M{"$in": filter.Value}
		case "$nin":
			mongoFilter[filter.Field] = bson.M{"$nin": filter.Value}
		case "$exists":
			mongoFilter[filter.Field] = bson.M{"$exists": filter.Value}
		case "$type":
			mongoFilter[filter.Field] = bson.M{"$type": filter.Value}
		}
	}

	return mongoFilter, nil
}
