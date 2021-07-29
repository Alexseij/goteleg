package goteleg

// import (
// 	"net/url"
// 	"strconv"
// )

// func setDefaultOptinons(option DefaultOption, val *url.Values) {
// 	if option.AllowSendingWithoutReply {
// 		val.Add("allow_sending_without_reply", "true")
// 	}
// 	if option.DisableNotification {
// 		val.Add("disable_notification", "true")
// 	}
// 	if option.ReplyToMessageID != 0 {
// 		val.Add("reply_to_message_id", strconv.Itoa(option.ReplyToMessageID))
// 	}

// }

// func createQuery(options interface{}) url.Values {
// 	val := url.Values{}

// 	switch options.(type) {
// 	case *MessageOption:
// 		msgOption := options.(*MessageOption)
// 		if msgOption == nil {
// 			return val
// 		}

// 		setDefaultOptinons(msgOption.DefaultOption, &val)

// 		if msgOption.DisableWebPagePreview {
// 			val.Add("disable_web_page_preview", "true")
// 		}
// 		val.Add("parse_mode", string(msgOption.ParseMode))

// 	case *PhotoOption:
// 		photoOption := options.(*PhotoOption)
// 		if photoOption == nil {
// 			return val
// 		}

// 		setDefaultOptinons(photoOption.DefaultOption, &val)

// 		if photoOption.Caption != "" {
// 			val.Add("caption", photoOption.Caption)
// 		}

// 		if photoOption.CaptionEntities != nil {
// 			val.Add("caption_entities", photoOption.CaptionEntities)
// 		}

// 		val.Add("parse_mode", string(photoOption.ParseMode))
// 	}
// 	return val
// }

// func parseOptions(options interface{}) string {
// 	return createQuery(options).Encode()
// }
