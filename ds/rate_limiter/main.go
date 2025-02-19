package main

import (
	"sync/atomic"
	"time"
)

// Develop a rate limiter which can accept t requests in a given time interval s seconds.
// If the number of requests exceeds t in s seconds, the rate limiter should reject the request.

type RingBuffer struct {
	data []int64
	size int
	head int32
	tail int32
}

func newRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data: make([]int64, size),
		size: size,
	}
}

func (r *RingBuffer) add(val int64) bool {
	for {
		tail := atomic.LoadInt32(&r.tail)
		head := atomic.LoadInt32(&r.head)

		if tail == (head)%int32(r.size) && head != 0 {
			return false
		}

		nextTail := (tail + 1) % int32(r.size)
		if atomic.CompareAndSwapInt32(&r.tail, tail, nextTail) {
			r.data[r.tail] = val
			return true
		}
	}
}

func (r *RingBuffer) get() int64 {
	for {
		head := atomic.LoadInt32(&r.head)
		if head == r.tail {
			return 0
		}
		nextHead := (head + 1) % int32(r.size)
		if atomic.CompareAndSwapInt32(&r.head, head, nextHead) {
			return r.data[r.head]
		}
	}
}

var rb = newRingBuffer(1)

func main() {
	var ct = time.Now()
	rateLimit(ct.Unix())
	rateLimit(ct.Unix())
}

func rateLimit(t int64) {
	if !rb.add(t) {
		println("Request rejected")
		return
	}
}
