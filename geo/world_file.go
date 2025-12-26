package geo

import (
	"encoding/json"
	"os"
)

type SphereData struct {
	CenterX      float64 `json:"cx"`
	CenterY      float64 `json:"cy"`
	CenterZ      float64 `json:"cz"`
	Radius       float64 `json:"r"`
	MaterialType string  `json:"mat_type"` // "lambertian", "metal", "dielectric", "light"
	AlbedoR      float64 `json:"ar"`
	AlbedoG      float64 `json:"ag"`
	AlbedoB      float64 `json:"ab"`
	Fuzz         float64 `json:"fuzz,omitempty"`
}

func SaveWorld(filename string, data []SphereData) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(data)
}

func LoadWorld(filename string) ([]SphereData, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var data []SphereData
	err = json.NewDecoder(f).Decode(&data)
	return data, err
}
