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
	name       *string
	ign        *string
	image      *string
	age        *string
	country    *string
	team       *Team
	statistics Statistics
}
