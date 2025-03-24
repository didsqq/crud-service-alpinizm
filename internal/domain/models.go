package domain

import "time"

type Alpinist struct {
	ID                int64
	Email             string
	Surname           string
	Name              string
	Address           string
	Phone             string
	Sex               string
	ID_sport_category int64
}

type Climb struct {
	ID          int64     `db:"ID_mountain_climbs"`
	ID_group    int64     `db:"ID_groups"`
	ID_mountain int64     `db:"ID_mountain"`
	ID_category int64     `db:"ID_category"`
	Start_date  time.Time `db:"Start_date_"`
	End_date    time.Time `db:"End_date_"`
	Total       string    `db:"Total"`
}
