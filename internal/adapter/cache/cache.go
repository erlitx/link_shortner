package cache

import (
	"github.com/google/uuid"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
	"sync"
)

type Cache struct {
	mx sync.RWMutex
	m  map[uuid.UUID]domain.Profile
}

func New() *Cache {
	return &Cache{
		m: make(map[uuid.UUID]domain.Profile),
	}
}

func (c *Cache) Add(key uuid.UUID, profile domain.Profile) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = profile
}

func (c *Cache) Get(key uuid.UUID) (domain.Profile, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	p, ok := c.m[key]
	if !ok {
		return p, domain.ErrNotFound
	}

	return p, nil
}

func (c *Cache) Update(key uuid.UUID, profile domain.Profile) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = profile
}

func (c *Cache) Delete(key uuid.UUID) {
	c.mx.Lock()
	defer c.mx.Unlock()

	delete(c.m, key)
}
