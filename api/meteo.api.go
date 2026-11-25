package api

import (
	"encoding/json"
	"fmt"
    "meteo/model"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,

}

func RechercheComunne(nom string) ([]models.Commune, error){

	endpoint := fmt.Sprintf("https://geo.api.gouv.fr/communes?nom=%s&fields=nom,code,codesPostaux,codeDepartement,centre,population")

	resp , err := httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("erreur réseau API Gouv: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Gouv a répondu avec le statut %d", resp.StatusCode)
	}
	var communes []models.Commune
	if err := json.NewDecoder(resp.Body).Decode(&communes); err != nil {
		return nil, fmt.Errorf("erreur décodage JSON: %w", err)
	}
	return communes, nil
}	
	
func RecupererCommuneParCode(codeINSEE string) (*models.Commune, error) {
	endpoint := fmt.Sprintf(
		"https://geo.api.gouv.fr/communes/%s?fields=nom,code,codesPostaux,codeDepartement,centre,population",
		url.PathEscape(codeINSEE),
	)
	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("erreur réseau API Gouv: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil // Commune non trouvée
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Gouv a répondu avec le statut %d", resp.StatusCode)
	}
	var c models.Commune
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, fmt.Errorf("erreur décodage JSON: %w", err)
	}
	return &c, nil
}