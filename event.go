package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"hltv/enum"
	"hltv/model"
	"hltv/utils"
	"strconv"
	"strings"
	"time"
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
	dateStartUnix, _ := doc.Find("td.eventdate span[data-unix]").First().Attr("data-unix")
	dateStartInt64, _ := strconv.ParseInt(dateStartUnix, 10, 64)
	dateStart := time.Unix(dateStartInt64, 0)
	dateEndUnix, _ := doc.Find("td.eventdate span[data-unix]").Last().Attr("data-unix")
	dateEndInt64, _ := strconv.ParseInt(dateEndUnix, 10, 64)
	dateEnd := time.Unix(dateEndInt64, 0)
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

/* Return a list of events satisfied given conditions
   Use enum.EventSizeAny to represent any event size and
   use time.Month(0) to represent any month
*/
func (h *HLTV) GetEvents(size enum.EventSize, month time.Month) (events []model.EventResults, err error) {
	res, err := utils.GetQuery(h.Url + "/events/")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	timeLayout := "January 2006" //time layout for string time conversion
	doc.Find(".events-month").Each(func(i int, eventsEachMonth *goquery.Selection) {
		currMonthString := eventsEachMonth.Find(".standard-headline").Text()
		currTime, _ := time.Parse(timeLayout, currMonthString)

		/* decide which event type should be added to events
		   month == 0 means search event at any month */
		if month == 0 || month == currTime.Month() {
			// TODO better error handling for parseEvent
			if size == enum.EventSizeAny || size == enum.EventSizeBig {
				parsedEvents, _ := parseEvent(eventsEachMonth.Find("a.big-event"), currTime.Month(), enum.EventSizeBig)
				events = append(events, parsedEvents)
			}
			if size == enum.EventSizeAny || size == enum.EventSizeSmall {
				parsedEvents, _ := parseEvent(eventsEachMonth.Find("a.small-event"), currTime.Month(), enum.EventSizeSmall)
				events = append(events, parsedEvents)
			}
		}
	})

	return events, nil
}

func parseEvent(eventSelection *goquery.Selection, month time.Month, eventSize enum.EventSize) (result model.EventResults, err error) {
	/* initialize goquery selector */
	var dateSelector string = ""
	var nameSelector string = ""
	var locationSelector string = ""

	if eventSize == enum.EventSizeSmall {
		dateSelector = ".eventDetails .col-desc span[data-unix]"
		nameSelector = ".col-value .text-ellipsis"
		locationSelector = ".smallCountry img"
	} else {
		/* selector for Big Event */
		//dateSelector = "span[data-unix]"
		dateSelector = ".additional-info .col-date"
		nameSelector = ".big-event-name"
		locationSelector = ".location-top-teams img"
	}

	var events []model.SimpleEvent
	eventSelection.Each(func(i int, selection *goquery.Selection) {
		/* basic information */
		idS, _ := selection.Attr("href")
		id, _ := strconv.Atoi(strings.Split(idS, "/")[2])
		eventName := selection.Find(nameSelector).Text()

		/* country */
		countryName, _ := selection.Find(locationSelector).Attr("title")
		countryCodeS := strings.Split(utils.PopSlashSource(selection.Find(locationSelector)), ".")[0]

		/* Date */
		//TODO fix date, it is totally broken
		dateStartUnix, _ := selection.Find(dateSelector).Eq(0).Attr("data-unix")
		dateStartInt64, _ := strconv.ParseInt(dateStartUnix, 10, 64)
		dateStart := time.Unix(dateStartInt64, 0)

		dateEndUnix, _ := selection.Find(dateSelector).Eq(1).Attr("data-unix")
		dateEndInt64, _ := strconv.ParseInt(dateEndUnix, 10, 64)
		dateEnd := time.Unix(dateEndInt64, 0)

		/* Get teams and prizepool */
		var numberOfTeam int // number of teams
		var prizePool string

		if (eventSize == enum.EventSizeSmall) {
			numberOfTeam, _ = strconv.Atoi(selection.Find(".col-value").Eq(1).Text())
			prizePool = selection.Find(".prizePoolEllipsis").Text()
		} else {
			numberOfTeam, _ = strconv.Atoi(selection.Find(".additional-info tr").Eq(0).Find("td").Eq(2).Text())
			prizePool = selection.Find(".additional-info tr").Eq(0).Find("td").Eq(1).Text()
		}

		//eventType := selection.Find("table tr").Eq(0).Find("td").Eq(3).Text()
		// TODO fix type to EventType

		currEvent := model.SimpleEvent{
			ID:            id,
			Name:          eventName,
			DateStart:     &dateStart,
			DateEnd:       &dateEnd,
			PrizePool:     prizePool,
			NumberOfTeams: &numberOfTeam,
			Location:      model.Country{
				Name: countryName,
				Code: countryCodeS,
			},
			Type:          enum.EventTypeMajor, //TODO temp
		}
		events = append(events, currEvent)
	})

	return model.EventResults{
		Month:  month,
		Events: events,
	}, nil
}
