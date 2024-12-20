package services

import "fmt"

func extractAnimeInfo(result map[string]interface{}, key string) (string, error) {
	if data, ok := result["data"].(map[string]interface{}); ok {
		if media, ok := data[key].(map[string]interface{}); ok {
			return formatAnimeData(media), nil
		}
	}
	return "Anime no encontrado.", nil
}

// extractAnimes extrae la información de una lista de animes
func extractAnimes(mediaList []interface{}) string {
	var resultString string
	for _, media := range mediaList {
		if mediaMap, ok := media.(map[string]interface{}); ok {
			resultString += formatAnimeData(mediaMap) + "\n"
		}
	}
	return resultString
}

// formatAnimeData formatea los datos del anime para la salida
func formatAnimeData(media map[string]interface{}) string {
	title := media["title"].(map[string]interface{})
	return fmt.Sprintf("Title (Romaji): %v\nTitle (Native): %v\nGenres: %v\n",
		title["romaji"],
		title["native"],
		media["genres"])
}
