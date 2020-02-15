package hltv

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/Olament/HLTV-Go/model"
	"github.com/Olament/HLTV-Go/utils"
)

func (h *HLTV) GetUpcomingMatches(teamIDs []int) (upcomingMatches []*model.UpcomingMatch, err error) {
	// Build query string parameters
	queryString := ""
	if len(teamIDs) > 0 {
		q := url.Values{}
		for _, teamID := range teamIDs {
			q.Add("team", strconv.Itoa(teamID))
		}
		queryString = q.Encode()
	}
	res, err := utils.GetQuery(h.Url + "/matches/?" + queryString)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".upcoming-match").Each(func(i int, selection *goquery.Selection) {
		matchHref, _ := selection.Find(".a-reset").First().Attr("href")
		matchID, _ := strconv.Atoi(strings.Split(matchHref, "/")[2])
		matchTimestamp, _ := selection.Find("div.time").First().Attr("data-unix")
		date := utils.UnixTimeStringToTime(matchTimestamp)

		eventName, _ := selection.Find(".event-logo").Attr("alt")
		eventID, _ := strconv.Atoi(strings.Split(utils.PopSlashSource(selection.Find("img.event-logo")), ".")[0])

		team1Name := selection.Find("div.team").First().Text()
		team1ID, _ := strconv.Atoi(utils.PopSlashSource(selection.Find("img.logo").First()))

		team2Name := selection.Find("div.team").Last().Text()
		team2ID, _ := strconv.Atoi(utils.PopSlashSource(selection.Find("img.logo").Last()))

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
		}
		upcomingMatches = append(upcomingMatches, match)
	})

	return upcomingMatches, nil
}
