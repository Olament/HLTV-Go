package model

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

type FullPlayerStats struct {
	Name       *string
	Ign        *string
	Image      *string
	Age        *string
	Country    *string
	Team       *Team
	Statistics Statistics
}
