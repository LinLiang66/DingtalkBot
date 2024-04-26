package routers

import (
	"DingtalkBot/controller"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
}

// 用户路由
func routerUser(engine *gin.Engine) {
	con := &controller.WebhookController{}
	// 添加新的路由
	webhookGroup := engine.Group("/webhook")
	{
		webhookGroup.POST("/event/:appid", con.EventHandlerFunc) //钉钉机器人消息事件处理
	}

}
