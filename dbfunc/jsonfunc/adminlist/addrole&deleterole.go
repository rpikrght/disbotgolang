package adminlist

import (
	"slices"

	"github.com/disgoorg/disgo/events"
)

func addrole(event *events.RoleUpdate) error {
	getadminrole(event)
	for i := range roles {
		if roles[i].IDrole == event.RoleID.String() {
			return nil
		}
	}
	roles = append(roles, listrole{IDrole: event.RoleID.String()})
	return nil
}

func deleterole(event *events.RoleUpdate) error {
	getadminrole(event)
	for i := range roles {
		if roles[i].IDrole == event.RoleID.String() {
			roles = slices.Delete(roles, i, i+1)
		}
	}
	return nil
}
