package models

import (
	"gorm.io/gorm"
)

type PlayersData struct {
	gorm.Model
	PlayerName string `json:"playerName"`
	Url        string `json:"url"`
}

type ScrappedData struct {
	gorm.Model
	Hash string `json:"scrappedHash" gorm:"column:hash;size:255;index"`
	Data string `json:"data" gorm:"column:data"`
}
