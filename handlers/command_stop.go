package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/processor"
)

var commandStopHandler = handlers.NewCommand("stop", commandStop)

func commandStop(b *gotgbot.Bot, ctx *ext.Context) error {
	stopped, err := processor.Stop()
	if err != nil {
		return err
	}
	if stopped {
		_, err = ctx.Message.Reply(b, "Stopped.", nil)
	} else {
		_, err = ctx.Message.Reply(b, "Not streaming.", nil)
	}
	return err
}
