package telegram

import (
	"log"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		// TODO: AddPage()
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:

	}
}

func isAddCmd(text string) bool {

}
