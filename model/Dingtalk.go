package model

type AtUser struct {
	DingtalkId string `json:"dingtalkId"`
	StaffId    string `json:"staffId"`
}

type RichTextNode struct {
	Text         string `json:"text,omitempty"`
	DownloadCode string `json:"downloadCode,omitempty"`
	Type         string `json:"type,omitempty"`
}

type Content struct {
	DownloadCode string         `json:"downloadCode"`
	FileName     string         `json:"fileName"`
	Duration     int64          `json:"duration"`
	VideoType    string         `json:"videoType"`
	Recognition  string         `json:"recognition"`
	RichText     []RichTextNode `json:"richText"`
}

type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DingtalkMessage struct {
	SenderPlatform            string      `json:"senderPlatform"`
	ConversationID            string      `json:"conversationId"`
	AtUsers                   []AtUser    `json:"atUsers"`
	ChatbotCorpID             string      `json:"chatbotCorpId"`
	ChatbotUserID             string      `json:"chatbotUserId"`
	MsgID                     *string     `json:"msgId"`
	SenderNick                string      `json:"senderNick"`
	IsAdmin                   bool        `json:"isAdmin"`
	SenderStaffID             string      `json:"senderStaffId"`
	SessionWebhookExpiredTime int64       `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64       `json:"createAt"`
	SenderCorpID              string      `json:"senderCorpId"`
	ConversationType          string      `json:"conversationType"`
	SenderID                  *string     `json:"senderId"`
	ConversationTitle         string      `json:"conversationTitle"`
	IsInAtList                bool        `json:"isInAtList"`
	SessionWebhook            string      `json:"sessionWebhook"`
	Text                      TextContent `json:"text"`
	RobotCode                 *string     `json:"robotCode"`
	Msgtype                   string      `json:"msgtype"`
	Content                   Content     `json:"content"`
}

type TextContent struct {
	Content string `json:"content"`
}

type DingTalkActionCardButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type DingTalkActionCard struct {
	Title          string                     `json:"title"`
	Text           string                     `json:"text"`
	BtnOrientation string                     `json:"btnOrientation"`
	SingleUrl      string                     `json:"singleUrl"`
	Btns           []DingTalkActionCardButton `json:"btns"`
}

type DingTalkCardMessage struct {
	MsgType    string             `json:"msgtype"`
	ActionCard DingTalkActionCard `json:"actionCard"`
}
