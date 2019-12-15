package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
	"github.com/tidwall/gjson"
	"github.com/Olament/HLTV-Go/enum"
	"github.com/Olament/HLTV-Go/model"
	"github.com/Olament/HLTV-Go/utils"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (h *HLTV) GetPlayer(id int) (player *model.FullPlayer, err error) {
	res, err := utils.GetQuery(h.Url + "/player/" + strconv.Itoa(id) + "/-")
	defer res.Body.Close()
	if err != nil {
		return nil, err
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
	mapsPlayed, _ := strconv.Atoi(stats[3])
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
		Player: model.Player{
			Name: name,
			ID:   &id,
		},
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

func (h *HLTV) GetPlayerByName(name string) (player *model.FullPlayer, err error) {
	res, err := utils.GetQuery(h.Url + "/search?term=" + name)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	query := string(body)
	query = query[1 : len(query)-1] // stupid and ugly workaround, need fix
	id := gjson.Get(query, "players.0.id")

	return h.GetPlayer(int(id.Int()))
}

type PlayerStatsQuery struct {
	StartDate  string //YYYY-MM-DD
	EndDate    string
	MatchType  enum.MatchType
	RankFilter enum.RankingFilter
}

func (h *HLTV) GetPlayerStats(id int, q PlayerStatsQuery) (playerStats *model.FullPlayerStats, err error) {
	queryString, _ := query.Values(q)

	res, err := utils.GetQuery(h.Url + "/stats/players/" + strconv.Itoa(id) + "/-?" + queryString.Encode())
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	/* basic info */
	name := doc.Find(".summaryRealname div").Text()
	ign := doc.Find(".context-item-name").Text() // in-game name
	image, _ := doc.Find(".context-item-image").Attr("src")
	age, _ := strconv.Atoi(strings.Split(doc.Find(".summaryPlayerAge").Text(), " ")[0])

	countryName, _ := doc.Find(".summaryRealname .flag").Attr("title")
	countryCode := strings.Split(utils.PopSlashSource(doc.Find(".summaryRealname .flag")), ".")[0]

	team := doc.Find(".SummaryTeamname").Text()
	teamIDS, _ := doc.Find(".SummaryTeamname").Find("a").Attr("href")
	teamid, _ := strconv.Atoi(strings.Split(teamIDS, "/")[3])

	/* stat */
	var roundsContributed float64 //this is special, treat differently
	doc.Find(".summaryStatBreakdown .summaryStatBreakdownDataValue").Each(func(i int, selection *goquery.Selection) {
		if i == 2 {
			roundsContributed, _ = strconv.ParseFloat(strings.ReplaceAll(selection.Text(), "%", ""), 32)
		}
	})

	stats := make([]string, 0)
	doc.Find(".stats-row").Find("span").Each(func(i int, selection *goquery.Selection) {
		if i%2 != 0 {
			stats = append(stats, selection.Text())
		}
	})

	kills, _ := strconv.Atoi(stats[0])
	headshots, _ := strconv.ParseFloat(strings.ReplaceAll(stats[1], "%", ""), 32)
	death, _ := strconv.Atoi(stats[2])
	kdRatio, _ := strconv.ParseFloat(stats[3], 32)
	damgePerRound, _ := strconv.ParseFloat(stats[4], 32)
	grenadeDamge, _ := strconv.ParseFloat(stats[5], 32)
	mapsPlayed, _ := strconv.Atoi(stats[6])
	roundsPlayed, _ := strconv.Atoi(stats[7])
	killsPerRound, _ := strconv.ParseFloat(stats[8], 32)
	assistsPerRound, _ := strconv.ParseFloat(stats[9], 32)
	deathsPerRound, _ := strconv.ParseFloat(stats[10], 32)
	savedByTeammatePerRound, _ := strconv.ParseFloat(stats[11], 32)
	savedTeammatesPerRound, _ := strconv.ParseFloat(stats[12], 32)
	rating, _ := strconv.ParseFloat(stats[13], 32)

	return &model.FullPlayerStats{
		Player: model.Player{
			Name: name,
			ID:   &id,
		},
		Ign:   &ign,
		Image: &image,
		Age:   &age,
		Country: &model.Country{
			Name: countryName,
			Code: countryCode,
		},
		Team: &model.Team{
			Name: team,
			ID:   &teamid,
		},
		Statistics: model.Statistics{
			Kills:                   kills,
			Headshots:               float32(headshots),
			Death:                   death,
			KDRatio:                 float32(kdRatio),
			DamgePerRound:           float32(damgePerRound),
			GrenadeDamge:            float32(grenadeDamge),
			MapsPlayed:              mapsPlayed,
			RoundsPlayed:            roundsPlayed,
			KillsPerRound:           float32(killsPerRound),
			AssistsPerRound:         float32(assistsPerRound),
			DeathsPerRound:          float32(deathsPerRound),
			SavedByTeammatePerRound: float32(savedByTeammatePerRound),
			SavedTeammatesPerRound:  float32(savedTeammatesPerRound),
			Rating:                  float32(rating),
			RoundsContributed:       float32(roundsContributed),
		},
	}, nil
}

func (h *HLTV) GetPlayerRanking(q PlayerStatsQuery) (players []model.PlayerRanking, err error) {
	queryString, _ := query.Values(q)

	res, err := utils.GetQuery(h.Url + "/stats/players?" + queryString.Encode())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, &utils.HTTPError{
			Code:        res.StatusCode,
			Description: http.StatusText(res.StatusCode),
		}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".player-ratings-table tbody tr").Each(func(i int, selection *goquery.Selection) {
		idS, _ := selection.Find(".playerCol a").First().Attr("href")
		id, _ := strconv.Atoi(strings.Split(idS, "/")[3])
		name := selection.Find(".playerCol").Text()
		rating, _ := strconv.ParseFloat(selection.Find(".ratingCol").Text(), 32)

		players = append(players, model.PlayerRanking{
			Player: model.Player{
				Name: name,
				ID:   &id,
			},
			Rating: float32(rating),
		})
	})

	return players, nil
}
