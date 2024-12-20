package handlers

import (
	"bot_telegram/services"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

//todo: documentation

// * Initials Commands

// Start function sends a welcome message to the user when they start the bot.
// It takes a pointer to a Bot instance and a Context as parameters.
func Start(b *gotgbot.Bot, ctx *ext.Context) error {

	// Reply to the user's effective message with a formatted welcome message.
	_, err := ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf(
			starMsg,
			b.User.Username,
		),
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// Help function sends a help message to the user, including inline buttons for additional actions.
// It takes a pointer to a Bot instance and a Context as parameters.
func Help(b *gotgbot.Bot, ctx *ext.Context) error {

	// Reply to the user's effective message with the help message and inline buttons.
	_, err := ctx.EffectiveMessage.Reply(
		b,
		helpMsg,
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: *helpButtons(),
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// * Messages

// searchAnimesInAir function searches for currently airing anime and sends the results to the user.
// It takes a pointer to a Bot instance, a Context, and a page number as parameters.
func searchAnimesInAir(b *gotgbot.Bot, ctx *ext.Context, page int) error {
	//todo: Refactorizar el nombre q esta en candela
	//	mejor formato a mensaje

	fmt.Printf("Searching Animes in Air\n")

	// Get the list of animes that are currently airing based on the provided page number.
	msg, err := services.GetAnimesInAir(page)
	if err != nil {
		return err
	}

	// Reply to the user's effective message with the list of animes.
	_, err = ctx.EffectiveMessage.Reply(
		b,
		msg,
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: *navigateButtons(page, searchAnimesAirText, "none"),
			},
		},
	)

	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	return nil
}

func searchAnimeForName(b *gotgbot.Bot, ctx *ext.Context, name string) error {

	//todo: mejor formato al mensaje
	// ver si es posible q me envie tb una img del anime

	fmt.Printf("Searching Anime for %s\n", name)
	msg, err := services.GetAnimeData(name)
	if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(
		b,
		msg,
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func searchAnimesForGenre(b *gotgbot.Bot, ctx *ext.Context, genre string, page int) error {
	fmt.Println("Searching Animes for Genre")
	msg, err := services.GetAnimesForGenres(genre, page)
	if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(
		b,
		msg,
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: *navigateButtons(page, searchAnimesGenres, genre),
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
func searchAnimesThisTemp(b *gotgbot.Bot, ctx *ext.Context) error {
	//todo: implement
	return nil
}

// * Unknown Command

// unknownCommand function handles unrecognized commands from the user.
// It takes a pointer to a Bot instance and a Context as parameters.
func unknownCommand(b *gotgbot.Bot, ctx *ext.Context) error {

	// Reply to the user's effective message with a message indicating the command is unknown.
	_, err := ctx.EffectiveMessage.Reply(
		b,
		unknownCommandMsg,
		&gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// * Buttons

// navigateButtons generates navigation buttons for paginated content.
// It accepts the current page number and a command string as parameters.
func navigateButtons(page int, cmd, details string) (buttons *[][]gotgbot.InlineKeyboardButton) {

	buttons = &[][]gotgbot.InlineKeyboardButton{
		{
			{ //Previous Button
				Text:         "⬅️ previous",
				CallbackData: fmt.Sprintf("navigate:%d cmd:%s details:%s", page-1, cmd, details),
			},
			{ // Delete Button
				Text:         "🗑️ delete",
				CallbackData: "delete_message",
			},
			{ // Next Button
				Text:         "➡️ next",
				CallbackData: fmt.Sprintf("navigate:%d cmd:%s details:%s", page+1, cmd, details),
			},
		},
	}
	return
}

// helpButtons generates help-related buttons.
func helpButtons() (buttons *[][]gotgbot.InlineKeyboardButton) {

	buttons = &[][]gotgbot.InlineKeyboardButton{
		{
			{
				Text:                         "search Anime for name",
				SwitchInlineQueryCurrentChat: &searchAnimeText,
			},
			{
				Text:                         "search Anime in Air",
				SwitchInlineQueryCurrentChat: &searchAnimesAirText,
			},
		},
		{
			{
				Text:                         "search Anime For Genres",
				SwitchInlineQueryCurrentChat: &searchAnimesGenres,
			},
		},
	}
	return
}
