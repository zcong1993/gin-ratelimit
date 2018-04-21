package ratelimit

import (
	"sync/atomic"
	"time"
)

// Item is limit state for a ip
type Item struct {
	num      int64
	expireIn int64
}

// Ratelimiter is struct for ratelimitor
type Ratelimiter struct {
	s         map[string]*Item
	duration  int64
	rateLimit int64
}

// NewLimiter is a constructor for Ratelimiter
func NewLimiter(duration, rateLimit int64) *Ratelimiter {
	s := &Ratelimiter{
		s:         make(map[string]*Item),
		duration:  duration,
		rateLimit: rateLimit,
	}
	return s
}

func (s *Ratelimiter) incr(k string) {
	i, ok := s.s[k]
	if !ok {
		s.s[k] = &Item{num: 1, expireIn: time.Now().Unix() + s.duration}
		return
	}
	atomic.AddInt64(&i.num, 1)
}

// ShouldLimit return if user should be block
func (s *Ratelimiter) ShouldLimit(k string) bool {
	i, ok := s.s[k]
	if !ok {
		s.incr(k)
		return false
	}
	s.incr(k)
	if i.expireIn <= time.Now().Unix() {
		delete(s.s, k)
		s.incr(k)
		return false
	}
	if atomic.LoadInt64(&i.num) > s.rateLimit {
		return true
	}
	return false
}
