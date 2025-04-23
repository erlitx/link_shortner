package cache

import (
	"container/list"
	"sync"

	"github.com/erlitx/link_shortner/internal/domain"
	"github.com/erlitx/link_shortner/internal/dto"
)

type Cache struct {
	capacity int
	mu       sync.Mutex
	order    *list.List
	mx       sync.RWMutex
	items    map[string]*list.Element
}

type entry struct {
	key   string
	value domain.URL
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		order:    list.New(),
		items:    make(map[string]*list.Element),
	}
}

func (c *Cache) Set(url domain.URL) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if el, found := c.items[string(url.ShortURL)]; found {
		el.Value.(*entry).value = url
		c.order.MoveToFront(el)
		return
	}

	if c.order.Len() >= c.capacity {
		last := c.order.Back()
		if last != nil {
			c.order.Remove(last)                     // Remove from list
			delete(c.items, last.Value.(*entry).key) // Remove from map
		}
	}

	e := &entry{key: string(url.ShortURL), value: url}
	el := c.order.PushFront(e)         // Пушим в лист сразу впереди
	c.items[string(url.ShortURL)] = el // Записываем в мапу
}

func (c *Cache) Get(input dto.GetURLInput) (domain.URL, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, found := c.items[input.ShortUrl]; found {
		c.order.MoveToFront(el)
		return el.Value.(*entry).value, true
	}
	
	return domain.URL{}, false
}
