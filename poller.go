package goteleg

import "log"

type poller interface {
	startPolling(*Bot)
}

type longPoller struct {
	offset         int
	limit          int
	timeout        int
	allowedUpdates []string
}

type LongPollerSettings struct {
	Limit          int
	Timeout        int
	AllowedUpdates []string
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

func newLongPoller(setting LongPollerSettings) *longPoller {
	return &longPoller{
		limit:          setting.Limit,
		timeout:        setting.Timeout,
		allowedUpdates: setting.AllowedUpdates,
	}
}
