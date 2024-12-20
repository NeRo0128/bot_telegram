package handlers

import (
	"bot_telegram/services"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

// NavigateCallback function handles navigation commands from the user.
// It takes a pointer to a Bot instance and a Context as parameters.
func NavigateCallback(b *gotgbot.Bot, ctx *ext.Context) error {

	var (
		cmd     string // Command extracted from the callback data.
		msg     string // Message to be sent back to the user.
		page    int    // Current page number for navigation.
		details string // todo
	)

	// Obtain the page number and command from the callback data.
	data := ctx.CallbackQuery.Data
	_, err := fmt.Sscanf(data, "navigate:%d cmd:%s details:%s", &page, &cmd, &details)
	if err != nil {
		return fmt.Errorf("datos de devolución de llamada inválidos: %w", err)
	}

	if page <= 0 {
		return FloatAlertCallbackQuery(
			b,
			ctx,
			"Este es el inicio",
		)
	}

	// Determine the action based on the command.
	switch cmd {
	case searchAnimesAirText:
		msg, err = services.GetAnimesInAir(page) // Fetch the animes currently in air for the specified page.
	case searchAnimesGenres:
		msg, err = services.GetAnimesForGenres(details, page)
	default:
		return fmt.Errorf("EL comando no coincide")
	}

	if err != nil {
		return err
	}

	// Edit the previous message with the new content and navigation buttons.
	_, _, err = b.EditMessageText(
		msg,
		&gotgbot.EditMessageTextOpts{
			ChatId:    ctx.EffectiveChat.Id,
			MessageId: ctx.EffectiveMessage.MessageId,
			ParseMode: "Markdown",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: *navigateButtons(page, cmd, details),
			},
		},
	)
	if err != nil {
		log.Printf("No se pudo editar el mensaje anterior: %v", err)
	}

	return nil
}

// DeleteMessageCallback function deletes the message associated with the callback query.
// It takes a pointer to a Bot instance and a Context as parameters.
func DeleteMessageCallback(b *gotgbot.Bot, ctx *ext.Context) error {

	// Delete the message in the chat that triggered the callback query.
	_, err := b.DeleteMessage(
		ctx.EffectiveChat.Id,
		ctx.CallbackQuery.Message.GetMessageId(),
		&gotgbot.DeleteMessageOpts{},
	)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar el mensaje: %w", err)
	}
	return nil
}

// FloatAlertCallbackQuery function sends an alert notification to the user.
// It takes a pointer to a Bot instance, a Context, and a message string as parameters.
func FloatAlertCallbackQuery(b *gotgbot.Bot, ctx *ext.Context, msg string) error {

	// Send an alert notification to the user based on the provided message.
	_, err := b.AnswerCallbackQuery(
		ctx.CallbackQuery.Id,
		&gotgbot.AnswerCallbackQueryOpts{
			Text:      msg,
			ShowAlert: false,
		},
	)
	if err != nil {
		return fmt.Errorf("error al enviar la notificación: %w", err)
	}

	return nil
}
