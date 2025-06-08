package telegram

import (
	"context"
	"log"
	"net/url"
	"strings"
	"telegram_bot/clients/telegram"
	"telegram_bot/lib/e"
	"telegram_bot/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		return p.savePage(ctx, chatID, text, username)
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:

	}
}

func (p *Processor) savePage(ctx context.Context, chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()

	sendMsg := NewMessageSender(chatID, p.tg)

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExist(ctx, page)
	if err != nil {
		return err
	}

	if isExists {
		return sendMsg(msgAlreadyExist)
	}

	if err := p.storage.Save(ctx, page); err != nil {
		return err
	}

	if err := sendMsg(msgSaved); err != nil {
		return err
	}

	return nil
}

func NewMessageSender(chatID int, tg *telegram.Client) func(string) error {
	return func(msg string) error {
		return tg.SendMessage(chatID, msg)
	}
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
