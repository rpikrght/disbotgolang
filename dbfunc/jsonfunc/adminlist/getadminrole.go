package adminlist

import (
	"encoding/json"
	"log"
	"os"

	"github.com/disgoorg/disgo/events"
)

func getadminrole(event *events.RoleUpdate) error {
	rolesdb, err := os.ReadFile("./db/" + event.GuildID.String() + "/adminrolelist.json")
	if err != nil {
		log.Printf("roleupdateerr: %v", err)
		return err
	}

	err = json.Unmarshal(rolesdb, &roles)
	if err != nil {
		log.Printf("unmarshalerr: %v", err)
	}
	return nil
}
