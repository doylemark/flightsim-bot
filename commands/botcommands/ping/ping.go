package ping

import (
	"time"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Ping command for checking bot status and response time
func Ping(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgTime, err := time.Parse(time.RFC3339, string(message.Timestamp))
	responseTime := time.Now().Sub(msgTime)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("🏓 Pong! response time was %s 😍", responseTime.String()))
}