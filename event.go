package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"hltv/enum"
	"hltv/model"
	"hltv/utils"
	"strconv"
	"strings"
)

func (h *HLTV) GetEvent(id int) (fullEvent *model.FullEvent, err error) {
	res, err := utils.GetQuery(h.Url + "/events/" + strconv.Itoa(id) + "/-")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	/* get basic information */
	name := doc.Find(".eventname").Text()
	dateStart, _ := doc.Find("td.eventdate span[data-unix]").First().Attr("data-unix")
	dateEnd, _ := doc.Find("td.eventdate span[data-unix]").Last().Attr("data-unix")
	prizePool := doc.Find("td.prizepool").Text()

	/* get location */
	countryName, _ := doc.Find("img.flag").Attr("title")
	countryID := strings.Split(utils.PopSlashSource(doc.Find("img.flag")), ".")[0]

	/* Team information */
	var teams []model.EventTeam
	doc.Find(".team-box").Each(func(i int, selection *goquery.Selection) {
		teamName, _ := selection.Find(".logo").Attr("title")
		teamID, _ := strconv.Atoi(utils.PopSlashSource(selection.Find(".logo")))
		reasonForParticipation := strings.Trim(selection.Find(".sub-text").Text(), " ")
		rankDuringEvent, _ := strconv.Atoi(strings.ReplaceAll(selection.Find(".event-world-rank").Text(), "#", ""))

		currTeam := model.EventTeam{
			Team: model.Team{
				Name: teamName,
				ID:   &teamID,
			},
			ReasonForParticipation: reasonForParticipation,
			RankDuringEvent:        &rankDuringEvent,
		}

		teams = append(teams, currTeam)
	})

	var relativeEvents []model.Event
	doc.Find(".related-event").Each(func(i int, selection *goquery.Selection) {
		eventName := selection.Find(".event-name").Text()
		eventIDS, _ := selection.Find("a").Attr("href")
		eventID, _ := strconv.Atoi(strings.Split(eventIDS, "/")[2])

		currEvent := model.Event{
			Name: eventName,
			ID:   &eventID,
		}
		relativeEvents = append(relativeEvents, currEvent)
	})

	var prizeDistribution []model.EventPrizeDistribution
	doc.Find(".placement").Each(func(i int, selection *goquery.Selection) {
		place := selection.Children().Eq(1).Text()
		prize := selection.Find(".prizeMoney").First().Text()
		/* sometimes the winning team not only get the prize money but also get other forms of
		prize. For example, team who won the IEM Katowice 2019 may be invited directly to
		IEM Katowice 2020 */
		otherPrize := selection.Find(".prizeMoney").First().Next().Text()
		var qualifiedEvent *model.Event
		if otherPrize != "" {
			for _, event := range relativeEvents {
				if event.Name == otherPrize {
					qualifiedEvent = &event
				}
			}
		}

		var team *model.Team
		if doc.Find(".team").Children().Length() != 0 {
			teamName := selection.Find(".team a").Text()
			teamIDS, _ := selection.Find(".team a").Attr("href")
			teamID, _ := strconv.Atoi(strings.Split(teamIDS, "/")[2])

			team = &model.Team{
				Name: teamName,
				ID:   &teamID,
			}
		}

		currDistribution := model.EventPrizeDistribution{
			Place:        place,
			Prize:        prize,
			OtherPrize:   otherPrize,
			QualifiesFor: qualifiedEvent,
			Team:         team,
		}

		prizeDistribution = append(prizeDistribution, currDistribution)
	})

	var formats []model.EventFormat
	doc.Find(".format tr").Each(func(i int, selection *goquery.Selection) {
		currFormat := model.EventFormat{
			Type:        selection.Find(".format-header").Text(),
			Description: selection.Find(".format-data").Text(),
		}
		formats = append(formats, currFormat)
	})

	var mapPool []enum.MapSlug
	doc.Find(".map-pool-map-holder").Each(func(i int, selection *goquery.Selection) {
		mapPool = append(mapPool, enum.MapSlug(selection.Find(".map-pool-map-name").Text()))
	})

	return &model.FullEvent{
		ID:        id,
		Name:      name,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		PrizePool: prizePool,
		Teams:     teams,
		Location: model.Country{
			Name: countryName,
			Code: countryID,
		},
		PrizeDistribution: prizeDistribution,
		Formats:           formats,
		RelatedEvents:     relativeEvents,
		MapPool:           mapPool,
	}, nil
}
