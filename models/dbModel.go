package models

import "gorm.io/gorm"

type MatchStatUrl struct {
	gorm.Model
	PlayerName    string `json:"playerName"`
	Url           string `json:"url"`
	ScrappingTime string `json:"scrappingTime"`
}
