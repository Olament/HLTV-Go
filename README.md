<h1 align="center">
  <img src="https://www.hltv.org/img/static/TopLogo2x.png" alt="HLTV logo" width="200">
  <br>
  Go port of <a href="https://github.com/gigobyte/HLTV">HLTV Node.js API</a>
  <br>
</h1>

## Getting Started
```Golang
func main() {
	h := hltv.HLTV{
		Url:       "https://www.hltv.org",
		StaticURL: "",
	}

	events, _ := h.GetEvent(5100)
	res, _ := json.MarshalIndent(events, "", "    ")
	fmt.Println(string(res))
}
```

<details><summary>Output</summary>
<p>
	
```json
{
    "ID": 5100,
    "Name": "WESG 2019 Greater China Regional Finals",
    "DateStart": "2019-12-11T05:00:00-06:00",
    "DateEnd": "2019-12-14T05:00:00-06:00",
    "PrizePool": "$34,100",
    "Teams": [
        {
            "Name": "TYLOO",
            "ID": 4863,
            "ReasonForParticipation": "South China",
            "RankDuringEvent": 21
        },
        {
            "Name": "ViCi",
            "ID": 7606,
            "ReasonForParticipation": "South China",
            "RankDuringEvent": 33
        },
        {
            "Name": "R-Stars",
            "ID": 10472,
            "ReasonForParticipation": "West China",
            "RankDuringEvent": 65
        },
        {
            "Name": "Invictus",
            "ID": 7966,
            "ReasonForParticipation": "West China",
            "RankDuringEvent": 66
        },
        {
            "Name": "EHOME",
            "ID": 7024,
            "ReasonForParticipation": "Qualifier",
            "RankDuringEvent": 78
        },
        {
            "Name": "OneThree",
            "ID": 10022,
            "ReasonForParticipation": "East China",
            "RankDuringEvent": 103
        },
        {
            "Name": "Lynn Vision",
            "ID": 8840,
            "ReasonForParticipation": "North China",
            "RankDuringEvent": 119
        },
        {
            "Name": "ahq",
            "ID": 10052,
            "ReasonForParticipation": "Taiwan",
            "RankDuringEvent": 187
        },
        {
            "Name": "4Bowie",
            "ID": 10529,
            "ReasonForParticipation": "East China",
            "RankDuringEvent": 0
        },
        {
            "Name": "L9L",
            "ID": 10480,
            "ReasonForParticipation": "Hong Kong",
            "RankDuringEvent": 0
        },
        {
            "Name": "Macau Legend",
            "ID": 9622,
            "ReasonForParticipation": "Macau",
            "RankDuringEvent": 0
        },
        {
            "Name": "DMC",
            "ID": 10528,
            "ReasonForParticipation": "North China",
            "RankDuringEvent": 0
        }
    ],
    "Location": {
        "Name": "China",
        "Code": "CN"
    },
    "PrizeDistribution": [
        {
            "Place": "1st",
            "Prize": "$21,325",
            "OtherPrize": "WESG 2019 World Finals",
            "QualifiesFor": null,
            "Team": {
                "Name": "Lynn Vision",
                "ID": 8840
            }
        },
        {
            "Place": "2nd",
            "Prize": "$8,530",
            "OtherPrize": "WESG 2019 APAC Finals",
            "QualifiesFor": {
                "Name": "WESG 2019 Taiwan",
                "ID": 5057
            },
            "Team": {
                "Name": "ViCi",
                "ID": 7606
            }
        },
        {
            "Place": "3rd",
            "Prize": "$4,265",
            "OtherPrize": "WESG 2019 APAC Finals",
            "QualifiesFor": {
                "Name": "WESG 2019 Taiwan",
                "ID": 5057
            },
            "Team": {
                "Name": "TYLOO",
                "ID": 4863
            }
        },
        {
            "Place": "4th",
            "Prize": "",
            "OtherPrize": "WESG 2019 APAC Finals",
            "QualifiesFor": {
                "Name": "WESG 2019 Taiwan",
                "ID": 5057
            },
            "Team": {
                "Name": "R-Stars",
                "ID": 10472
            }
        },
        {
            "Place": "5-6th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "OneThree",
                "ID": 10022
            }
        },
        {
            "Place": "5-6th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "Invictus",
                "ID": 7966
            }
        },
        {
            "Place": "7-8th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "ahq",
                "ID": 10052
            }
        },
        {
            "Place": "7-8th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "4Bowie",
                "ID": 10529
            }
        },
        {
            "Place": "9-10th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "DMC",
                "ID": 10528
            }
        },
        {
            "Place": "9-10th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "EHOME",
                "ID": 7024
            }
        },
        {
            "Place": "11-12th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "Macau Legend",
                "ID": 9622
            }
        },
        {
            "Place": "11-12th",
            "Prize": "",
            "OtherPrize": "",
            "QualifiesFor": null,
            "Team": {
                "Name": "L9L",
                "ID": 10480
            }
        }
    ],
    "Formats": null,
    "RelatedEvents": [
        {
            "Name": "WESG 2019 APAC Finals",
            "ID": 4996
        },
        {
            "Name": "WESG 2019 Hong Kong",
            "ID": 5050
        },
        {
            "Name": "WESG 2019 Taiwan",
            "ID": 5057
        }
    ],
    "MapPool": [
        "Dust2",
        "Mirage",
        "Inferno",
        "Nuke",
        "Train",
        "Overpass",
        "Vertigo"
    ]
}
```
</p>
</details>

