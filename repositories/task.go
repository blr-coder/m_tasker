package repositories

import (
	"context"
	"github.com/blr-coder/m_tasker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskClient struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func (c *TaskClient) Create(task models.Task) (models.Task, error) {
	result, err := c.Collection.InsertOne(c.Ctx, task)
	if err != nil {
		return models.Task{}, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *TaskClient) Get(id string) (models.Task, error) {
	task := models.Task{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	err = c.Collection.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (c *TaskClient) Delete(id string) (models.Task, error) {
	panic("implement me")
}
