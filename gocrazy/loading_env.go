package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadingEnvMain() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_ACCESS_KEY")

	fmt.Printf("S3: %s and secret: %s", s3Bucket, secretKey)
}
