package goteleg

type dispatcher struct {
	handlers map[string]func(*Update)
	stop     chan struct{}
}

func (d *dispatcher) listenUpdatesChan(updates chan *Update) {
	for {
		select {
		case update := <-updates:
			d.handleUpdate(update)
		case <-d.stop:
			close(d.stop)
			return
		}
	}
}

func newDispatcher(limits int) *dispatcher {
	return &dispatcher{
		stop:     make(chan struct{}),
		handlers: make(map[string]func(*Update)),
	}
}
