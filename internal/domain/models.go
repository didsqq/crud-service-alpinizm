package domain

import "time"

type User struct {
	ID                int64  `db:"ID_alpinist"`
	Surname           string `db:"Surname"`
	Name              string `db:"Name_"`
	Address           string `db:"Address_"`
	Phone             string `db:"Phone"`
	Sex               string `db:"Sex"`
	ID_sport_category int    `db:"ID_sport_category"`
	Username          string `db:"Username"`
	Password          string `db:"Password_"`
}

type Climb struct {
	ID          int64     `db:"ID_mountain_climbs"`
	ID_group    int64     `db:"ID_groups"`
	ID_mountain int64     `db:"ID_mountain"`
	ID_category int64     `db:"ID_category"`
	Start_date  time.Time `db:"Start_date_"`
	End_date    time.Time `db:"End_date_"`
	Total       string    `db:"Total"`
	Photo_Url   string    `db:"Photo_url"`
}
