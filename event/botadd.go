package event

import (
	"os"

	"disbotgolang/dbfunc/jsonfunc/adminlist"
	"disbotgolang/dbfunc/logsfunc"

	"github.com/disgoorg/disgo/events"
)

func BotAdd(event *events.GuildJoin) {
	_, err := os.Stat("./db/" + event.GuildID.String())
	if os.IsNotExist(err) {
		if err := logsfunc.Create(event.GuildID.String()); err != nil {
			return
		}
	}
	_, err = os.Stat("./db/" + event.GuildID.String() + "/adminlist.json")
	if os.IsNotExist(err) {
		if err := adminlist.Create(event); err != nil {
			return
		}
	}
	_, err = os.Stat("./db/" + event.GuildID.String() + "/adminrolelist.json")
	if os.IsNotExist(err) {
		if err := adminlist.CreateRole(event); err != nil {
			return
		}
	}
}
