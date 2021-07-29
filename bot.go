package goteleg

import "net/http"

type Setting struct {
	Token  string
	Client *http.Client
	LongPollerSettings
}

type Bot struct {
	poller     *longPoller
	dispatcher *dispatcher
	API
}

func NewBot(settings Setting) *Bot {
	if settings.Token == "" {
		panic("goteleg : use token")
	}

	if settings.Limit == 0 {
		settings.Limit = 100
	}

	if settings.Client == nil {
		settings.Client = http.DefaultClient
	}

	dispatcher := newDispatcher(settings.Limit)

	longPoller := newLongPoller(settings.LongPollerSettings)

	bot := &Bot{
		API:        newApi(settings.Token, settings.Client),
		dispatcher: dispatcher,
		poller:     longPoller,
	}

	return bot
}

func (b *Bot) Start() {
	go b.poller.startPolling(b.API)
	b.dispatcher.listenUpdatesChan(b.poller.updates)
}

func (b *Bot) AddHandler(command string, handler func(*Update)) {
	b.dispatcher.handlers[command] = handler
}

func (d *dispatcher) handleUpdate(update *Update) {
	if update.Msg != nil {
		str := update.Msg.Text
		if str[0] == '/' {
			if handler, ok := d.handlers[str]; ok {
				handler(update)
			}
		}
	}
}

func (b *Bot) Stop() {
	b.dispatcher.stop <- struct{}{}
}
