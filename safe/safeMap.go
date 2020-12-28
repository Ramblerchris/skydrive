package safe

import "sync"

type SafeMap struct {
	sync.RWMutex
	Map map[string]int
}

func (sm *SafeMap) ReadMap(key string) int {
	sm. RLock()
	defer sm. RUnlock()
	return sm.Map[key]
}
func (sm *SafeMap) WriteMap(key string, value int) {
	sm. Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}


func (sm *SafeMap) DeleteMap(key string) {
	sm. RLock()
	defer sm. RUnlock()
	delete(sm.Map,key)
}
