package handlers

import (
	"bot_telegram/services"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strconv"
)

func GenresInlineCallback(b *gotgbot.Bot, ctx *ext.Context) error {

	query := ctx.InlineQuery.Query
	fmt.Println(query)

	genres, err := services.GetGenres()
	if err != nil {
		return err
	}

	var results []gotgbot.InlineQueryResult

	for i, genre := range genres {
		msg := fmt.Sprintf("@anime_list_666_bot ans_gnr %s", genre)
		results = append(results, &gotgbot.InlineQueryResultArticle{
			Id:    strconv.Itoa(i),
			Title: genre,
			InputMessageContent: gotgbot.InputTextMessageContent{
				MessageText: msg,
			},
		})
	}

	// Envía la respuesta a la consulta en línea
	_, err = b.AnswerInlineQuery(ctx.InlineQuery.Id, results, nil)
	if err != nil {
		return fmt.Errorf("failed to answer inline query: %w", err)
	}
	return nil
}
