package handlers

import (
	"DingtalkBot/model"
	"context"
)

type MessageHandlerInterface interface {
	msgReceivedHandler(ctx context.Context, event *model.DingtalkMessage) error
}

type HandlerType string

const (
	GroupHandler = "group"
	UserHandler  = "personal"
)

// handlers 所有消息类型类型的处理器
var handlers MessageHandlerInterface

func InitHandlers() {
	handlers = NewMessageHandler()
}

func Handler(ctx context.Context, event *model.DingtalkMessage) error {
	return handlers.msgReceivedHandler(ctx, event)
}

func NewMessageHandler() MessageHandlerInterface {
	return &MessageHandler{}
}

func judgeChatType(event *model.DingtalkMessage) HandlerType {
	chatType := event.ConversationType
	if chatType == "1" {
		return UserHandler
	}
	if chatType == "2" {
		return GroupHandler
	}
	return "otherChat"
}
