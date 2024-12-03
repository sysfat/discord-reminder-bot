package commands

import (
	"dcbot/internal/logger"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	stopReminders = make(chan bool)
)

func StartReminder(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reminder started! You will receive water reminders every 15 minutes and exercise reminders every 45 minutes.",
		},
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to send start reminder response: %v", err))
		return
	}

	go func() {
		tickerWater := time.NewTicker(15 * time.Minute)
		tickerExercise := time.NewTicker(45 * time.Minute)
		defer tickerWater.Stop()
		defer tickerExercise.Stop()

		for {
			select {
			case <-tickerWater.C:
				_, err := s.ChannelMessageSend(i.ChannelID, "Time to drink water! 💧")
				if err != nil {
					logger.Error(fmt.Sprintf("Failed to send water reminder: %v", err))
				}
			case <-tickerExercise.C:
				_, err := s.ChannelMessageSend(i.ChannelID, "Time to exercise! 🚶‍♂️")
				if err != nil {
					logger.Error(fmt.Sprintf("Failed to send exercise reminder: %v", err))
				}
			case <-stopReminders:
				return
			}
		}
	}()
}

func FinishReminder(s *discordgo.Session, i *discordgo.InteractionCreate) {
	stopReminders <- true

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reminder stopped! Stay hydrated and active! 😊",
		},
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to send finish reminder response: %v", err))
	}
}
