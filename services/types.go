package services

import "fmt"

type ResponseMsg interface {
	Respond() string
}
type ResponseInline interface {
	Respond() []string
}

type AnimeResponse struct {
	Data struct {
		Media struct {
			Title       AnimeTitle `json:"title"`
			Description string     `json:"description"`
			Episodes    int        `json:"episodes"`
			Genres      []string   `json:"genres"`
		} `json:"Media"`
	} `json:"data"`
}

func (a AnimeResponse) Respond() string {

	msg := `
	*Title* 
		_Romaji_: %s
		_Native_: %s
	*Description*: 
		%s
	*Episodes*: %d
	*Genres*: 
	%v
	`

	res := fmt.Sprintf(
		msg,
		a.Data.Media.Title.Romaji,
		a.Data.Media.Title.Native,
		a.Data.Media.Description,
		a.Data.Media.Episodes,
		a.Data.Media.Genres,
	)

	return res
}

type AnimesInAirResponse struct {
	Data struct {
		Page struct {
			Media []struct {
				Title  AnimeTitle `json:"title"`
				Genres []string   `json:"genres"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}

func (a AnimesInAirResponse) Respond() string {
	msg := `
	*Title:* 
	  - *Romaji*: _%s_
	  - *Native*: _%s_
	*Genres*: 
	%v
	`
	var res string
	for _, anime := range a.Data.Page.Media {
		res += fmt.Sprintf(
			msg,
			anime.Title.Romaji,
			anime.Title.Native,
			anime.Genres,
		)
	}
	return res
}

type AnimesForGenresResponse struct {
	Data struct {
		Page struct {
			Media []struct {
				Title     AnimeTitle     `json:"title"`
				StartDate AnimeStartDate `json:"startDate"`
			} `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}

func (a AnimesForGenresResponse) Respond() string {
	msg := `
	*Title:* 
	  - *Romaji*: _%s_
	  - *Native*: _%s_
	*Start Year*: %d
	`
	var res string
	for _, anime := range a.Data.Page.Media {
		res += fmt.Sprintf(
			msg,
			anime.Title.Romaji,
			anime.Title.Native,
			anime.StartDate.Year,
		)
	}

	return res
}

type GenreResponseInline struct {
	Data struct {
		Genres []string `json:"Genres"`
	} `json:"data"`
}

func (a GenreResponseInline) Respond() []string {
	return a.Data.Genres
}

type AnimeTitle struct {
	Romaji string `json:"romaji"`
	Native string `json:"native"`
}
type AnimeStartDate struct {
	Year int `json:"year"`
}
