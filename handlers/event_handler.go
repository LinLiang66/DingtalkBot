package handlers

import (
	"DingtalkBot/utils"
	"context"
)

type MsgInfo struct {
	handlerType    HandlerType
	msgType        string
	MsgId          *string
	chatId         *string
	UserId         *string
	QParsed        string
	imageKey       []string
	sessionId      *string
	Appid          string
	SessionWebhook string
}
type ActionInfo struct {
	handler *MessageHandler
	Ctx     *context.Context
	Info    *MsgInfo
}

type Action interface {
	Execute(a *ActionInfo) bool
}

// ProcessedUniqueAction 消息去重处理
type ProcessedUniqueAction struct { //幂等判断消息唯一性
}

func (*ProcessedUniqueAction) Execute(a *ActionInfo) bool {
	if utils.RedisClient.KEYEXISTS(*a.Ctx, "robot:message_event:"+*a.Info.MsgId) {
		return false
	}
	err := utils.RedisClient.SetStrWithExpire(*a.Ctx, "robot:message_event:"+*a.Info.MsgId, "Message has been handle", 25200)
	if err != nil {
		return false
	}
	return true
}

// EmptyAction 空内容处理
type EmptyAction struct { /*空消息*/
}

func (*EmptyAction) Execute(a *ActionInfo) bool {
	if len(a.Info.QParsed) == 0 {
		//sendMsg(*a.Ctx, "🤖️：您好，请问有什么可以帮到您~", a.Info.chatId, a.Info.Appid)
		return false
	}
	return true
}

// HelpAction 帮助
type HelpAction struct { /*帮助*/
}

func (*HelpAction) Execute(a *ActionInfo) bool {
	if foundClear, _ := utils.ContainsSpecificContent(a.Info.QParsed,
		"help|帮助"); foundClear {
		sendHelpCard(*a.Ctx, a.Info)
		return false
	}
	return true
}
