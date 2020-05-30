package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/doylemark/flightsim-bot/commands/env"

	"github.com/bwmarrin/discordgo"
)

// HandleMetar parses a ICAO from message content and calls GetMetar
func HandleMetar(session *discordgo.Session, message *discordgo.MessageCreate) {
	stringArr := strings.Fields(message.Content)

	if len(stringArr) != 2 {
		session.ChannelMessageSend(message.ChannelID, "You must provide an icao code and nothing else, ex. `!metar EIDW`")
		return
	}

	getMetar(stringArr[1], message.ChannelID, session)
}

// GetMetar fetches and returns a METAR object from avwx api
func getMetar(icao string, channel string, session *discordgo.Session) {
	url := "http://metar.vatsim.net/metar.php?id=" + icao

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		session.ChannelMessageSend(channel, "There was an error fetching your METAR")
		return
	}

	req.Header.Set("Authorization", env.WxToken)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		session.ChannelMessageSend(channel, "There was an error fetching your METAR")
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		session.ChannelMessageSend(channel, "There was an error fetching your METAR")
		return
	}

	metar := string(body)

	if len(metar) == 0 {
		session.ChannelMessageSend(channel, "No METAR could be found for the provided ICAO")
		return
	}

	_, err = session.ChannelMessageSend(channel, string(body))

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(channel, "There was an error fetching your METAR")
		return
	}
}