## API
* Player
	* [GetPlayer](#getplayer)
	* [GetPlayerByName](#getplayerbynamename)
	* [GetPlayerStats](#getplayerstats)
	    

#### GetPlayer
```golang
//GetPlayer(id int) (player *model.FullPlayer, err error)
h.GetPlayer(7798)
```

<details><summary>Output</summary>
<p>
	
```json
{
    "Name": "Aleksandr Kostyliev",
    "ID": 7998,
    "Ign": "s1mple",
    "Image": "https://static.hltv.org//images/playerprofile/bodyshot/compressed/7998.png",
    "Age": 22,
    "Country": {
        "Name": "Ukraine",
        "Code": "UA"
    },
    "Team": {
        "Name": "Natus Vincere",
        "ID": 4608
    },
    "Twitter": "https://twitter.com/s1mpleO",
    "Twitch": "http://www.twitch.tv/s1mple",
    "Facebook": "https://www.facebook.com/Officials1mple",
    "Statistics": {
        "Kills": 0,
        "Headshots": 53.3,
        "Death": 0,
        "KDRatio": 0,
        "DamgePerRound": 0,
        "GrenadeDamge": 0,
        "MapsPlayed": 34,
        "RoundsPlayed": 0,
        "KillsPerRound": 0.86,
        "AssistsPerRound": 0,
        "DeathsPerRound": 0.61,
        "SavedByTeammatePerRound": 0,
        "SavedTeammatesPerRound": 0,
        "Rating": 1.31,
        "RoundsContributed": 74.6
    },
    "Achievements": [
        {
            "Event": {
                "Name": "StarLadder Major Berlin 2019",
                "ID": 4443
            },
            "Place": "1/4 final"
        },
        {
            "Event": {
                "Name": "IEM Katowice 2019",
                "ID": 3883
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "FACEIT Major 2018",
                "ID": 3564
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "ELEAGUE Major 2018",
                "ID": 3247
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "PGL Major Krakow 2017",
                "ID": 2720
            },
            "Place": "Group stage"
        },
        {
            "Event": {
                "Name": "ELEAGUE Major 2017",
                "ID": 2471
            },
            "Place": "1/4 final"
        },
        {
            "Event": {
                "Name": "ESL One Cologne 2016",
                "ID": 2062
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "MLG Columbus 2016",
                "ID": 2027
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "DreamHack Winter 2014",
                "ID": 1553
            },
            "Place": "1/4 final"
        },
        {
            "Event": {
                "Name": "ESL Pro League Season 10 Finals",
                "ID": 4697
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "BLAST Pro Series Copenhagen 2019",
                "ID": 4702
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "DreamHack Masters Malmö 2019",
                "ID": 4553
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "ESL One Cologne 2019",
                "ID": 4281
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "StarSeries i-League Season 7",
                "ID": 4240
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "IEM Katowice 2019",
                "ID": 3883
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "BLAST Pro Series Lisbon 2018",
                "ID": 4205
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "BLAST Pro Series Copenhagen 2018",
                "ID": 3701
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "EPICENTER 2018",
                "ID": 3985
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "FACEIT Major 2018",
                "ID": 3564
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "ELEAGUE CS:GO Premier 2018",
                "ID": 3515
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "ESL One Cologne 2018",
                "ID": 3392
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "CS:GO Asia Championships 2018",
                "ID": 3714
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "StarSeries i-League Season 5",
                "ID": 3666
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "ESL Pro League Season 7 Finals",
                "ID": 3373
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "DreamHack Masters Marseille 2018",
                "ID": 3491
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "StarSeries i-League Season 4",
                "ID": 3486
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "ELEAGUE Major 2018",
                "ID": 3247
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "DreamHack Open Winter 2017",
                "ID": 2574
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "ESL One Cologne 2017",
                "ID": 2635
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "SL i-League StarSeries Season 3 Finals",
                "ID": 2683
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "EPICENTER: Moscow",
                "ID": 2410
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "ESL One New York 2016",
                "ID": 2239
            },
            "Place": "1st"
        },
        {
            "Event": {
                "Name": "ESL One Cologne 2016",
                "ID": 2062
            },
            "Place": "2nd"
        },
        {
            "Event": {
                "Name": "MLG Columbus 2016",
                "ID": 2027
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "ESWC 2015",
                "ID": 1707
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "SLTV StarSeries XIII Finals",
                "ID": 1607
            },
            "Place": "3rd"
        },
        {
            "Event": {
                "Name": "Game Show Season 1 Finals",
                "ID": 1502
            },
            "Place": "2nd"
        }
    ]
}
```

</p>
</details>

#### GetPlayerByName

```golang
//GetPlayerByName(name string) (player *model.FullPlayer, err error)
h.GetPlayerByName("s1mple")
```

#### GetPlayerStats
```golang
//GetPlayerStats(id int, q PlayerStatsQuery) (playerStats *model.FullPlayerStats, err error)
h.GetPlayerStats(7998, hltv.PlayerStatsQuery{
		StartDate:  "2019-01-01", //YYYY-MM-DD
		EndDate:    "2019-12-31",
		MatchType:  enum.MatchTypeLAN,
		RankFilter: enum.RankingFilterTop20,
})
```
<details><summary>Output</summary>
<p>	
	
```json
{
    "Name": "Aleksandr Kostyliev",
    "ID": 7998,
    "Ign": "s1mple",
    "Image": "https://static.hltv.org/images/playerprofile/thumb/7998/400.jpeg?v=22",
    "Age": 22,
    "Country": {
        "Name": "Ukraine",
        "Code": "UA"
    },
    "Team": {
        "Name": "Natus Vincere",
        "ID": 4608
    },
    "Statistics": {
        "Kills": 2956,
        "Headshots": 42.7,
        "Death": 2069,
        "KDRatio": 1.43,
        "DamgePerRound": 85.1,
        "GrenadeDamge": 2.8,
        "MapsPlayed": 129,
        "RoundsPlayed": 3470,
        "KillsPerRound": 0.85,
        "AssistsPerRound": 0.1,
        "DeathsPerRound": 0.6,
        "SavedByTeammatePerRound": 0.08,
        "SavedTeammatesPerRound": 0.11,
        "Rating": 1.3,
        "RoundsContributed": 74.8
    }
}
```

</p>
</details>
