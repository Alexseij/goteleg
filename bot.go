package goteleg

type Setting struct {
	Token string
	LongPollerSettings
}

type Bot struct {
	*dispatcher
	poller
	API
}

func NewBot(settings Setting) *Bot {
	if settings.Token == "" {
		panic("Set token!!")
	}

	if settings.Limit == 0 {
		settings.Limit = 100
	}

	d := newDispatcher(settings.Limit)

	longPoller := newLongPoller(settings.LongPollerSettings)

	bot := &Bot{
		API:        newApi(settings.Token),
		dispatcher: d,
		poller:     longPoller,
	}

	return bot
}

func (b *Bot) Start() {
	go b.startPolling(b)
	b.listenHttp()
}

func (b *Bot) AddHandler(command string, handler func(*Update)) {
	b.handlers[command] = handler
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
	b.stop <- struct{}{}
}
