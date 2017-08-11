package assorted

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
	"bufio"
	"time"
	"strings"
)

// this function connects to redis to load the reference data
func loadWords(conn redis.Conn) {
	// get the curse_words from redis
	curse_words, _ := getValuesFromRedisSet(conn, "curse_words")

	if len(curse_words) == 0 {
		addRefData(conn)
	}
}

// This is the main functiot that gets invoked from the main.go
func SwearFilterMain() {
	// connect to redis server
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("Cannot connect to redis: %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	// load the reference data
	loadWords(conn)

	// start the main loop
	for {
		fmt.Println("Enter some sentences with the swear words. Then press Enter or \"q\" to quit")
		bio := bufio.NewReader(os.Stdin)
		line, _, _ := bio.ReadLine()

		if string(line) == "q" {
			break
		}
		terms := strings.Split(string(line), " ")
		analyze(conn, terms)
	}

	fmt.Println("Session ended.")
}

// this function analyzes the user data and tries to find the swear words in a sentence
// it does this by creating a new key in redis called user_words. It then loads all the
// data user entered against this key and then takes the intersection of two sets of words
// i.e. the curse words and the user_words.
func analyze(conn redis.Conn, userData []string) {
	// asks redis to start a transaction as the next commands will have to be queued for atomic
	// execution using EXEC command
	conn.Send("MULTI")
	// delete the temp set "user_words
	conn.Send("DEL", "user_words")
	// add the user_data
	conn.Send("SADD", redis.Args{}.Add("user_words").AddFlat(userData)...)
	// take the intersection of both sets
	conn.Send("SINTER", "user_words", "curse_words")

	reply, err := conn.Do("EXEC")

	if err != nil {
		fmt.Printf("Error executing EXEC: %s",err)
		os.Exit(1)
	}

	values, _ := redis.Values(reply, nil)

	curse_words, err := redis.Strings(values[2], nil)
	if err != nil {
		fmt.Printf("Error parsing the reply from redis: %s", err)
	}

	if len(curse_words) > 0 {
		for _, word := range curse_words {
			fmt.Println(">> Found: ", word)
		}
	}
}

// a function to insert words into redis using a standard data source
func addRefData(conn redis.Conn) error {
	// open the list of swear words
	startTime := time.Now()
	file, err := os.Open("resources/swear_words.txt")
	if err != nil {
		fmt.Printf("Cannot read the file resources/swear_words.txt: %s\n", err)
		os.Exit(0)
	}
	defer file.Close()

	// create a slice variable
	scanner := bufio.NewScanner(file)
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

	return values, nil
}
