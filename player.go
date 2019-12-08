package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
	"github.com/tidwall/gjson"
	"hltv/enum"
	"hltv/model"
	"hltv/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

func (h *HLTV) GetPlayer(id int) (player *model.FullPlayer, err error) {
	res, _ := utils.GetQuery(h.Url+"/player/"+strconv.Itoa(id)+"/-")
	defer res.Body.Close()

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

func (h *HLTV) GetPlayerByName(name string) (player *model.FullPlayer, err error) {
	res, _ := utils.GetQuery(h.Url+"/search?term="+name)
	defer res.Body.Close()

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
	StartDate string //YYYY-MM-DD
	EndDate string
	MatchType enum.MatchType
	RankFilter enum.RankingFilter
}

func (h *HLTV) GetPlayerStats(id int, q PlayerStatsQuery) (stats *model.FullPlayerStats, err error) {
	queryString, _ := query.Values(q)

	res, _ := utils.GetQuery(h.Url+"/stats/players/"+strconv.Itoa(id)+"/-?" + queryString.Encode())
	defer res.Body.Close()

	/* basic info */


	return &model.FullPlayerStats{
		Name:       nil,
		Ign:        nil,
		Image:      nil,
		Age:        nil,
		Country:    nil,
		Team:       nil,
		Statistics: model.Statistics{
			Kills:                   0,
			Headshots:               0,
			Death:                   0,
			KDRatio:                 0,
			DamgePerRound:           0,
			GrenadeDamge:            0,
			MapsPlayed:              0,
			RoundsPlayed:            0,
			KillsPerRound:           0,
			AssistsPerRound:         0,
			DeathsPerRound:          0,
			SavedByTeammatePerRound: 0,
			SavedTeammatesPerRound:  0,
			Rating:                  0,
			RoundsContributed:       0,
		},
	}, nil
}
