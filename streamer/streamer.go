package streamer

import (
	"bot/processor"
	"bot/ytdl"
	"fmt"
	"net/url"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Item struct {
	Input string
	User  *gotgbot.User
	Video *ytdl.Video
}

var now Item

func Stream(b *gotgbot.Bot, input string, user *gotgbot.User) error {
	var video *ytdl.Video
	origInput := input
	_, err := url.ParseRequestURI(input)
	if err == nil {
		video, err = ytdl.Download(input)
		if err == nil {
			input = video.Url
		}
	}
	errc := make(chan error)
	go func() {
		for {
			err, ok := <-errc
			if !ok {
				break
			}
			b.SendMessage(user.Id, fmt.Sprintf("Failed to process: %s", err.Error()), nil)
		}
	}()
	err = processor.Process(input, errc)
	if err == nil {
		now.Input = origInput
		now.User = user
		now.Video = video
	} else {
		close(errc)
	}
	return err
}

func Now() (bool, Item) {
	return processor.Processing(), now
}
