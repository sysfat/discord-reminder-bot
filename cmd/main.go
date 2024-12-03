package main

import (
	"dcbot/config"
	"dcbot/internal/logger"
	"dcbot/internal/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()

	logger.Info("Creating Discord bot session...")
	bot, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	logger.Info("Connecting to Discord...")
	err = bot.Open()
	if err != nil {
		logger.Error(fmt.Sprintf("Error opening WebSocket connection %s", err))
		return
	}

	logger.Info("Registering commands and handlers...")
	utils.RegisterCommands(bot, config.GuildID)
	utils.RegisterHandlers(bot)

	logger.Info("Bot is running. Press CTRL+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	bot.Close()
	logger.Info("Bot stopped.")
}
