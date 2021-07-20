package goteleg

type FormatingOption string

const (
	MarkdownV2 FormatingOption = "MarkdownV2"
	HTML       FormatingOption = "HTML"
	Markdown   FormatingOption = "Markdown"
)

type DefaultOption struct {
	DisableNotification      bool
	ReplyToMessageID         int
	AllowSendingWithoutReply bool
}

type MessageOption struct {
	ParseMode             FormatingOption
	DisableWebPagePreview bool
	DefaultOption
}

type PhotoOption struct {
	Caption         string
	ParseMode       FormatingOption
	CaptionEntities []*MessageEntity
	DefaultOption
}
