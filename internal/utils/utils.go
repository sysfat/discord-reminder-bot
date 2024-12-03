package utils

import (
	"dcbot/internal/commands"
	"dcbot/internal/logger"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(bot *discordgo.Session, guildID string) {
	botCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "start",
			Description: "Start reminder",
		},
		{
			Name:        "finish",
			Description: "Finish reminder",
		},
	}

	for _, cmd := range botCommands {
		_, err := bot.ApplicationCommandCreate(bot.State.User.ID, guildID, cmd)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to register command %s: %v", cmd.Name, err))
		} else {
			logger.Info(fmt.Sprintf("Successfully registered command: %s", cmd.Name))
		}
	}
}

func RegisterHandlers(bot *discordgo.Session) {
	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		logger.Info(fmt.Sprintf("User %s triggered the /%s command", i.Member.User.Username, i.ApplicationCommandData().Name))

		switch i.ApplicationCommandData().Name {
		case "start":
			commands.StartReminder(s, i)
		case "finish":
			commands.FinishReminder(s, i)
		}
	})
}
