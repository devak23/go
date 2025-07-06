package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type WordCount struct {
	Word  string
	Count int
}

type Counter struct {
	data map[string]int
	lock sync.RWMutex
}

func (c *Counter) GetSortedResults() []WordCount {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// Convert map into slice for Sorting
	results := make([]WordCount, 0, len(c.data))
	for word, count := range c.data {
		results = append(results, WordCount{Word: word, Count: count})
	}

	// sort by count(descending), then by word Ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].Count == results[j].Count {
			return results[i].Word < results[j].Word
		}
		return results[i].Count > results[j].Count
	})

	return results
}

func (c *Counter) Update(w string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	// if count, exists := c.data[w]; !exists {
	// 	c.data[w] = 1
	// } else {
	// 	c.data[w] = count + 1
	// }
	c.data[w]++
}

func (c *Counter) PrintTokens() {
	// for k, v := range c.data {
	// 	fmt.Printf("%s appears %d times.\n", k, v)
	// }

	for _, result := range c.GetSortedResults() {
		fmt.Printf("'%s' appears %d times\n", result.Word, result.Count)
	}
}

func NewCounter() Counter {
	return Counter{
		data: make(map[string]int),
	}
}

// CleanseAndLowerCaseV2 is a better and efficient version of the CleanseAndLowerCase method below
func CleanseAndLowerCaseV2(line string) string {
	var builder strings.Builder
	for _, ch := range line {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == ' ' {
			builder.WriteRune(ch)
		}
	}
	return strings.ToLower(builder.String())
}

func CleanseAndLowerCase(line string) string {
	cleansedLine := regexp.MustCompile(`[^\w\s]`).ReplaceAllString(line, "")
	return strings.ToLower(cleansedLine)
}

func WordCountMain() {
	c := NewCounter()
	line := "Hello world! This is a test. Hello again, world."
	line = CleanseAndLowerCaseV2(line)

	tokens := strings.Fields(line)
	for _, token := range tokens {
		if token != "" {
			c.Update(token)
		}
	}
	c.PrintTokens()
}
