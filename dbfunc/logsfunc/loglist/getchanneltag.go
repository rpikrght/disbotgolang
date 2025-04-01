package loglist

import (
	"log"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func GetChannelTagC(event *events.MessageCreate) (string, error) {
	channel, err := event.Client().Rest().GetChannel(event.ChannelID)
	if err != nil {
		log.Printf("getchannelerr %v", err)
		return "err", err
	}
	channeltag := "err"
	if channel.Type() == discord.ChannelTypeGuildText {
		guildChannel := channel.(discord.GuildChannel)
		channeltag = guildChannel.Name()
	}
	return channeltag, nil
}

func GetChannelTagU(event *events.MessageUpdate) (string, error) {
	channel, err := event.Client().Rest().GetChannel(event.ChannelID)
	if err != nil {
		log.Printf("getchannelerr %v", err)
		return "err", err
	}
	channeltag := "err"
	if channel.Type() == discord.ChannelTypeGuildText {
		guildChannel := channel.(discord.GuildChannel)
		channeltag = guildChannel.Name()
	}
	return channeltag, nil
}
