package repositories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetDbCollection -> mongo collection
func GetDbCollection(dbContext context.Context, database string, collection string) (*mongo.Collection, error) {
	dbClient, err := mongo.Connect(dbContext, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Printf("unexpected error ocurred while trying to connect database: %s, error: %s", os.Getenv("DB_URI"), err.Error())
		return nil, err
	}

	dbCollection := dbClient.Database(database).Collection(collection)

	return dbCollection, nil
}
