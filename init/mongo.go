package init

import (
	"log"

	"github.com/CuriousCrow/project-planner-service/configs"
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
func MongoClient(appConfig *configs.AppConfig) *mongo.Client {

	credential := options.Credential{
		Username:   appConfig.Mongo.Username,
		Password:   appConfig.Mongo.Password,
		AuthSource: "admin", // База данных, где зарегистрирован пользователь
	}

	clientOptions := options.Client().
		ApplyURI(appConfig.Mongo.Url).
		SetAuth(credential)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	return client
}
