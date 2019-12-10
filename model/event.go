package model

import "hltv/enum"

type Event struct {
	Name string
	ID   *int
}

// TODO convert date to actual Date struct instead of string
type FullEvent struct {
	ID                int
	Name              string
	DateStart         string
	DateEnd           string
	PrizePool         string
	Teams             []EventTeam
	Location          Country
	PrizeDistribution []EventPrizeDistribution
	Formats           []EventFormat
	RelatedEvents     []Event
	MapPool           []enum.MapSlug
}

type EventFormat struct {
	Type        string
	Description string
}

type EventPrizeDistribution struct {
	Place        string
	Prize        string
	OtherPrize   string
	QualifiesFor *Event
	Team         *Team
}
