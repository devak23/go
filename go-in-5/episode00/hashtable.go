package episode00

import "errors"

var (
	ErrorNotFound = errors.New("not found")
)

type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
