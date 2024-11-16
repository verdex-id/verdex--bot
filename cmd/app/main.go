package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/verdex-id/verdex-bot/injector"
	"github.com/verdex-id/verdex-bot/internal/infrastructure/config"
)

func main() {
	config.Load()

	discordDelivery := injector.InjectDeliveryDiscord()
	discordDelivery.RegisterEvents()

	err := discordDelivery.Session.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Verdex Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discordDelivery.Session.Close()
}
