package logsfunc

import (
	"time"

	"disbotgolang/dbfunc/logsfunc/loglist"

	"github.com/disgoorg/disgo/events"
)

func AddC(event *events.MessageCreate) error {
	var logs logslist
	var err error
	logs.time = time.Now().Format("2006-01-02 15:04:05 MST")
	logs.message = loglist.GetContent(&event.Message)
	logs.messageid = event.Message.ID.String()
	logs.user = event.Message.Author.Tag()
	logs.userid = event.Message.Author.ID.String()
	logs.channel, err = loglist.GetChannelTagC(event)
	if err != nil {
		return err
	}
	logs.channelid = event.ChannelID.String()
	logs.typelog = "create"
	logs.GuildID = event.GuildID.String()
	return add(logs)
}

func AddU(event *events.MessageUpdate) error {
	var logs logslist
	var err error
	logs.time = time.Now().Format("2006-01-02 15:04:05 MST")
	logs.message = loglist.GetContent(&event.Message)
	logs.messageid = event.Message.ID.String()
	logs.user = event.Message.Author.Tag()
	logs.userid = event.Message.Author.ID.String()
	logs.channel, err = loglist.GetChannelTagU(event)
	if err != nil {
		return err
	}
	logs.channelid = event.ChannelID.String()
	logs.typelog = "update"
	logs.GuildID = event.GuildID.String()
	return add(logs)
}
