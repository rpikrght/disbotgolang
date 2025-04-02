package adminlist

import (
	"encoding/json"
	"log"
	"os"

	"github.com/disgoorg/disgo/events"
)

var roles []listrole

func UpdateRole(event *events.RoleUpdate) error {
	rolePath := "./db/" + event.GuildID.String() + "/adminrolelist.json"
	if event.Role.Permissions&402653190 > 0 {
		if err := addrole(event); err != nil {
			return err
		}
	} else {
		if err := deleterole(event); err != nil {
			return err
		}
	}
	newRoleDate, err := json.MarshalIndent(roles, "", " ")
	if err != nil {
		log.Printf("newroledateerr: %v", err)
		return err
	}
	if err := os.WriteFile(rolePath, newRoleDate, 0644); err != nil {
		log.Printf("writenewerr: %v", err)
		return err
	}
	return nil
}
