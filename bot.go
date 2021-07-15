package goteleg

import (
	"log"
)

type LongPollerSettings struct {
	Limit          int
	Timeout        int
	AllowedUpdates []string
}

type Setting struct {
	Token string
	LongPollerSettings
}

type poller interface {
	startPolling(*Bot)
}

type dispatcher struct {
	Handlers map[string]func(*Update)
	updates  chan *Update
	stop     chan struct{}
}

type Bot struct {
	*dispatcher
	poller
	API
}

type longPoller struct {
	offset         int
	limit          int
	timeout        int
	allowedUpdates []string
}

func (lp *longPoller) startPolling(b *Bot) {
	for {
		updates, err := b.GetUpdates(lp.offset+1, lp.timeout)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, update := range updates.Result {
			lp.offset = update.UpdateID
			b.updates <- update
		}
	}
}

func newDispatcher(limits int) *dispatcher {
	return &dispatcher{
		updates: make(chan *Update, limits),
		stop:    make(chan struct{}),
	}
}

func newLongPoller(setting LongPollerSettings) *longPoller {
	return &longPoller{
		limit:          setting.Limit,
		timeout:        setting.Timeout,
		allowedUpdates: setting.AllowedUpdates,
	}
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
	b.Handlers[command] = handler
}

func (d *dispatcher) handleUpdate(update *Update) {
	if update.Msg != nil {
		str := update.Msg.Text
		if str[0] == '/' {
			if handler, ok := d.Handlers[str]; ok {
				handler(update)
			}
		}
	}
}

func (d *dispatcher) listenHttp() {
	for {
		select {
		case update := <-d.updates:
			d.handleUpdate(update)
		case <-d.stop:
			close(d.stop)
			return
		}
	}
}

func (b *Bot) Stop() {
	b.stop <- struct{}{}
}
