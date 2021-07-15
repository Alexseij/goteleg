package goteleg

func GetChatID(update *Update) int64 {
	var chatID int64

	if update.Msg != nil {
		chatID = update.Msg.Chat.ID
	} else if update.EditedMsg != nil {
		chatID = update.EditedMsg.Chat.ID
	} else if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
	} else if update.ChannelPost != nil {
		chatID = update.ChannelPost.Chat.ID
	} else if update.EditedChannelPost != nil {
		chatID = update.EditedChannelPost.Chat.ID
	} else if update.InlineQuery != nil {
		chatID = update.InlineQuery.From.ID
	}

	return chatID
}
