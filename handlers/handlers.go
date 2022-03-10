package handlers

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

func Load(dp *ext.Dispatcher) {
	dp.AddHandler(commandNowHandler)
	dp.AddHandler(commandStopHandler)
	dp.AddHandler(commandStreamHandler)
}
