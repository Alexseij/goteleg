package goteleg

type Bot struct {
	API
	chatID int64
}

func NewBot(chatID int64, token string) *Bot {
	return &Bot{
		NewApi(token),
		chatID,
	}
}
