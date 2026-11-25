package models


type OpenMeteoResponse struct {
	Current struct {
		Temperature2m        float64 `json:"temperature_2m"`
		RelativeHumidity2m   float64 `json:"relative_humidity_2m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		WeatherCode          int     `json:"weather_code"`
		WindSpeed10m         float64 `json:"wind_speed_10m"`
	} `json:"current"`


	Daily struct {
		Time             []string  `json:"time"`
		WeatherCode      []int     `json:"weather_code"`
		Temperature2mMax []float64 `json:"temperature_2m_max"`
		Temperature2mMin []float64 `json:"temperature_2m_min"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
	} `json:"daily"`
}


type PrevisionJour struct {
	DateFr        string  
	Libelle       string 
	Icone         string  
	TempMin       float64
	TempMax       float64
	TempBarWidth  int   
	Precipitation float64
}


type PageMeteoData struct {
	Commune       Commune
	Actuelle      MeteoActuelle
	Previsions    []PrevisionJour
	CarteURL      string 
}

type MeteoActuelle struct {
	Temperature float64
	Ressentie   float64
	Humidite    float64
	Vent        float64
	Libelle     string
	Icone       string
}