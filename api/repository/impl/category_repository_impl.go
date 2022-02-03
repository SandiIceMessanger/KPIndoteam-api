package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"chatnews-api/api/model"
	repository "chatnews-api/api/repository"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	mongoDB *mongo.Database
}

func NewCategoryRepository(mongo *mongo.Database) repository.CategoryRepositoryI {
	return &CategoryRepository{
		mongoDB: mongo,
	}
}

func (cr CategoryRepository) GetAllDataCategory(ctx context.Context) (crudResp model.GetDataCategoryResponse, err error) {

	query, err := cr.mongoDB.Collection("categories").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetDataCategoryResponse{}, err
	}
	defer query.Close(ctx)

	listDataCategory := make([]model.DataCategory, 0)
	for query.Next(ctx) {
		var row model.DataCategory
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataCategory = append(listDataCategory, row)
	}

	crudResp = model.GetDataCategoryResponse{Data: listDataCategory}

	return crudResp, err
}

func (cr CategoryRepository) InsertDataCategory(ctx context.Context, req model.DataCategoryRequest) (err error) {
	// var cthAsia, _ = time.LoadLocation("Asia/Bangkok")

	dataReq := bson.M{
		"name":       req.Name,
		"created_at": time.Now(),
	}

	query, err := cr.mongoDB.Collection("categories").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		categoryID := oid.Hex()
		dataUpdateCategoryID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"_id": categoryID}}
		_, err := cr.mongoDB.Collection("category").UpdateOne(ctx, dataUpdateCategoryID, dataObjectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}

	return err
}

func (cr CategoryRepository) UpdateDataCategory(ctx context.Context, req model.DataCategoryRequest) (err error) {

	dataUpdateCategoryID := bson.M{"_id": req.ID}
	dataObjectID := bson.M{"$set": bson.M{
		"name":       req.Name,
		"updated_at": time.Now(),
	}}
	_, err = cr.mongoDB.Collection("categories").UpdateOne(ctx, dataUpdateCategoryID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (cr CategoryRepository) DeleteDataCategory(ctx context.Context, req model.DataCategoryRequest) (err error) {

	dataUpdateCategoryID := bson.M{"_id": req.ID}

	_, err = cr.mongoDB.Collection("categories").DeleteOne(ctx, dataUpdateCategoryID)
	if err != nil {
		log.Println("error")
	}

	return err
}
