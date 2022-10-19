package car_model

import (
	"context"

	"github.com/jhony/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarModel interface {
	GetAll(ctx context.Context) ([]entities.TCar, error)
}

type carModel struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) TCarModel {
	return &carModel{
		collection,
	}
}