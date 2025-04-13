// File: internal/service/shortener.go
package service

import (
	"fmt"
	"sort"
	"sync"
	"url-shortener/internal/storage"
	"url-shortener/internal/utils"
)

type ShortenerService struct {
	store  *storage.MemoryStore
	mutex  sync.Mutex
	counts map[string]int
}

type DomainStat struct {
	Domain string
	Count  int
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
	if domain != "" {
		s.counts[domain]++
	}

	return fmt.Sprintf("http://localhost:8080/%s", id)
}

func (s *ShortenerService) Resolve(id string) string {
	return s.store.GetURL(id)
}

func (s *ShortenerService) TopDomains(n int) []DomainStat {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range s.counts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var result []DomainStat
	for i := 0; i < n && i < len(ss); i++ {
		result = append(result, DomainStat{Domain: ss[i].Key, Count: ss[i].Value})
	}

	return result
}
