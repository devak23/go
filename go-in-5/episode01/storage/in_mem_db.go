package storage

import "sync"

type inMemDB struct {
	Data map[string][]byte
	lock sync.RWMutex
}

func NewInMemDB() DB {
	return &inMemDB{
		Data: make(map[string][]byte),
	}
}

func (db *inMemDB) Set(key string, val []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.Data[key] = val
	return nil
}

func (db *inMemDB) Get(key string) ([]byte, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	return db.Data[key], nil
}
