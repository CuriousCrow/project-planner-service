package init

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "mongodb://localhost:27017"
	MONGO_DB    = "mydb"
	MONGO_USER  = "rootuser"
	MONGO_PASS  = "rootpass"
)

func MongoClient(ctx context.Context) *mongo.Client {
	credential := options.Credential{
		Username:   MONGO_USER,
		Password:   MONGO_PASS,
		AuthSource: "admin", // База данных, где зарегистрирован пользователь
	}

	clientOptions := options.Client().
		ApplyURI(MONGODB_URL).
		SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}

	return client
}

func MongoDB(ctx context.Context) *mongo.Database {
	credential := options.Credential{
		Username:   MONGO_USER,
		Password:   MONGO_PASS,
		AuthSource: "admin", // База данных, где зарегистрирован пользователь
	}

	clientOptions := options.Client().
		ApplyURI(MONGODB_URL).
		SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Ошибка подключения к MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Ошибка при отключении: %v", err)
		}
	}()

	return client.Database(MONGO_DB)
}
