package model

type Statistics struct {
	Rating            float32 `json:"rating"`
	KillsPerRound     float32 `json:"kills_per_round"`
	MapsPlayed        int     `json:"maps_played"`
	DeathsPerRound    float32 `json:"deaths_per_round"`
	Headshots         float32 `json:"headshots"`
	RoundsContributed float32 `json:"rounds_contributed"`
}

type FullPlayer struct {
	ID           int           `json:"id"`
	Name         *string       `json:"name"`
	Ign          string        `json:"ign"`
	Image        *string       `json:"image"`
	Age          *int          `json:"age"`
	Country      Country       `json:"country"`
	Team         Team          `json:"team"`
	Twitter      *string       `json:"twitter"`
	Twitch       *string       `json:"twitch"`
	Facebook     *string       `json:"facebook"`
	Statistics   Statistics    `json:"statistics"`
	Achievements []Achievement `json:"achievements"`
}
