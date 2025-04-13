// File: internal/service/shortener.go
package service

import (
	"fmt"
	"sync"
	"url-shortener/internal/storage"
	"url-shortener/internal/utils"
)

type ShortenerService struct {
	store  *storage.MemoryStore
	mutex  sync.Mutex
	counts map[string]int
}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{
		store:  storage.NewMemoryStore(),
		counts: make(map[string]int),
	}
}

func (s *ShortenerService) Shorten(url string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if id, exists := s.store.GetID(url); exists {
		return fmt.Sprintf("http://localhost:8080/%s", id)
	}

	id := s.store.Save(url)
	domain := utils.ExtractDomain(url)
	s.counts[domain]++

	return fmt.Sprintf("http://localhost:8080/%s", id)
}
