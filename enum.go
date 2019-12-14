package hltv

type MapSlug string

const (
	MapSlugTBA         MapSlug = "tba"
	MapSlugTrain       MapSlug = "trn"
	MapSlugCobblestone MapSlug = "cbl"
	MapSlugInferno     MapSlug = "inf"
	MapSlugCache       MapSlug = "cch"
	MapSlugMirage      MapSlug = "mrg"
	MapSlugOverpass    MapSlug = "ovp"
	MapSlugDust2       MapSlug = "d2"
	MapSlugNuke        MapSlug = "nuke"
	MapSlugTuscan      MapSlug = "tcn"
	MapSlugVertigo     MapSlug = "vertigo"
	MapSlugSeason      MapSlug = "-"
	MapSlugDefault     MapSlug = "-"
)

// TODO resolve the conflict with maps full name

type MatchType string

const (
	MatchTypeLAN       MatchType = "Lan"
	MatchTypeOnline    MatchType = "Online"
	MatchTypeBigEvents MatchType = "BigEvents"
	MatchTypeMajors    MatchType = "Majors"
)

type RankingFilter string

const (
	RankingFilterTop5  RankingFilter = "Top5"
	RankingFilterTop10 RankingFilter = "Top10"
	RankingFilterTop20 RankingFilter = "Top20"
	RankingFilterTop30 RankingFilter = "Top30"
	RankingFilterTop50 RankingFilter = "Top50"
)

type EventSize string

const (
	EventSizeSmall EventSize = "small"
	EventSizeBig   EventSize = "big"
	EventSizeAny   EventSize = "any"
)

type EventType string

const (
	EventTypeOnline   EventType = "Online"
	EventTypeLocalLan EventType = "Local LAN"
	EventTypeIntlLan  EventType = "International LAN"
	EventTypeRegLan   EventType = "Regional LAN"
	EventTypeMajor    EventType = "Major"
	EventTypeOther    EventType = "Other"
)
