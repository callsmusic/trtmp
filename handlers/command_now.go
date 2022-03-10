package handlers

import (
	"fmt"
	"html"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/helpers"
	"bot/streamer"
)

var commandNowHandler = handlers.NewCommand("now", commandNow)

func commandNow(b *gotgbot.Bot, ctx *ext.Context) error {
	var err error
	streaming, now := streamer.Now()
	if !streaming {
		_, err = ctx.Message.Reply(b, "Not streaming.", nil)
	} else {
		title := "Custom Input"
		if now.Video != nil {
			if now.Video.Title != "" {
				title = now.Video.Title
			}
		}
		title = fmt.Sprintf("<a href=\"%s\">%s</a>", now.Input, html.EscapeString(title))
		duration := ""
		if now.Video != nil {
			if now.Video.Duration != 0 {
				duration = fmt.Sprintf(" (%s)", now.Video.Duration.String())
			}
		}
		title = title + duration
		_, err = ctx.Message.Reply(b, fmt.Sprintf("%s is currently streaming %s...", helpers.MentionUser(now.User), title), &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	}
	return err
}
