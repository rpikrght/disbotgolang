package logsfunc

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func add(logs logslist) error {
	_, err := os.Stat("./db/" + logs.GuildID)
	if os.IsNotExist(err) {
		Create(logs.GuildID)
	}

	logspath := "./db/" + logs.GuildID + "/message.db"

	logsdb, err := sql.Open("sqlite", logspath)
	if err != nil {
		log.Printf("sqlopenerr: %v", err)
	}
	defer logsdb.Close()

	statemen, err := logsdb.Prepare(`INSERT INTO messages(time,message,messageid,user,userid,channel,channelid,typelogs) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Printf("statemenerr: %v", err)
	}
	if _, err := statemen.Exec(logs.time, logs.message, logs.messageid, logs.user, logs.userid, logs.channel, logs.channelid, logs.typelog); err != nil {
		log.Printf("execerr: %v", err)
	}
	return nil
}
