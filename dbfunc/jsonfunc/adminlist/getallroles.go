package adminlist

import (
	"log"

	"github.com/disgoorg/disgo/events"
)

func getallroles(event *events.GuildJoin) error {
	allroles, err := event.Client().Rest().GetRoles(event.GuildID)
	if err != nil {
		log.Printf("gerroleserr: %v", err)
		return err
	}
	for _, role := range allroles {
		if role.Permissions&402653190 > 0 {
			moderrole = append(moderrole, listrole{IDrole: role.ID.String()})
		}
	}
	return nil
}
