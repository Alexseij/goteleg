package goteleg

type FormatingOption string

const (
	MarkdownV2 FormatingOption = "MarkdownV2"
	HTML       FormatingOption = "HTML"
	Markdown   FormatingOption = "Markdown"
)

type DefaultOption struct {
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageID         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	ParseMode                FormatingOption `json:"parse_mode"`
}

type MessageOption struct {
	DisableWebPagePreview bool `json:"disable_web_page_preview"`
	DefaultOption
}

type PhotoOption struct {
	Caption         string           `json:"caption"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	DefaultOption
}
