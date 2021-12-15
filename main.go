package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilhamabdlh/iteration1/httpserver"
	paginate "github.com/ilhamabdlh/iteration1/models"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
)


func insertExamples(db *mongo.Database) (insertedIds []interface{}, err error) {
	var data []interface{}
	for i := 0; i < 30; i++ {
		data = append(data, bson.M{
			"name":     fmt.Sprintf("product-%d", i),
			"quantity": float64(i),
			"price":    float64(i*10 + 5),
		})
	}
	result, err := db.Collection("products").InsertMany(
		context.Background(), data)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}


func main() {

	http.HandleFunc("/normal-pagination", func(w http.ResponseWriter, r *http.Request) {
		convertedPageInt, convertedLimitInt := getPageAndLimit(r)
		// Example for Normal Find query
		filter := bson.M{}
		limit := int64(convertedLimitInt)
		page := int64(convertedPageInt)
		db := httpserver.Connect()
		collection := db.Collection("products")
		projection := bson.D{
			{"name", 1},
			{"quantity", 1},
		}
		var products []models.Product
		paginatedData, err := paginate.New(collection).Context(ctx).Limit(limit).Page(page).Sort("price", -1).Sort("quantity", -1).Select(projection).Filter(filter).Decode(&products).Find()
		if err != nil {
			panic(err)
		}

		payload := struct {
			Data       []models.Product               `json:"data"`
			Pagination paginate.PaginationData `json:"pagination"`
		}{
			Pagination: paginatedData.Pagination,
			Data:       products,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payload)
	})

	_, insertErr := insertExamples(client.Database("atop"))
	if insertErr != nil {
		panic(insertErr)
	}

	fmt.Println("Application started on port http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getPageAndLimit(r *http.Request) (convertedPageInt int, convertedLimitInt int) {
	queryPageValue := r.FormValue("page")
	if queryPageValue != "" {
		convertedPageInt, _ = strconv.Atoi(queryPageValue)
	}

	queryLimitValue := r.FormValue("limit")
	if queryLimitValue != "" {
		convertedLimitInt, _ = strconv.Atoi(queryLimitValue)
	}

	return convertedPageInt, convertedLimitInt
}