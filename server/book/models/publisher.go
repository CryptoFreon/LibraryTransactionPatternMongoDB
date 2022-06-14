package models

import (
	"context"
	"github.com/Ja7ad/library/server/book/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colPub = global.Client.GetDatabase("library").Collection("publisher")

type Publisher struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}

func (b *Publisher) Insert(ctx context.Context) error {
	if _, err := colPub.InsertOne(ctx, b); err != nil {
		return err
	}
	return nil
}

func GetPublisherByName(ctx context.Context, publisherName string) (*Publisher, error) {
	publisher := &Publisher{}
	if err := colPub.FindOne(ctx, bson.M{"name": publisherName}).Decode(publisher); err != nil {
		return nil, err
	}
	return publisher, nil
}
