/*
   A simple "shared" place (aka global) to put various data related to just about anything we want

   Harken back to the C, C++ days for global accessed data.

   This probably should not be used for very frequently changing data.  But quiet good for sharing config like data
   across nodes/divergent processes.
*/

package shared

import (
	"sync"
)

//singleton
type globalData map[string]interface{}

var data globalData
var mu *sync.RWMutex

func init() {
	data = make(globalData, 0)
	mu = new(sync.RWMutex)
}

// Get something from the global puddle
func Get(name string) interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return data[name]
}

func Set(name string, indata interface{}) {
	mu.Lock()
	defer mu.Unlock()
	data[name] = indata
}

func GetAll() globalData {
	// makes a copy we don't want random post getalls to mess w/ what's in here
	cp := make(map[string]interface{})
	mu.RLock()
	defer mu.RUnlock()

	for k, v := range data {
		cp[k] = v
	}
	return globalData(cp)
}
