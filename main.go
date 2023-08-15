package main

import (
	"log"
	"os"
	"strings"

	"github.com/Bjorn248/go-chat-bot-bot/irc"
)

func main() {

	if os.Getenv("TWITCH_OAUTH_TOKEN") == "" {
		log.Fatal("TWITCH_OAUTH_TOKEN not set")
	}
	if os.Getenv("TWITCH_CHANNEL_NAME") == "" {
		log.Fatal("TWITCH_CHANNEL_NAME not set")
	}

	irc.Run(&irc.Config{
		Server:   "irc.chat.twitch.tv:6697",
		Channels: strings.Split(os.Getenv("TWITCH_CHANNEL_NAME"), ","),
		User:     "BjornTwitchBot",
		Nick:     "bjorntwitchbot",
		Password: os.Getenv("TWITCH_OAUTH_TOKEN"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != ""})
}
