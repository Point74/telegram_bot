package main

import (
	"context"
	"flag"
	"log"
	tgClient "telegram_bot/clients/telegram"
	"telegram_bot/consumer/event-consumer"
	"telegram_bot/events/telegram"
	"telegram_bot/storage/sqlite"
)

const (
	bathSize          = 100
	sqliteStoragePath = "data/sqlite/storage.db"
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(Host(), mustToken()),
		s,
	)

	log.Print("start service")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped ", err)
	}
}

func mustToken() string {
	token := flag.String(
		`tg-bot-token`,
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

func Host() string {
	host := flag.String(
		"",
		`api.telegram.org`,
		"host for access to telegram bot",
	)

	return *host
}
