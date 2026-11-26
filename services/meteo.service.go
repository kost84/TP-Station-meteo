package services
import (
	"strconv"
	"strings"
	"meteo/api"
	"meteo/model"
)
func SearchPokemon(name string) ([]models.Commune, error) {

	toutemeteo, err := api.GetCharacter()
	if err != nil {
		return nil, err
	}
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return toutemeteo, nil
	}
	var results []models.Commune
	for _, p := range toutemeteo {
		if strings.Contains(strings.ToLower(p.Name.Fr), name) {
			results = append(results, p)
		}
	}
	return results, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetPokemon(idParam string) (models.Commune, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetMeteo(id)
}