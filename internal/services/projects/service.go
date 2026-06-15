package projects

import "go.mongodb.org/mongo-driver/mongo"

type service struct {
	db *mongo.Database
}

func NewService(db *mongo.Database) *service {
	return &service{
		db: db,
	}
}
