package main

import (
	"github.com/Alexseij/goteleg"
)

func main() {
	settings := goteleg.Setting{
		Token: "YOUR_TOKEN",
	}

	bot := goteleg.NewBot(settings)

	bot.AddHandler("/start", func(u *goteleg.Update) {
		bot.SendMessage("Hi !", goteleg.GetChatID(u), nil)
	})

	bot.Start()
}
