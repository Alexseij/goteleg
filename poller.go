package goteleg

import (
	"log"
)

type longPoller struct {
	offset         int
	timeout        int
	limit          int
	allowedUpdates []string
	updates        chan *Update
}

type LongPollerSettings struct {
	Limit          int
	Timeout        int
	AllowedUpdates []string
}

func (lp *longPoller) startPolling(api API) {

	options := map[string]interface{}{
		"offset":  lp.offset,
		"timeout": lp.timeout,
		"limit":   lp.limit,
	}

	offset := lp.offset

	for {
		updates, err := api.GetUpdates(options)
		if err != nil {
			log.Println(err)
			continue
		}
		offset++
		options["offset"] = offset
		for _, update := range updates.Result {
			lp.offset = update.UpdateID
			lp.updates <- update
		}
	}
}

func newLongPoller(setting LongPollerSettings) *longPoller {
	return &longPoller{
		limit:          setting.Limit,
		timeout:        setting.Timeout,
		allowedUpdates: setting.AllowedUpdates,
		updates:        make(chan *Update, setting.Limit),
	}
}
