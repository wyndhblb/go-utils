/*
  This is a helper WaitGroup singleton that helps with doing shutdowns in a nice fashion

  Any sort of "Stop()" or "Shutdown()" command should add it self to the this using AddToShutdown

  the "caller" of the shutdown should then mark it as done.

  the "root" caller of the shutdown (usually a SIGINT signal) then should wait for everything to finish
  and finally "exit"
*/

package shutdown

import "sync"

// singleton
var _SHUTDOWN_WAITGROUP sync.WaitGroup

// AddToShutdown add counter to the shutdown wait group
func AddToShutdown() {
	_SHUTDOWN_WAITGROUP.Add(1)
}

// ReleaseFromShutdown remove from wait group
func ReleaseFromShutdown() {
	_SHUTDOWN_WAITGROUP.Done()
}

// WaitOnShutdown just wait for things to finish
func WaitOnShutdown() {
	_SHUTDOWN_WAITGROUP.Wait()
}
