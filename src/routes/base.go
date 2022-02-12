package routes

import "github.com/gin-gonic/gin"

//ConfigureRouters -> configure route handlers
func ConfigureRouters(router *gin.Engine) {
	router.GET("/v1/tasks", GetAllTasks)
	router.GET("/v1/tasks/:id", GetTaskByID)
	router.POST("/v1/tasks", CreateNewTask)
	router.PUT("/v1/tasks/:id", UpdateTask)
	router.DELETE("/v1/tasks/:id", DeleteTask)
}

//JSONError -> for http 500
func JSONError(err error) map[string]interface{} {
	return gin.H{"message": "unexpected error ocurred", "error": err.Error()}
}

//JSONValidator -> for http 400/401/422/404
func JSONValidator(code string, message string) map[string]interface{} {
	return gin.H{"code": code, "message": message}
}
