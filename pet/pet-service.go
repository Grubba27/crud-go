package pet

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Service() T {
	db := CreatePetContext()

	service := T{
		getById: func(id string) Pet {
			var result Pet
			err := db.FindOne(
				context.TODO(),
				bson.D{{"_id", id}},
			).Decode(&result)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return Pet{}
				}
				log.Fatal(err)
			}
			return result
		},
		getByIds: func(ids []string) []Pet {
			var result []Pet
			err := db.FindOne(
				context.TODO(),
				bson.D{{"_id", ids}},
			).Decode(&result)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return []Pet{}
				}
				log.Fatal(err)
			}
			return result
		},
		deleteById: func(id string) *mongo.DeleteResult {
			res, err := db.DeleteOne(
				context.TODO(),
				bson.D{{"_id", id}},
			)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return nil
				}
				log.Fatal(err)
			}
			return res
		},
		deleteByIds: func(ids []string) *mongo.DeleteResult {
			res, err := db.DeleteMany(
				context.TODO(),
				bson.D{{"_id", ids}},
			)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return nil
				}
				log.Fatal(err)
			}
			return res
		},
		updatePet: func(id string, pet Pet) Pet {
			var result Pet
			err := db.FindOneAndReplace(
				context.TODO(),
				bson.D{{"_id", id}},
				pet,
				options.FindOneAndReplace().SetUpsert(true),
			).Decode(&result)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					return Pet{}
				}
				log.Fatal(err)
			}
			return result
		},
		createPet: func(pet Pet) Pet {
			res, err := db.InsertOne(context.TODO(), pet)
			if err != nil {
				log.Fatal(err)
			}
			pet._id = res.InsertedID
			return pet
		},
	}

	return service
}

type T struct {
	getById     func(id string) Pet
	getByIds    func(ids []string) []Pet
	deleteById  func(id string) *mongo.DeleteResult
	deleteByIds func(ids []string) *mongo.DeleteResult
	updatePet   func(id string, pet Pet) Pet
	createPet   func(pet Pet) Pet
}
