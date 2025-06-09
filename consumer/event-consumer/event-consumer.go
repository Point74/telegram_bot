package event_consumer

import (
	"context"
	"log"
	"telegram_bot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	ctx := context.Background()

	for {
		gotEvents, err := c.fetcher.Fetch(ctx, c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer :%s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(ctx, gotEvents); err != nil {
			log.Printf("[ERR] handling events: %s", err)
		}
	}
}

func (c Consumer) handleEvents(ctx context.Context, events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(ctx, event); err != nil {
			log.Printf("[ERR] can't handle event :%s", err.Error())

			continue
		}
	}

	return nil
}
