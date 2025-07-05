package episode00

import (
	. "fmt"
	"gopkg.in/redis.v3"
)

// RedisClientWrapper is a wrapper for the Client type the gopkg.in/redis.v3 package
// This type allows to redefine the Get and Set method using the parameter types and return types that match the
// Hashtable interface.
type RedisClientWrapper struct {
	Client *redis.Client
}

// Get method of the RedisClientWrapper type
func (rcw *RedisClientWrapper) Get(key string) ([]byte, error) {
	val, ok := rcw.Client.Get(key).Result()
	if ok != nil {
		return nil, Errorf("error while getting key: %s", key)
	}
	return []byte(val), nil
}

// Set method of the RedisClientWrapper type
func (rcw *RedisClientWrapper) Set(key string, value []byte) error {
	return rcw.Client.Set(key, value, 0).Err()
}
