package goteleg

type dispatcher struct {
	handlers map[string]func(*Update)
	updates  chan *Update
	stop     chan struct{}
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

func newDispatcher(limits int) *dispatcher {
	return &dispatcher{
		updates:  make(chan *Update, limits),
		stop:     make(chan struct{}),
		handlers: make(map[string]func(*Update)),
	}
}
