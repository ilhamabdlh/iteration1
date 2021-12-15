package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
)


// func Run(ctx, context.context){
// 	r:= mux.NewRouter()
// 	c:= cors.New(cors.options{
// 		AllowedOrigins: []string{"*"},
// 		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
// 		AllowedMethods: []string{"GET, POST, PUT, OPTIONS"},
// 		AllowCredentials: true,
// 	})
// 	handler := c.Handler(r)

// 	srv := &http.Server{
// 		Address: "3001",
// 		WriteTimeout: time.Second * 5,
// 		ReadTimeout: time.Second * 5,
// 		Handler: handler,
// 	}
// }



func Connect() (*mongo.Database, error) {
    clientOptions := options.Client()
    clientOptions.ApplyURI("mongodb://localhost:27017")
    client, err := mongo.NewClient(clientOptions)
	ctx := context.Background()
    if err != nil {
        return nil, err
    }
    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }
	collection := client.Database("atop")
    return collection, nil
}


