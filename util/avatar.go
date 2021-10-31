package util

import "net/url"

var (
	style = "open-peeps"
)

func GenerateAvatarUrl(username string) string {
	return "https://avatars.dicebear.com/api/" + style + "/" + url.QueryEscape(username) + ".svg?size=256"
}
