package model

type Statistics struct {
	Rating            float32
	KillsPerRound     float32
	MapsPlayed        int
	DeathsPerRound    float32
	Headshots         float32
	RoundsContributed float32
}

type FullPlayer struct {
	ID           int
	Name         *string //optional
	Ign          string
	Image        *string //optional
	Age          *int    //optional
	Country      Country
	Team         Team
	Twitter      *string
	Facebook     *string
	Statistics   Statistics
	Achievements []Achievement
}
