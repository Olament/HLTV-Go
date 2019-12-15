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
