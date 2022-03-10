package main

import (
	"bot/handlers"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	token := os.Getenv("BOT_TOKEN")
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic(err)
	}
	updater := ext.NewUpdater(nil)
	handlers.Load(updater.Dispatcher)
	err = updater.StartPolling(
		bot, &ext.PollingOpts{
			DropPendingUpdates: true,
		},
	)
	if err != nil {
		panic(err)
	}
	updater.Idle()
}
