package helpers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func MentionUser(user *gotgbot.User) string {
	name := user.FirstName
	if user.LastName != "" {
		name += " " + user.LastName
	}
	return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", user.Id, name)
}
