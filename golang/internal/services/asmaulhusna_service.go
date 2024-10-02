package services

import (
	"encoding/json"
	"io"
	"os"

	"github.com/haikelz/asmaul-husna-api/internal/configs"
	"github.com/haikelz/asmaul-husna-api/internal/entities"
)

func GetAsmaulHusna() ([]*entities.AsmaulHusna, error) {
	var asmaulHusna []*entities.AsmaulHusna

	if err := configs.Gorm().Find(&asmaulHusna).Error; err != nil {
		return nil, err
	}

	return asmaulHusna, nil
}

func InsertLoopAsmaulHusna() ([]entities.AsmaulHusna, error) {
	// search where the file
	jsonFile, err := os.Open("internal/configs/asmaulhusna.json")

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	// Read the JSON file and convert the data to array byte
	fileData, _ := io.ReadAll(jsonFile)

	var asmaulHusna []entities.AsmaulHusna

	errLoad := json.Unmarshal(fileData, &asmaulHusna)

	if errLoad != nil {
		return nil, errLoad
	}

		for i := range asmaulHusna {
			data := &entities.AsmaulHusna{Urutan: asmaulHusna[i].Urutan, Latin: asmaulHusna[i].Latin, Arab: asmaulHusna[i].Arab, Arti: asmaulHusna[i].Arti}

			if dbErr := configs.Gorm().Create(&data).Error; dbErr != nil {
			}
		}

	return asmaulHusna, nil
}
