package hltv

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"

	"github.com/Olament/HLTV-Go/model"
	"github.com/Olament/HLTV-Go/utils"
)

type UpcomingMatchesQuery struct {
	Team []int
}

func (h *HLTV) GetUpcomingMatches(q UpcomingMatchesQuery) (upcomingMatches []*model.UpcomingMatch, err error) {
	// Build query string parameters
	queryString, _ := query.Values(q)
	res, err := utils.GetQuery(h.Url + "/matches/?" + queryString.Encode())
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".upcomingMatch").Each(func(i int, selection *goquery.Selection) {
		matchHref, _ := selection.Find("a.match").First().Attr("href")
		matchID, _ := strconv.Atoi(strings.Split(matchHref, "/")[2])
		matchTimestamp, _ := selection.Find(".matchTime").First().Attr("data-unix")
		date := utils.UnixTimeStringToTime(matchTimestamp)

		eventName := selection.Find(".matchEventName").First().Text()
		eventID, _ := strconv.Atoi(
			strings.Split(utils.PopSlashSource(selection.Find("img.matchEventLogo")), ".")[0])

		team1Name := selection.Find(".team1 .matchTeamName").First().Text()
		team1IDStr, _ := selection.Attr("team1")
		team1ID, _ := strconv.Atoi(team1IDStr)

		team2Name := selection.Find(".team2 .matchTeamName").Last().Text()
		team2IDStr, _ := selection.Attr("team2")
		team2ID, _ := strconv.Atoi(team2IDStr)

		format := selection.Find(".matchMeta").Last().Text()

		match := &model.UpcomingMatch{
			ID: &matchID,
			Team1: model.Team{
				Name: team1Name,
				ID:   &team1ID,
			},
			Team2: model.Team{
				Name: team2Name,
				ID:   &team2ID,
			},
			Date: date,
			Event: model.Event{
				Name: eventName,
				ID:   &eventID,
			},
			Format: format,
		}
		upcomingMatches = append(upcomingMatches, match)
	})

	return upcomingMatches, nil
}
