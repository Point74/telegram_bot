package main

import (
	"flag"
	"log"
)

func main() {
	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
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

	flag.Parse()

	return *host
}
