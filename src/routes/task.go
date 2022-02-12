package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lexgalante/go.gin/src/models"
	"github.com/lexgalante/go.gin/src/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

//GetAllTasks -> retrieve all taks
func GetAllTasks(c *gin.Context) {
	tasks, err := repositories.FilterTasks(bson.D{{}})
	if err != nil {
		log.Print("unexpected error occurred while trying to filter tasks: ", err.Error())
		c.JSON(http.StatusInternalServerError, JSONError(err))
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

//GetTaskByID -> retrieve task by ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, JSONValidator("INVALID_PARAMETER", "parameter [id] invalid"))
		return
	}

	task, err := repositories.FindTask(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, JSONValidator("NOT_FOUND", "task not found"))
			return
		}

		log.Print("unexpected error occurred while trying to filter tasks: ", err.Error())

		c.JSON(http.StatusInternalServerError, JSONError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

//CreateNewTask -> create new task
func CreateNewTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		log.Print("unexpected error occurred while trying to desserialize task: ", err.Error())
		c.JSON(http.StatusBadRequest, JSONValidator("INVALID_PAYLOAD", fmt.Sprintf("check payload request: %s", err.Error())))
		return
	}

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err = repositories.InsertTask(&task)
	if err != nil {
		log.Print("unexpected error occurred while trying to insert task: ", err.Error())
		c.JSON(http.StatusInternalServerError, JSONError(err))
		return
	}

	c.JSON(http.StatusCreated, task)
}

//UpdateTask -> update task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, JSONValidator("INVALID_PARAMETER", "parameter [id] invalid"))
		return
	}

	var task models.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		log.Print("unexpected error occurred while trying to desserialize task: ", err.Error())
		c.JSON(http.StatusBadRequest, JSONValidator("INVALID_PAYLOAD", fmt.Sprintf("check payload request: %s", err.Error())))
		return
	}

	task.UpdatedAt = time.Now()

	err = repositories.ChangeTask(id, &task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, JSONValidator("NOT_FOUND", "task not found"))
			return
		}

		log.Print("unexpected error occurred while trying to update task: ", id, " error: ", err.Error())
		c.JSON(http.StatusInternalServerError, JSONError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

//DeleteTask -> delete task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, JSONValidator("INVALID_PARAMETER", "parameter [id] invalid"))
		return
	}

	err := repositories.RemoveTask(id)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, JSONValidator("NOT_FOUND", "task not found"))
			return
		}

		log.Print("unexpected error occurred while trying to update task: ", id, " error: ", err.Error())
		c.JSON(http.StatusInternalServerError, JSONError(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
