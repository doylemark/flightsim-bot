package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Connect Connects discord bot to discord
func Connect(token string) {
	Session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err)
		return
	}

	Session.AddHandler(MessageCreate)

	err = Session.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now online, press CTRL+C to terminate the process!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}