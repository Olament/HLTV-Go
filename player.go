package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"hltv/model"
	"hltv/utils"
	"net/http"
	"strconv"
	"strings"
)

func (h *HLTV) GetPlayers(id int) (player *model.FullPlayer, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", h.Url+"/player/"+strconv.Itoa(id)+"/-", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, &HTTPError{
			Code:        res.StatusCode,
			Description: http.StatusText(res.StatusCode),
		}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	/* basic information */
	name := strings.Trim(doc.Find(".playerRealname").Text(), " ") // player's real name
	ign := doc.Find(".playerNickname").Text()                     // player's in-game name
	image, _ := doc.Find(".bodyshot-img").Attr("src")
	age, _ := strconv.Atoi(strings.Split(doc.Find(".playerAge .listRight").Text(), " ")[0])

	/* social media */
	twitter, _ := doc.Find(".twitter").Parent().Attr("href")
	twitch, _ := doc.Find(".twitch").Parent().Attr("href")
	facebook, _ := doc.Find(".facebook").Parent().Attr("href")

	country, _ := doc.Find(".playerRealname .flag").Attr("alt")
	code := strings.Split(utils.PopSlashSource(doc.Find(".playerRealname .flag")), ".")[0]

	teamname := strings.Trim(doc.Find(".playerTeam a").Text(), " ")
	teamhref, _ := doc.Find(".playerTeam a").Attr("href")
	var teamid int
	if len(teamhref) > 0 {
		teamhref = strings.Split(teamhref, "/")[2]
		teamid, _ = strconv.Atoi(teamhref)
	}

	/* Player statistics */
	stats := getMapStat(doc)
	rating, _ := strconv.ParseFloat(stats[0], 32)
	killsPerRound, _ := strconv.ParseFloat(stats[1], 32)
	headshots, _ := strconv.ParseFloat(strings.ReplaceAll(stats[2], "%", ""), 32)
	mapsPlayed, _ := strconv.Atoi(strings.ReplaceAll(stats[3], "%", ""))
	deathPerRound, _ := strconv.ParseFloat(stats[4], 32)
	roundsContributed, _ := strconv.ParseFloat(strings.ReplaceAll(stats[5], "%", ""), 32)

	/* achievement */
	var achievements []model.Achievement
	doc.Find(".achievement-table .team").Each(func(i int, selection *goquery.Selection) {
		/* prepare entry for event */
		eventName := selection.Find(".tournament-name-cell a").Text()
		eventidS, _ := selection.Find(".tournament-name-cell a").Attr("href")
		eventID, _ := strconv.Atoi(strings.Split(eventidS, "/")[2])
		eventPlace := selection.Find(".achievement").Text()

		currAchievement := model.Achievement{
			Event: model.Event{
				Name: eventName,
				ID:   &eventID,
			},
			Place: eventPlace,
		}

		achievements = append(achievements, currAchievement)
	})

	return &model.FullPlayer{
		ID:    id,
		Name:  &name,
		Ign:   ign,
		Image: &image,
		Age:   &age,
		Country: model.Country{
			Name: country,
			Code: code,
		},
		Team: model.Team{
			Name: teamname,
			ID:   &teamid,
		},
		Twitter:  &twitter,
		Twitch:   &twitch,
		Facebook: &facebook,
		Statistics: model.Statistics{
			Rating:            float32(rating),
			KillsPerRound:     float32(killsPerRound),
			MapsPlayed:        mapsPlayed,
			DeathsPerRound:    float32(deathPerRound),
			Headshots:         float32(headshots),
			RoundsContributed: float32(roundsContributed),
		},
		Achievements: achievements,
	}, nil
}

/* a helper method that returns an array of stats */
func getMapStat(doc *goquery.Document) (stats []string) {
	doc.Find(".tab-content .two-col").Find(".cell").Find(".statsVal").Each(
		func(i int, selection *goquery.Selection) {
			stats = append(stats, selection.Text())
		})
	return stats
}
