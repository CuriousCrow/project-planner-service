package projects

import "go.mongodb.org/mongo-driver/v2/mongo"

type service struct {
	db *mongo.Database
}

// NewService ...
func NewService(db *mongo.Database) *service {
	return &service{
		db: db,
	}
}
