package model

import (
	"hltv/enum"
	"time"
)

type Event struct {
	Name string
	ID   *int
}

type FullEvent struct {
	ID                int
	Name              string
	DateStart         time.Time
	DateEnd           time.Time
	PrizePool         string
	Teams             []EventTeam
	Location          Country
	PrizeDistribution []EventPrizeDistribution
	Formats           []EventFormat
	RelatedEvents     []Event
	MapPool           []enum.MapSlug
}

type EventResults struct {
	Month time.Month
	Events []SimpleEvent
}
type SimpleEvent struct {
	ID            int
	Name          string
	DateStart     *time.Time
	DateEnd       *time.Time
	PrizePool     string
	NumberOfTeams *int
	Location      Country
	Type          enum.EventType
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
