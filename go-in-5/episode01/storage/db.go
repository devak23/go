package storage

import "errors"

var (
	ErrorNotFound = errors.New("Item not found")
)

type DB interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}
