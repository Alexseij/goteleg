package goteleg

type FormatingOption string

const (
	MarkdownV2 FormatingOption = "MarkdownV2"
	HTML                       = "HTML"
	Markdown                   = "Markdown"
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
