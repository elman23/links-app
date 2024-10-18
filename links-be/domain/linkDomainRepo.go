package domain

import (
	"fmt"
	"links-be/db"
	"links-be/dto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ILinkRepo interface {
	CreateLink(dto.Link) (*dto.Link, error)
	FindLink(string) (*dto.Link, error)
	DeleteLink(string) (*string, error)
	FindLinks() ([]dto.Link, error)
}

type LinkRepoMongo struct {
}

func (lrm LinkRepoMongo) CreateLink(link dto.Link) (*dto.Link, error) {

	collection, client, context, cancel := db.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	link.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context, link)
	if err != nil {
		fmt.Printf("Error in writing object %s", err)
		return nil, err
	}

	fmt.Printf("Inserted link %s", result.InsertedID)

	return &link, nil
}

func (lrm LinkRepoMongo) FindLink(id string) (*dto.Link, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := db.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	l := dto.Link{}
	filter := bson.D{{Key: "_id", Value: oid}}

	err = collection.FindOne(context, filter).Decode(&l)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No document found")
		return nil, err
	} else if err != nil {
		fmt.Printf("Error in mongo %v", err)
		return nil, err
	}

	return &l, nil
}

func (lrm LinkRepoMongo) FindLinks() ([]dto.Link, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := db.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	l := make([]dto.Link, 0, 10)
	filter := bson.D{}

	cursor, err := collection.Find(context, filter)
	defer cursor.Close(context)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No document found")
		return nil, err
	} else if err != nil {
		fmt.Printf("Error in mongo %v", err)
		return nil, err
	}

	for cursor.Next(context) {
		var result dto.Link
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		l = append(l, result)
	}

	return l, nil
}

func (lrm LinkRepoMongo) DeleteLink(id string) (*string, error) {
	fmt.Println("Connection setup is called")
	collection, client, context, cancel := db.SetupMongoDB()

	defer closeConnection(client, context, cancel)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	filter := bson.D{{Key: "_id", Value: oid}}

	result, err := collection.DeleteOne(context, filter)
	if err != nil {
		fmt.Printf("Error with delete %v", err)
	}

	resultStr := fmt.Sprintf("Results deleted %v, id of the item %v", result.DeletedCount, oid)
	return &resultStr, nil
}

func NewLinkRepoMongo() ILinkRepo {
	return LinkRepoMongo{}
}
