package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func ConnectOpenAI() *openai.Client {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	TOKEN := os.Getenv("OPEN_AI_TOKEN")
	client := openai.NewClient(TOKEN)

	fmt.Println("Connected to Open AI")
	return client
}
