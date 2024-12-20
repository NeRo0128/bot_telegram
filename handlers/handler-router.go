package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"regexp"
)

//todo: dar un formato mas legible a los comando
//	documentation

var re = regexp.MustCompile(`^@anime_list_666_bot\s+(\w+)(?:\s+(.+))?$`)

// Routes
// search Anime name : an.name
// search Animes genre : ans.gnr
// search Anime in air : ans.air
// search anime this temp : ans.temp
// todo: refactorizar nombre de los comandos de mensajes
var (
	searchAnimeText     = "an_name"
	searchAnimesAirText = "ans_air"
	searchAnimesGenres  = "ans_gnr"
)

func AnalyzeMessages(b *gotgbot.Bot, ctx *ext.Context) error {

	if ctx.Message == nil {
		return nil
	}
	text := ctx.Message.Text

	if re.MatchString(text) {

		matches := re.FindStringSubmatch(text)
		command := matches[1]
		argument := ""
		if len(matches) > 2 {
			argument = matches[2]
		}

		switch command {
		case searchAnimeText:
			return searchAnimeForName(b, ctx, argument)
		case searchAnimesAirText:
			return searchAnimesInAir(b, ctx, 1)
		case searchAnimesGenres:
			return searchAnimesForGenre(b, ctx, argument, 1)
		}
	}

	return unknownCommand(b, ctx)
}
