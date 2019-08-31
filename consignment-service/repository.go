package main

import (
	"context"

	pb "github.com/jobin212/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignemnt); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
