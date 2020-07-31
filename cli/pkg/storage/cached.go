package storage

import (
	"path"
	"strings"
)

// CachedStorage wraps another storage, caching a prefix in a local directory.
//
// SyncCache() syncs cachePrefix locally, which you must call before doing any
// reads. It is not done automatically so you can control output to the user about
// syncing.
//
// If a read hits a path starting with cachePrefix, it will use the local cached version.
type CachedStorage struct {
	storage      Storage
	cachePrefix  string
	cacheDir     string
	cacheStorage *DiskStorage
	isSynced     bool
}

func NewCachedStorage(store Storage, cachePrefix string, cacheDir string) (*CachedStorage, error) {
	// This doesn't actually return an error, but catch in case of future errors
	cacheStorage, err := NewDiskStorage(cacheDir)
	if err != nil {
		return nil, err
	}
	return &CachedStorage{
		storage:      store,
		cachePrefix:  cachePrefix,
		cacheDir:     cacheDir,
		cacheStorage: cacheStorage,
		isSynced:     false,
	}, nil
}

func (s *CachedStorage) Get(p string) ([]byte, error) {
	if strings.HasPrefix(p, s.cachePrefix) {
		return s.cacheStorage.Get(p)
	}
	return s.storage.Get(p)
}

func (s *CachedStorage) Put(p string, data []byte) error {
	// FIXME: potential for cache and remote to get out of sync on error
	if strings.HasPrefix(p, s.cachePrefix) {
		if err := s.cacheStorage.Put(p, data); err != nil {
			return err
		}
	}
	return s.storage.Put(p, data)
}

func (s *CachedStorage) GetDirectory(storagePath string, localPath string) error {
	if strings.HasPrefix(storagePath, s.cachePrefix) {
		return s.cacheStorage.GetDirectory(storagePath, localPath)
	}
	return s.storage.GetDirectory(storagePath, localPath)
}

func (s *CachedStorage) PutDirectory(localPath string, storagePath string) error {
	// FIXME: potential for cache and remote to get out of sync on error
	if strings.HasPrefix(storagePath, s.cachePrefix) {
		if err := s.cacheStorage.PutDirectory(localPath, storagePath); err != nil {
			return err
		}
	}
	return s.storage.PutDirectory(localPath, storagePath)

}

func (s *CachedStorage) List(p string) ([]string, error) {
	if strings.HasPrefix(p, s.cachePrefix) {
		return s.cacheStorage.List(p)
	}
	return s.storage.List(p)
}

func (s *CachedStorage) MatchFilenamesRecursive(results chan<- ListResult, path string, filename string) {
	if strings.HasPrefix(path, s.cachePrefix) {
		s.cacheStorage.MatchFilenamesRecursive(results, path, filename)
		return
	}
	s.storage.MatchFilenamesRecursive(results, path, filename)
}

func (s *CachedStorage) SyncCache() error {
	return s.storage.GetDirectory(s.cachePrefix, path.Join(s.cacheDir, s.cachePrefix))
}
