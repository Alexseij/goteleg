package goteleg

import "log"

type longPoller struct {
	offset         int
	timeout        int
	allowedUpdates []string
	updates        chan *Update
}

type LongPollerSettings struct {
	Limit          int
	Timeout        int
	AllowedUpdates []string
}

func (lp *longPoller) startPolling(api API) {
	for {
		updates, err := api.GetUpdates(lp.offset+1, lp.timeout)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, update := range updates.Result {
			lp.offset = update.UpdateID
			lp.updates <- update
		}
	}
}

func newLongPoller(setting LongPollerSettings) *longPoller {
	return &longPoller{
		timeout:        setting.Timeout,
		allowedUpdates: setting.AllowedUpdates,
		updates:        make(chan *Update, setting.Limit),
	}
}
