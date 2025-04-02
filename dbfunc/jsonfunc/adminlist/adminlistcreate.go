package adminlist

import (
	"encoding/json"
	"log"
	"os"

	"github.com/disgoorg/disgo/events"
)

type list struct {
	ID string `json:"id"`
}

func Create(event *events.GuildJoin) error {
	owner := list{ID: event.Guild.OwnerID.String()}
	adminPath := "./db/" + event.GuildID.String() + "/adminlist.json"
	adminfile, err := os.Create(adminPath)
	if err != nil {
		log.Printf("createadminerr: %v", err)
		return err
	}
	defer adminfile.Close()

	jsonadmin, err := json.MarshalIndent(owner, "", " ")
	if err != nil {
		log.Printf("marshalerr: %v", err)
		return err
	}
	if err := os.WriteFile(adminPath, jsonadmin, 0644); err != nil {
		log.Printf("writeerr: %v", err)
		return err
	}
	return nil
}
