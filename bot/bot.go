package bot

import (
	"fmt"
	"strings"
	"time"

	"discordbot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())

	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := s.ChannelMessageSend
	fmt.Println(m.Author, m.Content, m.Message.ID)
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		m.Content = strings.Replace(m.Content, config.BotPrefix, "", -1)
		fmt.Println(m.Content)
		arg := strings.Split(m.Content, " ")
		fmt.Println(arg)
		cmd := arg[0]
		fmt.Println(m.Message.ID)
		if cmd == "lpb" {
			_, _ = send(m.ChannelID, "Le plus beau c'est Célian et Paco")
		}
		if cmd == "time" {
			currentTime := time.Now()
			_, _ = send(m.ChannelID, "Time : "+currentTime.Format("02-01-2006 15:04:05"))

		}
		if cmd == "sondage" {
			if len(arg) > 1 {

				prout, _ := send(m.ChannelID, "Le sondage de "+m.Author.Username+" est : "+strings.Join(arg[1:], " "))
				_ = s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
				_ = s.MessageReactionAdd(m.ChannelID, prout.ID, "👍")
				_ = s.MessageReactionAdd(m.ChannelID, prout.ID, "👎")
			} else {
				send(m.ChannelID, "Veuillez mettre l'intitulé du sondage")
			}
		}
		if cmd == "echo" {
			if len(arg) > 1 {
				_, _ = send(m.ChannelID, strings.Join(arg[1:], " "))
				_ = s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
			} else {
				_, _ = send(m.ChannelID, "veuillez mettre un argument")
			}
		}
	}
}
