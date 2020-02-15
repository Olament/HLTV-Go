package hltv

import (
	"testing"
	"time"
)

func TestUpcomingMatches(t *testing.T) {
	h := HLTV{
		Url:       "https://www.hltv.org",
		StaticURL: "",
	}

	matches, err := h.GetUpcomingMatches(UpcomingMatchesQuery{})
	if err != nil {
		t.Errorf("GetUpcomingMatches returned an error %v", err)
		return
	}
	if len(matches) == 0 {
		t.Errorf("GetUpcomingMatches returned no matches")
		return
	}

	m := matches[0]
	if *m.ID == 0 {
		t.Errorf("Matches have no ID")
	}
	if m.Date == time.Unix(0, 0) {
		t.Errorf("Match has wrong date")
	}
	if m.Team1.Name == "" {
		t.Errorf("Teams have no name")
	}
	if *m.Team1.ID == 0 {
		t.Errorf("Teams have no ID")
	}
	if m.Event.Name == "" {
		t.Errorf("Events have no Name")
	}
	if *m.Event.ID == 0 {
		t.Errorf("Events have no ID")
	}
}
