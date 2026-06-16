package init

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	// MongoDbURL ...
	MongoDbURL = "mongodb://localhost:27017"
	// MongoDB ...
	MongoDB = "mydb"
	// MongoUser ...
	MongoUser = "rootuser"
	// MongoPass ...
	MongoPass = "rootpass"
)

// MongoClient ...
func MongoClient() *mongo.Client {
	credential := options.Credential{
		Username:   MongoUser,
		Password:   MongoPass,
		AuthSource: "admin", // База данных, где зарегистрирован пользователь
	}

	clientOptions := options.Client().
		ApplyURI(MongoDbURL).
		SetAuth(credential)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	return client
}
