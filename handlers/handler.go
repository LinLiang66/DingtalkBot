package handlers

import (
	"DingtalkBot/model"
	"context"
	"fmt"
)

// 责任链
func chain(data *ActionInfo, actions ...Action) bool {
	for _, v := range actions {
		if !v.Execute(data) {
			return false
		}
	}
	return true
}

type MessageHandler struct {
}

func judgeMsgType(event *model.DingtalkMessage) (string, error) {
	msgType := event.Msgtype
	switch msgType {
	case "text", "picture", "richText":
		return msgType, nil
	default:
		return "", fmt.Errorf("unknown message type: %v", msgType)
	}

}

func (m MessageHandler) msgReceivedHandler(ctx context.Context, event *model.DingtalkMessage) error {
	handlerType := judgeChatType(event)
	if handlerType == "otherChat" {
		return nil
	}
	msgType, err := judgeMsgType(event)
	if err != nil {
		fmt.Printf("error getting message type: %v\n", err)
		return nil
	}

	qParsed := event.Text.Content
	msgId := event.MsgID
	chatId := event.RobotCode
	sessionId := event.SenderID

	msgInfo := MsgInfo{
		msgType:        msgType,
		MsgId:          msgId,
		UserId:         sessionId,
		chatId:         chatId,
		QParsed:        qParsed,
		sessionId:      sessionId,
		Appid:          event.ChatbotCorpID,
		SessionWebhook: event.SessionWebhook,
	}
	data := &ActionInfo{
		Ctx:     &ctx,
		handler: &m,
		Info:    &msgInfo,
	}
	actions := []Action{
		&ProcessedUniqueAction{}, //避免重复处理
		&EmptyAction{},           //空消息处理
		&HelpAction{},            //帮助处理
	}
	chain(data, actions...)

	return nil
}
