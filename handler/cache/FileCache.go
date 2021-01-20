package cache

import (
	"github.com/skydrive/beans"
	"github.com/skydrive/db"
	"github.com/skydrive/logger"
	"sync"
)

const TAG = "FileCache.go"

var hitTargetCount = 0
var filecache = NewFileCache()

type FileCache struct {
	sync.RWMutex
	Map map[string]beans.File
}

func NewFileCache() *FileCache {
	tm := new(FileCache)
	tm.Map = make(map[string]beans.File)
	return tm
}

func (sm *FileCache) ReadFileCache(key string) (beans.File, bool) {
	sm.RLock()
	defer sm.RUnlock()
	byToken, exist := sm.Map[key]
	return byToken, exist
}

func (sm *FileCache) WriteFileCache(key string, value beans.File) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}

func (sm *FileCache) DeleteFileCache(key string) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.Map, key)
}

func AddOrUpdateFileMeta(filemeta beans.File) {
	filecache.WriteFileCache(filemeta.Filesha1, filemeta)
}

func GetFileMeta(sha1 string) (*beans.File, bool) {
	if userinfo, isExist := filecache.ReadFileCache(sha1); isExist {
		hitTargetCount++
		logger.Info(TAG, "文件缓存获取成功次数:", hitTargetCount, userinfo)
		return &userinfo, true
	}
	//else {
	//	logger.Error(TAG, "缓存获取失败", userinfo)
	//}
	if meta, err := db.GetFileInfoBySha1(sha1); err == nil {
		userinfo := beans.GetFileObject(*meta)
		filecache.WriteFileCache(sha1, *userinfo)
		return userinfo, true
	}
	return nil, false
}
