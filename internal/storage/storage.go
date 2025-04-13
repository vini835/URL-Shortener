package storage

import (
	"strconv"
	"sync"
)

type MemoryStore struct {
	urlToID map[string]string
	idToURL map[string]string
	counter int
	mutex   sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		urlToID: make(map[string]string),
		idToURL: make(map[string]string),
		counter: 1,
	}
}

func (s *MemoryStore) Save(url string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if id, exists := s.urlToID[url]; exists {
		return id
	}

	id := strconv.Itoa(s.counter)
	s.counter++

	s.urlToID[url] = id
	s.idToURL[id] = url

	return id
}

func (s *MemoryStore) GetID(url string) (string, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id, exists := s.urlToID[url]
	return id, exists
}

func (s *MemoryStore) GetURL(id string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.idToURL[id]
}
