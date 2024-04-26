package controller

import (
	"DingtalkBot/handlers"
	"DingtalkBot/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WebhookController struct {
}

func (controller *WebhookController) EventHandlerFunc(c *gin.Context) {
	appid := c.Param("appid")
	if len(appid) == 0 {
		c.JSON(200, gin.H{
			"message":   "appid Cannot be empty!!",
			"code":      -1,
			"success":   false,
			"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
		})
		return
	}

	var event *model.DingtalkMessage
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
			"code":      500,
			"success":   false,
			"msg":       "参数错误",
		})
		return
	}
	if appid != *event.RobotCode {
		c.JSON(200, gin.H{
			"message":   "appid is invalid",
			"code":      -1,
			"success":   false,
			"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "success",
		"code":      200,
		"success":   true,
		"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
	})
	err := handlers.Handler(c, event)
	if err != nil {
		return
	}
	return

}
