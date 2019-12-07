package hltv

import (
	"github.com/PuerkitoBio/goquery"
	"hltv/model"
	"net/http"
	"strconv"
)
type HLTV struct {
	Url       string
	StaticURL string
}

func (h *HLTV) GetPlayers(id int) (player *model.FullPlayer, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", h.Url + "/player/" + strconv.Itoa(id) + "/-", nil)
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

	name := doc.Find(".playerRealname").Text()

	return &model.FullPlayer{
		ID:           0,
		Name:         &name,
		Ign:          "",
		Image:        nil,
		Age:          nil,
		Country:      model.Country{
			Name: "",
			Code: "",
		},
		Team:         model.Team{
			Name: "",
			ID:   nil,
		},
		Twitter:      nil,
		Facebook:     nil,
		Statistics:   model.Statistics{
			Rating:            0,
			KillsPerRound:     0,
			MapsPlayed:        0,
			DeathsPerRound:    0,
			Headshots:         0,
			RoundsContributed: 0,
		},
		Achievements: nil,
	}, nil
}
