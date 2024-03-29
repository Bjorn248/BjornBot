package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-chat-bot/bot"
	"github.com/nicklaw5/helix/v2"
)

func uptime(command *bot.Cmd) (msg string, err error) {

	client, err := helix.NewClient(&helix.Options{
		ClientID:     os.Getenv("GO_TWITCH_CLIENTID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
	})
	if err != nil {
		msg = "Error creating twitch api client"
		return
	}

	authResp, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		msg = "Error getting twitch access token"
		return
	}

	// Set the access token on the client
	client.SetAppAccessToken(authResp.Data.AccessToken)

	resp, err := client.GetStreams(&helix.StreamsParams{
		First:      1,
		UserLogins: []string{"bjorn_248"},
	})
	if err != nil {
		msg = "Error retrieving stream information from twitch API"
		return
	}

	if len(resp.Data.Streams) == 0 {
		msg = "Stream is currently offline"
		return
	}

	streamStartTime := resp.Data.Streams[0].StartedAt

	streamUptime := time.Since(streamStartTime)

	// Remove millseconds from Duration, we don't need that level of precision
	// This doesn't round, it just flat out removes everything after the dot
	uptimeString := strings.Split(streamUptime.String(), ".")[0]

	msg = fmt.Sprintf("Stream Uptime: %vs\n", uptimeString)

	return
}

func init() {
	bot.RegisterCommand(
		"uptime",
		"Returns the Channel Uptime",
		"",
		uptime)
}
