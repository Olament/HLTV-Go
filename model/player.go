package model

type Player struct {
	Name string
	ID   *int
}

type PlayerRanking struct {
	Player
	Rating float32
}

type FullPlayer struct {
	Player
	Ign          string
	Image        *string
	Age          *int
	Country      Country
	Team         Team
	Twitter      *string
	Twitch       *string
	Facebook     *string
	Statistics   Statistics
	Achievements []Achievement
}

type FullPlayerStats struct {
	Player
	Ign        *string
	Image      *string
	Age        *int
	Country    *Country
	Team       *Team
	Statistics Statistics
}

type Statistics struct {
	Kills                   int
	Headshots               float32
	Death                   int
	KDRatio                 float32
	DamgePerRound           float32
	GrenadeDamge            float32
	MapsPlayed              int
	RoundsPlayed            int
	KillsPerRound           float32
	AssistsPerRound         float32
	DeathsPerRound          float32
	SavedByTeammatePerRound float32
	SavedTeammatesPerRound  float32
	Rating                  float32
	RoundsContributed       float32
}
