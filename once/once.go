// a little util that is like "sync.Once" except we can reset it
// very handy for "start/stop" operations such that things won't start/stop more then once
// See http://golang.org/pkg/sync/#Once

package once

import (
	"sync"
	"sync/atomic"
)

// Once little struct to Do a function once
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do do the function but only do it if things have not been reset (or never run)
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// Reset resets the state such that a function can run again
func (o *Once) Reset() {
	o.m.Lock()
	defer o.m.Unlock()
	atomic.StoreUint32(&o.done, 0)
}

// StartStop start/stop once in a group
type StartStop struct {
	start Once
	stop  Once
}

// Start a function just once and rest the Stop caller, so we can stop things if needed
func (o *StartStop) Start(f func()) {
	o.start.Do(func() {
		f()
		o.stop.Reset()
	})
}

// Stop run a stop function only once and reset the start blocker
func (o *StartStop) Stop(f func()) {
	o.stop.Do(func() {
		f()
		o.start.Reset()
	})
}
