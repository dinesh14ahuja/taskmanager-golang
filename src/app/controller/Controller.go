package controller

import (
	"boilerplate-webserver-code/src/app/model"
	"boilerplate-webserver-code/src/constant"
	"boilerplate-webserver-code/src/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewTask(c *gin.Context) {
	var taskData = model.Task{}
	taskData.New()
	err := c.BindJSON(&taskData)
	//body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.FatalLn("Failed to read request body." + err.Error())
	}
	constant.Tasklist[taskData.TaskNumber] = taskData
	c.JSON(200, gin.H{
		"message": "Task saved successfully",
	})
}

func GetTasks(c *gin.Context) {
	var response = make([]model.Task, 0)
	for _, value := range constant.Tasklist {
		response = append(response, value)
	}
	c.JSON(200, response)
}

func DeleteTask(c *gin.Context) {
	tasknumber := c.Param("tasknumber")
	for key := range constant.Tasklist {
		tasknumberValue, _ := strconv.Atoi(tasknumber)
		if key == tasknumberValue {
			delete(constant.Tasklist, key)
			c.JSON(200, gin.H{
				"message": "Task deleted successfully",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message": "Task not found",
	})

}
