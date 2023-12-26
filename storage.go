package storage

import (
	"container/heap"
	"errors"
	"sync"
	"time"
)

const (
	defaultTTL = 10 * time.Hour
)

type Storage interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
}

type storage struct {
	data       map[string]interface{}
	expiration map[string]time.Time
	mu         sync.Mutex
	ttlHeap    TTLHeap
}

func New() Storage {
	s := &storage{
		data:       make(map[string]interface{}),
		expiration: make(map[string]time.Time),
		ttlHeap:    make(TTLHeap, 0),
	}
	heap.Init(&s.ttlHeap)
	go s.TTLChecker()
	return s
}

func (s *storage) Get(key string) (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if expiration, ok := s.expiration[key]; ok {
		if expiration.Before(now) {
			s.deleteKey(key)
			return nil, errors.New("key expired")
		}
		return s.data[key], nil
	}
	return nil, errors.New("key not found")

}

func (s *storage) Set(key string, value interface{}, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	if ttl == 0 {
		ttl = defaultTTL
	}
	expiration := time.Now().Add(ttl)
	s.expiration[key] = expiration
	s.ttlHeap.Push(&Entry{Key: key, Expiration: expiration})
}

func (s *storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.deleteKey(key)
}

func (s *storage) TTLChecker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			for {
				if s.ttlHeap.Len() == 0 {
					break
				}
				entry := s.ttlHeap[0]
				if entry.Expiration.After(time.Now()) {
					break
				}
				s.ttlHeap.Pop()
				s.deleteKey(entry.Key)
			}
			s.mu.Unlock()
		}
	}
}

func (s *storage) deleteKey(key string) {
	delete(s.data, key)
	delete(s.expiration, key)
}
