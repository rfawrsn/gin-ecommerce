package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	Products  string             `json:"product"`
	Suppliers []string           `json:"suppliers"`
	Time      []int              `json:"time"`
}

func InsertProduct(product Product) error {
	collection := MongoClient.Database(DB).Collection(CollName)
	inserted, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a record with id:", inserted.InsertedID)
	return err
}

func InsertMany(products []Product) error {
	//Convert to a slice of interface{}
	newProducts := make([]interface{}, len(products))
	for i, product := range products {
		newProducts[i] = product
	}

	collection := MongoClient.Database(DB).Collection(CollName)
	result, err := collection.InsertMany(context.TODO(), newProducts)
	if err != nil {
		panic(err)
	}
	log.Println(result)

	return err
}

func UpdateProduct(productId string, product Product) error {
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"product": product.Products, "suppliers": product.Suppliers, "time": product.Time}}

	collection := MongoClient.Database(DB).Collection(CollName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("New record: ", result)
	return nil
}

func DeleteProduct(productId string) error {
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	collection := MongoClient.Database(DB).Collection(CollName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Println("Delete result: ", result)
	return nil
}

func Find(productName string) (Product, error) {
	var result Product
	filter := bson.D{{"product", productName}}
	collection := MongoClient.Database(DB).Collection(CollName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Product{}, err
	}
	return result, nil
}

func FindAll(productName string) []Product {
	var results []Product

	filter := bson.D{{"product", productName}}

	collection := MongoClient.Database(DB).Collection(CollName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}

	return results

}

func ListAll(productName string) []Product {
	var results []Product

	collection := MongoClient.Database(DB).Collection(CollName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), results)
	if err != nil {
		log.Fatal(err)
	}

	return results

}

func DeleteAll() error {
	collection := MongoClient.Database(DB).Collection(CollName)
	delResult, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Records deleted: ", delResult.DeletedCount)

	return err

}

func FindById(productId string) (Product, error) {
	var result Product
	id, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return Product{}, err
	}
	filter := bson.M{"_id": id}
	collection := MongoClient.Database(DB).Collection(CollName)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Product{}, err
	}
	return result, nil
}
