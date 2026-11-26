package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"meteo/helpers"
	"meteo/model"
	"meteo/services"
)


type ListPage struct {
	Search   string         
	Erreur   string         
	Communes []models.Commune 
}


func Accueil(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, http.StatusOK, "accueil", ListPage{})
}


func ListVilles(w http.ResponseWriter, r *http.Request) {

	search := strings.TrimSpace(r.URL.Query().Get("q"))


	if len(search) == 0 {
		helpers.Render(w, http.StatusBadRequest, "accueil", ListPage{
			Search: search,
			Erreur: "Veuillez saisir un nom de commune.",
		})
		return
	}
	if len(search) < 2 {
		helpers.Render(w, http.StatusBadRequest, "accueil", ListPage{
			Search: search,
			Erreur: "Le nom doit comporter au moins 2 caractères.",
		})
		return
	}


	communes, err := services.SearchCommunes(search)
	if err != nil {
		renderAPIError(w, err)
		return
	}


	if len(communes) == 0 {
		helpers.RenderError(w, http.StatusNotFound, "Aucune commune trouvée pour cette recherche.")
		return
	}

	helpers.Render(w, http.StatusOK, "ville", ListPage{
		Search:   search,
		Communes: communes,
	})
}


func ShowMeteo(w http.ResponseWriter, r *http.Request) {

	code := strings.TrimSpace(r.URL.Query().Get("code"))
	if code == "" {
		helpers.RenderError(w, http.StatusBadRequest, "Code INSEE de la commune manquant.")
		return
	}

	meteoData, err := services.GetMeteo(code)
	if err != nil {
		renderAPIError(w, err)
		return
	}

	helpers.Render(w, http.StatusOK, "meteo", meteoData)
}


func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

func renderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound,
			"Commune introuvable.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest,
			"Paramètre de requête invalide.")
	case errors.Is(err, models.ErrUnavailable):
		log.Println("API distante indisponible :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable,
			"Service météo momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError,
			"Une erreur interne est survenue.")
	}
}