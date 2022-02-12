package repositories

import (
	"context"
	"log"
	"time"

	"github.com/lexgalante/go.gin/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FilterTasks -> filter collection tasks
func FilterTasks(filter interface{}) ([]*models.Task, error) {
	var tasks []*models.Task

	dbContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbCollection, err := GetDbCollection(dbContext, "tasks", "tasks")
	if err != nil {
		return tasks, err
	}

	dbCursor, err := dbCollection.Find(dbContext, filter)
	if err != nil {
		return tasks, err
	}
	defer dbCursor.Close(dbContext)

	for dbCursor.Next(dbContext) {
		var task models.Task
		err := dbCursor.Decode(&task)
		if err != nil {
			return tasks, nil
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

//FindTask -> find task by id
func FindTask(id string) (*models.Task, error) {
	var task *models.Task

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	dbContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbCollection, err := GetDbCollection(dbContext, "tasks", "tasks")
	if err != nil {
		return task, err
	}

	err = dbCollection.FindOne(dbContext, bson.M{"_id": objectID}).Decode(&task)

	return task, err
}

//InsertTask -> create new task
func InsertTask(task *models.Task) error {
	dbContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbCollection, err := GetDbCollection(dbContext, "tasks", "tasks")
	if err != nil {
		return err
	}

	result, err := dbCollection.InsertOne(dbContext, task)
	if err != nil {
		return err
	}

	log.Print("new task add: ", result.InsertedID)

	return nil
}

//ChangeTask -> update task
func ChangeTask(id string, task *models.Task) error {
	dbContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbCollection, err := GetDbCollection(dbContext, "tasks", "tasks")
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	task.ID = objectID

	result, err := dbCollection.UpdateByID(dbContext, objectID, bson.M{"$set": task})
	if err != nil {
		return err
	}

	log.Print(result.ModifiedCount, " tasks updated")

	return nil
}

//RemoveTask -> delete task
func RemoveTask(id string) error {
	dbContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbCollection, err := GetDbCollection(dbContext, "tasks", "tasks")
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := dbCollection.DeleteOne(dbContext, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	log.Print(result.DeletedCount, " tasks deleted")

	return nil
}
