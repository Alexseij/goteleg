package goteleg

import (
	"net/url"
	"strconv"
)

func createQuery(options interface{}) url.Values {
	val := url.Values{}

	switch options.(type) {
	case *MessageOption:
		msgOption := options.(*MessageOption)
		if msgOption == nil {
			return val
		}
		if msgOption.AllowSendingWithoutReply {
			val.Add("allow_sending_without_reply", "true")
		}
		if msgOption.DisableNotification {
			val.Add("disable_notification", "true")
		}
		if msgOption.DisableWebPagePreview {
			val.Add("disable_web_page_preview", "true")
		}
		if msgOption.ReplyToMessageID != 0 {
			val.Add("reply_to_message_id", strconv.Itoa(msgOption.ReplyToMessageID))
		}
		val.Add("parse_mode", string(msgOption.ParseMode))
	}
	return val
}

func parseOptions(options interface{}) string {
	return createQuery(options).Encode()
}
