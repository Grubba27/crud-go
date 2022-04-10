package pet

import (
	database "crud-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePetContext() *mongo.Collection {
	return database.DB.Collection("Pets")
}
