package ytdl

import (
	"encoding/json"
	"errors"
	"os/exec"
	"time"
)

type Video struct {
	Duration time.Duration
	Title    string
	Url      string
}

var (
	ErrUnexpectedOutput = errors.New("unexpected output")

	args = []string{
		"--format",
		"best",
		"--geo-bypass",
		"--quiet",
		"--dump-json",
	}
)

func getArgs(input string) []string {
	return append(args, input)
}

func Download(input string) (*Video, error) {
	video := &Video{}
	var result interface{}
	output, err := exec.Command("youtube-dl", getArgs(input)...).Output()
	if err != nil {
		return video, err
	}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return video, err
	}
	data, ok := result.(map[string]interface{})
	if !ok {
		return video, ErrUnexpectedOutput
	}
	duration, ok := data["duration"].(float64)
	if !ok {
		return video, ErrUnexpectedOutput
	}
	video.Duration = time.Duration(duration) * time.Second
	url, ok := data["url"].(string)
	if !ok {
		return video, ErrUnexpectedOutput
	}
	video.Url = url
	title, ok := data["title"].(string)
	if !ok {
		return video, ErrUnexpectedOutput
	}
	video.Title = title
	return video, nil
}
