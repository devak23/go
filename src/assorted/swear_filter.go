package assorted

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
	"bufio"
	"time"
)

func SwearFilterMain() {
	// connect to redis server
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("Cannot connect to redis: %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	// delete the user-defined words
	conn.Do("DEL", "user_words")

	// get the curse_words from redis
	curse_words, err := getValuesFromRedisSet(conn, "curse_words")
	fmt.Println(curse_words)

	if len(curse_words) == 0 {
		addWordsToSet(conn)
	}

	//now for the main program
}

// a function to insert words into redis using a standard data source
func addWordsToSet(conn redis.Conn) error {
	// open the list of swear words
	startTime := time.Now()
	file, err := os.Open("resources/swear_words.txt")
	if err != nil {
		fmt.Printf("Cannot read the file resources/swear_words.txt: %s\n", err)
		os.Exit(0)
	}
	defer file.Close()


	// create a slice variable
	scanner:= bufio.NewScanner(file)
	for scanner.Scan() {
		_, err = conn.Do("SADD", "curse_words", scanner.Text())
	}

	fmt.Printf("Time taken to read the file: %d microseconds", time.Since(startTime)/1000)
	return err
}


func getValuesFromRedisSet(conn redis.Conn, key string) ([]interface{}, error) {
	values, err := redis.Values(conn.Do("SMEMBERS", key))
	if err != nil {
		fmt.Printf("Error reading values from Redis: %s\n", err)
		return nil, err
	}

	return values,nil
}