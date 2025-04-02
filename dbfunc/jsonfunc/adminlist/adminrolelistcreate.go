package adminlist

import (
	"encoding/json"
	"log"
	"os"

	"github.com/disgoorg/disgo/events"
)

type listrole struct {
	IDrole string `json:"idrole"`
}

var moderrole []listrole

func CreateRole(event *events.GuildJoin) error {
	adminPath := "./db/" + event.GuildID.String() + "/adminrolelist.json"
	adminfile, err := os.Create(adminPath)
	if err != nil {
		log.Printf("createadminerr: %v", err)
		return err
	}
	defer adminfile.Close()

	if err := getallroles(event); err != nil {
		log.Printf("getroleserr: %v", err)
	}

	jsonadmin, err := json.MarshalIndent(moderrole, "", " ")
	if err != nil {
		log.Printf("marshalerr: %v", err)
		return err
	}
	if err := os.WriteFile(adminPath, jsonadmin, 0644); err != nil {
		log.Printf("writeerr: %v", err)
		return err
	}
	moderrole = moderrole[:0]
	return nil
}
