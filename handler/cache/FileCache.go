package cache

import (
	"github.com/skydrive/response"
	"sync"
)

type FileCache struct {
	sync.RWMutex
	Map map[string]response.UserFile
}

func NewFileCache() *FileCache {
	tm := new(FileCache)
	tm.Map = make(map[string]response.UserFile)
	return tm
}

func (sm *FileCache) ReadFileCache(key string) (response.UserFile,bool) {
	sm.RLock()
	defer sm.RUnlock()
	byToken, exist := sm.Map[key]
	return byToken,exist
}

func (sm *FileCache) WriteFileCache(key string, value response.UserFile) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}

func (sm *FileCache) DeleteFileCache(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.Map, key)
}
