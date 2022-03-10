package streamer

import (
	"bot/processor"
	"bot/ytdl"
	"net/url"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Item struct {
	Input string
	User  *gotgbot.User
	Video *ytdl.Video
}

var now Item

func Stream(input string, user *gotgbot.User) error {
	var video *ytdl.Video
	origInput := input
	_, err := url.ParseRequestURI(input)
	if err == nil {
		video, err = ytdl.Download(input)
		if err == nil {
			input = video.Url
		}
	}
	err = processor.Process(input)
	if err == nil {
		now.Input = origInput
		now.User = user
		now.Video = video
	}
	return err
}

func Now() (bool, Item) {
	return processor.Processing(), now
}
