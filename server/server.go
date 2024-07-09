package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"task4/db"
	"task4/routes"
	"task4/services"
	"task4/web"
)

func main() {
	// Connect to the MongoDB database
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Create a MongoDB instance for the database operations
	mongoDB := db.NewMongo(client)

	// Enable CORS (Cross-Origin Resource Sharing) only in the production profile
	cors := os.Getenv("profile") == "prod"

	// Create a UserRouter instance with a UserService and App instance
	userRouter := routes.NewUserRouter(services.NewUserService(mongoDB))

	// Create and start the web application
	app := web.NewApp(cors, userRouter)
	err = app.Serve()
	log.Println("Error", err)
}

// clientOptions returns the MongoDB client options based on the profile.
func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
