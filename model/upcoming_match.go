package model

import "time"

type UpcomingMatch struct {
	ID     *int
	Team1  Team
	Team2  Team
	Date   time.Time
	Event  Event
	Format string
}
