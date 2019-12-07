package model

type FullPlayer struct {
	ID           int
	Name         *string
	Ign          string
	Image        *string
	Age          *int
	Country      Country
	Team         Team
	Twitter      *string
	Twitch       *string
	Facebook     *string
	Statistics   Statistics
	Achievements []Achievement
}
