package event

import (
	"disbotgolang/dbfunc/logsfunc"
	"log"

	"github.com/disgoorg/disgo/events"
)

func MessageUpdate(event *events.MessageUpdate) {
	if err := logsfunc.AddU(event); err != nil {
		log.Printf("logserr: %v", err)
		return
	}
}
