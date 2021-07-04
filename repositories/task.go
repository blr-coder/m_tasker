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

func (c *TaskClient) Get(id string) (task models.Task, err error) {
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

func (c *TaskClient) List() (tasks models.Tasks, err error) {
	cursor, err := c.Collection.Find(c.Ctx, bson.M{})
	if err != nil {
		return tasks, err
	}

	for cursor.Next(c.Ctx) {
		task := models.Task{}
		_ = cursor.Decode(&task)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (c *TaskClient) Search(filter *models.TaskFilter) (tasks models.Tasks, err error) {
	cursor, err := c.Collection.Aggregate(c.Ctx, c.filterToPipeline(filter))
	if err != nil {
		return tasks, err
	}

	for cursor.Next(c.Ctx) {
		task := models.Task{}
		err = cursor.Decode(&task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (c *TaskClient) Done(id string) (models.Task, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: _id}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "done", Value: true},
	}}}

	_, err = c.Collection.UpdateOne(c.Ctx, filter, updater)
	if err != nil {
		return models.Task{}, err
	}

	return c.Get(id)
}

func (c *TaskClient) Update(id string, task *models.Task) (models.Task, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}

	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"done":        task.Done,
		},
	}

	err = c.Collection.FindOneAndUpdate(c.Ctx, filter, update).Err()
	if err != nil {
		return models.Task{}, err
	}

	return c.Get(id)
}

func (c *TaskClient) Delete(id string) (models.TaskDelete, error) {
	deleted := models.TaskDelete{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return deleted, err
	}

	res, err := c.Collection.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return deleted, err
	}

	deleted.DeletedCount = res.DeletedCount

	return deleted, nil
}

func (c *TaskClient) filterToPipeline(filter *models.TaskFilter) (pipeline mongo.Pipeline) {
	if filter.Title != "" {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{"title": filter.Title}}})
	}
	if filter.Description != "" {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{"description": filter.Description}}})
	}
	if filter.Done != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{"done": filter.Done}}})
	}

	return pipeline
}
