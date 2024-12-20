package services

// todo: documentation
//
//	ver bien el tema de los campos q resivo
const (
	queryGetAnimeByName = `
	{
        Media(search: "%s") {
            title {
                romaji
                english
                native
            }
            description
            episodes
            genres
        }
    }`

	queryGetAnimesInAir = `
	{
        Page (perPage: 5, page: %d) {
            media(status: RELEASING, format: TV) {
                title {
                    romaji
                    native
                }
                genres
            }
        }
    }`
	queryGetGenres = `
        query {
            Genres
        }
    `
	queryGetAnimesForGenres = `
	{
		Page(page: %d, perPage: 5) {
			media(genre: "%s") {
				title {
					native
					romaji
				},
				startDate {
					year
			  	}
			}
		}
	}`
)
