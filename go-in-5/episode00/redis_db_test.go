package episode00

import (
	"gopkg.in/redis.v3"
	"testing"
)

func TestBusinessLogicWithRedis(t *testing.T) {
	// initialize Redis client here
	ht := RedisClientWrapper{
		redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password
			DB:       0,  // use default DB
		}),
	}

	BusinessLogic(&ht)
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("Expected %s, got %s", "world", val)
	}
}
