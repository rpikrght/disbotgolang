package event

import (
	"disbotgolang/dbfunc/jsonfunc/adminlist"

	"github.com/disgoorg/disgo/events"
)

func RoleUpdate(event *events.RoleUpdate) error {
	err := adminlist.UpdateRole(event)
	return err
}
