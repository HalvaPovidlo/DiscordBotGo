package firestore

import (
	"github.com/HalvaPovidlo/discordBotGo/pkg/contexts"
	"sync"
	"time"

	"github.com/HalvaPovidlo/discordBotGo/internal/storage"
)

type Item struct {
	song    storage.Song
	updated time.Time
}

type CacheKey string

type SongsCache struct {
	sync.RWMutex
	songs map[string]Item
}

func NewSongsCache(ctx contexts.Context, expirationTime time.Duration) *SongsCache {
	c := &SongsCache{
		songs: make(map[string]Item),
	}
	c.expireProcess(ctx, expirationTime)
	return c
}

func (c *SongsCache) Get(k CacheKey) (*storage.Song, bool) {
	c.RLock()
	defer c.RUnlock()
	s, ok := c.songs[string(k)]
	if !ok {
		return nil, false
	}
	return &s.song, true
}

func (c *SongsCache) Set(k CacheKey, song *storage.Song) {
	if song == nil {
		return
	}
	c.Lock()
	c.songs[string(k)] = Item{
		song:    *song,
		updated: time.Now(),
	}
	c.Unlock()
}

func (c *SongsCache) KeyFromID(s storage.SongID) CacheKey {
	return CacheKey(s.String())
}

func (c *SongsCache) expireProcess(ctx contexts.Context, expirationTime time.Duration) {
	ticker := time.NewTicker(expirationTime)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				c.Lock()
				now := time.Now()
				for k, v := range c.songs {
					if v.updated.Before(now.Add(-expirationTime)) {
						delete(c.songs, k)
					}
				}
				c.Unlock()
			}
		}
	}()
}
