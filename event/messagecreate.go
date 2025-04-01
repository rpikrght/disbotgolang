package event

import (
	"disbotgolang/dbfunc/logsfunc"
	"log"

	"github.com/disgoorg/disgo/events"
)

func MessageCreate(event *events.MessageCreate) {
	if err := logsfunc.AddToLogsListC(event); err != nil {
		log.Printf("logserr: %v", err)
		return
	}
}
