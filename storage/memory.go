package storage

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type memory struct {
	Store
	sync.RWMutex

	data    map[string][]int
	created map[string]time.Time

	ttl time.Duration
}

// NewMemoryStore creates a Store based on internal memory.
func NewMemoryStore(ttl time.Duration) Store {
	return &memory{
		data:    make(map[string][]int),
		created: make(map[string]time.Time),

		ttl: ttl,
	}
}

func (m *memory) CreateSession() string {
	m.cleanup()
	session := uuid.Must(uuid.NewRandom()).String()
	m.created[session] = time.Now()
	return session
}

func (m *memory) StoreGrid(session string, data []int) error {
	m.Lock()
	defer m.Unlock()
	m.data[session] = data

	return nil
}

func (m *memory) LoadGrid(session string) ([]int, error) {
	m.RLock()
	defer m.RUnlock()
	return m.data[session], nil
}

func (m *memory) cleanup() {
	var clean []string
	for session, created := range m.created {
		if time.Since(created) > m.ttl {
			clean = append(clean, session)
		}
	}

	if len(clean) > 0 {
		m.Lock()
		defer m.Unlock()
		for _, session := range clean {
			delete(m.data, session)
			delete(m.created, session)
		}
	}
}
