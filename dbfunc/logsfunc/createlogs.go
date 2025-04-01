package logsfunc

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func CreateLogs(GuildID string) error {
	logsdir := "./db/" + GuildID
	if err := os.MkdirAll(logsdir, 0755); err != nil {
		log.Printf("logsdirerr: %v", err)
		return err
	}

	logspath := filepath.Join(logsdir, "message.db")

	logs, err := sql.Open("sqlite", logspath)
	if err != nil {
		log.Printf("openlogserr: %v", err)
	}
	defer logs.Close()

	if _, err := logs.Exec(`CREATE TABLE IF NOT EXISTS messages (
			time TEXT NOT NULL,
			message TEXT,
			messageid TEXT NOT NULL,
			user TEXT NOT NULL,
			userid TEXT NOT NULL,
			channel TEXT NOT NULL,
			channelid TEXT 	NOT NULL,
			typelogs TEXT NOT NULL
		);
	`); err != nil {
		log.Printf("createtableerr: %v", err)
		return err
	}
	log.Printf("table for guild %q create: %v", GuildID, err)
	return nil
}
