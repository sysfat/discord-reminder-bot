package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BotToken string
	GuildID  string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	BotToken = os.Getenv("BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")

	if BotToken == "" || GuildID == "" {
		log.Fatalf("Missing required environment variables: BOT_TOKEN or GUILD_ID")
	}

	fmt.Println("Configuration loaded successfully.")
}
