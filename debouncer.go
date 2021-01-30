package debouncer

import (
	"sync"
	"time"
)

// Debouncer - A debouncer for ensuring that only one event fires when
//             multiple triggers happen inside a defined window of time
type Debouncer struct {
	duration time.Duration
	timer    *time.Timer
	callback func()
	mutex    sync.Mutex
}

// New - Returns a new debouncer with the defined delay and callback
func New(limit time.Duration, callback func()) *Debouncer {
	return &Debouncer{
		duration: limit,
		callback: callback,
	}
}

// Trigger -  Calls the callback after duration of the last calling of trigger
func (g *Debouncer) Trigger() {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.timer != nil {
		if !g.timer.Stop() {
			<-g.timer.C
		}
		g.timer.Reset(g.duration)
	} else {
		g.timer = time.AfterFunc(g.duration, func() {
			g.mutex.Lock()
			g.timer = nil
			g.mutex.Unlock()
			g.callback()
		})
	}
}
