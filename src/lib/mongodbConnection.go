package lib

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type ShortUrl struct {
	Url  string
	Hash string
}

func Init() *mongo.Client {
	// TODO fix url .env
	err := godotenv.Load("/Users/andrea/go/src/urlShortner/src/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongodb_connection_url := os.Getenv("MONGODB_CONNECTION_URL")
	fmt.Println(mongodb_connection_url)

	clientOptions := options.Client().ApplyURI(mongodb_connection_url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// fmt.Printf("%T\n", client)

	return client
}

func GetCollection(client *mongo.Client, db_name string, collection_name string) *mongo.Collection {
	collection := client.Database(db_name).Collection(collection_name)
	// fmt.Printf("%T\n", collection)
	return collection

}

func EndDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}

func InsertURL(collection *mongo.Collection, url string) string {
	result := MakeShortUrl(url)
	insertResult, err := collection.InsertOne(context.TODO(), result)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return result.Hash
}

func QueryURL(collection *mongo.Collection, hash_str string) string {
	var result ShortUrl
	filter := bson.D{{"hash", hash_str}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return result.Url

}

func W_GetCollection() (*mongo.Client, *mongo.Collection) {
	client := Init()
	mongodb_dbname_shorturl := os.Getenv("MONGODB_DBNAME_SHORTURL")
	mongodb_collection_shorturldata := os.Getenv("MONGODB_COLLECTION_SHORTURLDATA")
	collection := GetCollection(client, mongodb_dbname_shorturl, mongodb_collection_shorturldata)
	return client, collection
}
