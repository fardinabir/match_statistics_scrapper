package models

type EspnStat struct {
	Date   string
	Opp    string
	Result string
	Min    string
	FG     string
	FT     string
	ThreeP string
	REB    string
	AST    string
	BLK    string
	STL    string
	PF     string
	TO     string
	PTS    string
}

type EuroBasket struct {
	Date        string
	Team        string
	AgainstTeam string
	Result      string
	Min         string
	Pts         string
	TwoFGP      string
	ThreeFGP    string
	FT          string
	RO          string
	RD          string
	RT          string
	AS          string
	PF          string
	BS          string
	ST          string
	TO          string
	RNK         string
}
