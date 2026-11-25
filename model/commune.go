package models


type Commune struct {
	Nom             string   `json:"nom"`
	Code            string   `json:"code"` 
	CodeDepartement string   `json:"codeDepartement"`
	CodesPostaux    []string `json:"codesPostaux"`
	Population      int      `json:"population"`
	Centre          Point    `json:"centre"`
}


type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"` 
}


func (c Commune) Longitude() float64 {
	if len(c.Centre.Coordinates) >= 2 {
		return c.Centre.Coordinates[0]
	}
	return 0.0
}

func (c Commune) Latitude() float64 {
	if len(c.Centre.Coordinates) >= 2 {
		return c.Centre.Coordinates[1]
	}
	return 0.0
}