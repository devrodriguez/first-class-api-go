package server

import (
	"context"
	"time"

	"github.com/devrodriguez/first-class-api-go/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() *gin.Engine {

	// New server
	server := gin.New()

	// Create db connection
	dbcli := dbConnect()

	// Serve static files
	server.Static("/s", "./storage")

	// Add middlewares
	server.Use(gin.Recovery(), middlewares.Logger())

	// Map URLs
	MapURLs(server, dbcli)

	return server
}

func dbConnect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/locateme?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	// Check connections
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	return client
}
