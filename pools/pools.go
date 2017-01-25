/**

pools

A collection of standard sync.Pools that one seems to use all the time

*/

package pools

import (
	"bytes"
	"hash"
	"hash/fnv"
	"sync"
)

// set up a byte[] sync pool for better GC allocation of byte arraies
// bufPool is a pool for staging buffers. Using a pool allows concurrency-safe
// reuse of buffers
var bytesPool sync.Pool

// GetBytes get a byte slice from the pool if length l
func GetBytes(l int) []byte {
	x := bytesPool.Get()
	if x == nil {
		return make([]byte, l)
	}
	buf := x.([]byte)
	if cap(buf) < l {
		return make([]byte, l)
	}
	return buf[:l]
}

// PutBytes put the byte slice back in the pool
func PutBytes(buf []byte) {
	bytesPool.Put(buf)
}

var bytesBufferPool sync.Pool

// GetBytesBuffer get a new bytes.Buffer from the pool, with its old contents reset
func GetBytesBuffer() *bytes.Buffer {
	x := bytesBufferPool.Get()
	if x == nil {
		return &bytes.Buffer{}
	}
	buf := x.(*bytes.Buffer)
	buf.Reset()
	return buf
}

// PutBytesBuffer put the bytes.Buffer back in the pool
func PutBytesBuffer(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}

// for dealing w/ read buffer copies
var GetSyncBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer([]byte{})
	},
}

var mutexPool = sync.Pool{
	New: func() interface{} {
		return new(sync.Mutex)
	},
}

// GetMutex get a simple mutex pointer
func GetMutex() *sync.Mutex {
	return mutexPool.Get().(*sync.Mutex)
}

// PutMutex put the mutex point back in the pool
func PutMutex(mu *sync.Mutex) {
	mutexPool.Put(mu)
}

var rwMutexPool = sync.Pool{
	New: func() interface{} {
		return new(sync.RWMutex)
	},
}

// GetRWMutex get a RW mutex from the pool
func GetRWMutex() *sync.RWMutex {
	return rwMutexPool.Get().(*sync.RWMutex)
}

// PutRWMutex put a RW mutex into the pool
func PutRWMutex(mu *sync.RWMutex) {
	rwMutexPool.Put(mu)
}

var waitGroupPool = sync.Pool{
	New: func() interface{} {
		return new(sync.WaitGroup)
	},
}

// GetWaitGroup get a sync.WaitGroup from the pool
func GetWaitGroup() *sync.WaitGroup {
	return waitGroupPool.Get().(*sync.WaitGroup)
}

// PutWaitGroup a waitgroup into the pool
func PutWaitGroup(mu *sync.WaitGroup) {
	mutexPool.Put(mu)
}

// a little hash pool for GC pressure easing

var fn64avPool sync.Pool

// GetFnv64a get a fnv64a from a pool
func GetFnv64a() hash.Hash64 {
	x := fn64avPool.Get()
	if x == nil {
		return fnv.New64a()
	}
	out := x.(hash.Hash64)
	out.Reset()
	return out
}

// PutFnv64a put a fnv64a into a pool
func PutFnv64a(spl hash.Hash64) {
	fn64avPool.Put(spl)
}
