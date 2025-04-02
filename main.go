package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"disbotgolang/event"
	"disbotgolang/token"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
)

func main() {

	token, err := token.String()
	if err != nil {
		log.Printf("tokenerr: %v", err)
	}

	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentGuildMembers,
				gateway.IntentsDirectMessage,
				gateway.IntentMessageContent,
			),
		),
		bot.WithEventListenerFunc(event.MessageCreate),
		bot.WithEventListenerFunc(event.MessageUpdate),
		bot.WithEventListenerFunc(event.BotAdd),
		bot.WithEventListenerFunc(event.RoleUpdate),
	)
	if err != nil {
		log.Printf("clienterr: %v", err)
		return
	}
	ctx := context.Background()
	if err := client.OpenGateway(ctx); err != nil {
		log.Printf("conecterr: %v", err)
		return
	}
	log.Printf("boton")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGABRT, os.Interrupt)
	<-sc
	log.Printf("botoff")
}
