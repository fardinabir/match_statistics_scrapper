package models

import (
	"gorm.io/gorm"
)

type MatchStatUrl struct {
	gorm.Model
	PlayerName string `json:"playerName"`
	Url        string `json:"url"`
}

type ScrappedData struct {
	gorm.Model
	ScrappedHash string `json:"scrappedHash" gorm:"column:scrapped_hash;size:30;index"`
	//LastScrappingTime string `json:"lastScrappingTime"`
	//LastScrappedData  string `json:"lastScrappedData"`
}
