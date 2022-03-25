package cache

import (
	"time"
	"ziglu_tech_test/data"
)

type Cache struct {
	Ttl        int64
	lastUpdate int64
	Feeds      data.Feed
}

func (c *Cache) GetFeed() []*data.Item {
	// check cache times
	if c.lastUpdate != 0 && c.lastUpdate+c.Ttl < time.Now().Unix() {
		// if withing cache times return cache
		return c.Feeds.Feed
	}

	// or get fresh feed
	c.Feeds.RefreshFeed()
	c.lastUpdate = time.Now().Unix()

	return c.Feeds.Feed
}
