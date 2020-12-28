package cache

import (
	"github.com/skydrive/db"
	"sync"
)

type TokenCache struct {
	sync.RWMutex
	Map map[string]db.TableUToken
}

func NewTokenMap() *TokenCache {
	tm := new(TokenCache)
	tm.Map = make(map[string]db.TableUToken)
	return tm
}

func (sm *TokenCache) ReadTokenMap(key string) (db.TableUToken, bool) {
	sm.RLock()
	defer sm.RUnlock()
	byToken, exist := sm.Map[key]
	return byToken, exist
}

func (sm *TokenCache) WriteTokenMap(key string, value db.TableUToken) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}

func (sm *TokenCache) DeleteTokenMap(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.Map, key)
}
