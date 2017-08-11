package assorted

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
)

func RedisPrimerMain()  {
	// make a connection to redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("Cannot connect to redis: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close() // when done, close the connection

	// check if the swear words are present in Redis
	values, err := getValuesFromSet(conn, "curse_words")

	if len(values) == 0 {
		// if they are not, load them
		conn.Do("SADD", "curse_words", "shit" ,"crap", "fuck")
	}

	// again get them from Redis
	values, err = getValuesFromSet(conn, "curse_words")

	// print the values received from redis. This is an indexed array
	array, _ := redis.Strings(values, nil)
	for index, curseWord := range array {
		fmt.Println(index, " => ", curseWord)
	}
}


// function takes in the redis connection and checks if the key lies in the SET of values
func getValuesFromSet(conn redis.Conn, key string) ([]interface{}, error) {
	values, err := redis.Values(conn.Do("SMEMBERS", key))
	if err != nil {
		fmt.Printf("Could not query redis: %s\n", err)
	}
	return values, err
}
