package loglist

import (
	"strings"

	"github.com/disgoorg/disgo/discord"
)

func GetContent(message *discord.Message) string {
	content := message.Content

	if len(message.Attachments) > 0 {
		urls := make([]string, len(message.Attachments))
		for i, a := range message.Attachments {
			urls[i] = a.URL
		}
		content += "//image " + strings.Join(urls, ", ")
	}

	if len(message.Embeds) > 0 {
		content += "//Embeds" + message.Embeds[0].Description
	}
	if content == "" {
		return "NULL"
	}
	return content
}
