package main

import (
	"github.com/Alexseij/goteleg"
)

func main() {
	settings := goteleg.Setting{
		Token: "1720667998:AAGL4BuASvV42Z1yGK_tqc6swos_xXQ3CTA",
	}

	bot := goteleg.NewBot(settings)

	bot.Handlers = map[string]func(*goteleg.Update){
		"/start": func(u *goteleg.Update) {
			bot.SendMessage("Hi !", goteleg.GetChatID(u), nil)
		},
	}

	bot.Start()
}
