package db

import (
	"storiesservice/pkg/utctime"
	"sync"
	"time"
)

type QueryCache struct {
	mtx                                 sync.Mutex
	cache                               interface{}
	lastUpdate                          time.Time
	forceUpdated                        bool
	forceUpdateInterval, updateInterval time.Duration
}

func NewQueryCache(forceUpdateInterval, updateInterval time.Duration) *QueryCache {
	return &QueryCache{
		mtx:                 sync.Mutex{},
		cache:               nil,
		lastUpdate:          time.Unix(0, 0),
		forceUpdated:        false,
		forceUpdateInterval: forceUpdateInterval,
		updateInterval:      updateInterval,
	}
}

func (q *QueryCache) ForceUpdate() {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.forceUpdated = true
}

func (q *QueryCache) Modify(callback func() interface{}) interface{} {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	t := utctime.Get()
	if q.forceUpdated && t.After(q.lastUpdate.Add(q.forceUpdateInterval)) || t.After(q.lastUpdate.Add(q.updateInterval)) {
		q.cache = callback()
		q.lastUpdate = utctime.Get()
		q.forceUpdated = false
	}
	return q.cache
}
