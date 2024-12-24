package server

import (
	handlers2 "bot_telegram/handlers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

// LoadServer This bot demonstrates some example interactions with commands on telegram.
// It has a basic start command with a bot intro.
// It also has a source command, which sends the bot sourcecode, as a file.
func LoadServer() error {

	
	//Get token from the environment variable
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("TOKEN environment variable is empty")
	}

	// Create bot from environment value.
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		return fmt.Errorf("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", handlers2.Start))
	dispatcher.AddHandler(handlers.NewCommand("help", handlers2.Help))
	dispatcher.AddHandler(handlers.NewInlineQuery(inlinequery.QueryPrefix("ans_gnr"), handlers2.GenresInlineCallback))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("delete_message"), handlers2.DeleteMessageCallback))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("navigate:"), handlers2.NavigateCallback))
	dispatcher.AddHandler(handlers.NewMessage(message.Text, handlers2.AnalyzeMessages))

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to start polling: " + err.Error())
	}

	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()

	return nil
}
