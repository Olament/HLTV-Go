package model

type Team struct {
	Name string
	ID   *int //optional
}

type EventTeam struct {
	Team
	ReasonForParticipation string
	RankDuringEvent        *int
}
