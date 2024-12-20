package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// * Services

// GetAnimeData obtiene los datos de un anime específico por nombre
func GetAnimeData(animeName string) (string, error) {
	query := fmt.Sprintf(queryGetAnimeByName, animeName)
	data, err := fetchData(query)
	if err != nil {
		return "", err
	}

	var animeResponse AnimeResponse
	if err := json.Unmarshal(data, &animeResponse); err != nil {
		return "", err
	}
	return animeResponse.Respond(), nil
}

// GetAnimesInAir retorna los animes que están actualmente en emisión
func GetAnimesInAir(page int) (string, error) {
	data, err := fetchData(fmt.Sprintf(queryGetAnimesInAir, page))
	if err != nil {
		return "", err
	}

	var animesInAirResponse AnimesInAirResponse
	if err := json.Unmarshal(data, &animesInAirResponse); err != nil {
		return "", err
	}

	return animesInAirResponse.Respond(), nil
}

func GetGenres() ([]string, error) {

	query := fmt.Sprintf(queryGetGenres)
	data, err := fetchData(query)
	if err != nil {
		return nil, err
	}

	var genreResponse GenreResponseInline
	if err := json.Unmarshal(data, &genreResponse); err != nil {
		return nil, fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return genreResponse.Respond(), nil
}

func GetAnimesForGenres(genre string, page int) (string, error) {

	query := fmt.Sprintf(queryGetAnimesForGenres, page, genre)
	data, err := fetchData(query)
	if err != nil {
		return "", err
	}

	var animesForGenresResponse AnimesForGenresResponse
	if err := json.Unmarshal(data, &animesForGenresResponse); err != nil {
		return "", fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return animesForGenresResponse.Respond(), nil

}

// * Private Functions

// fetchData realiza una solicitud POST a la API de AniList y decodifica la respuesta
func fetchData(query string) ([]byte, error) {
	payload := map[string]string{"query": query}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://graphql.anilist.co/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error al llamar a la API de AniList: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error al close body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("respuesta no exitosa de la API: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta de la API: %v", err)
	}

	return body, nil
}
