package main

import (
	"blackhatgo/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	// Fatal helpers - program exits on error
	config := utils.Must(os.Open("../data/config.json"))
	defer utils.TryClose(config, "config file")

	// Non-fatal logging
	if err := someOperation(); err != nil {
		utils.LogErrorF(err, "failed to perform operation")
	}

	// Conditional execution
	data, err := os.ReadFile("data.txt")
	utils.CheckF(err, "failed to read data file")

	// Different approaches for different scenarios

	// 1. Must succeed (panic on error)
	db := utils.Must(connectToDatabase())
	fmt.Println(db)

	// 2. Log and continue with default
	port, err := getConfigPort()
	utils.OrDefault(port, err, 8080)

	// 3. Log warning but continue
	if err := optionalCleanup(); err != nil {
		utils.LogWarnF(err, "cleanup failed, continuing anyway")
	}

	// 4. Ignore expected errors
	utils.IgnoreF(os.Remove("temp.txt"), "removing temp file")

	log.Printf("Server starting on port %d with %d bytes of data", port, len(data))
	log.Println("Application initialized successfully")
}

// Mock functions for demonstration
func someOperation() error {
	return nil
}

func connectToDatabase() (*struct{}, error) {
	return &struct{}{}, nil
}

func getConfigPort() (int, error) {
	return 0, fmt.Errorf("port not configured")
}

func optionalCleanup() error {
	return nil
}
