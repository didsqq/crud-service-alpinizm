package domain

import "time"

type User struct {
	ID                int64  `db:"id"`
	Surname           string `db:"surname"`
	Name              string `db:"name_"`
	Address           string `db:"address_"`
	Phone             string `db:"phone"`
	Sex               string `db:"sex"`
	ID_sport_category int    `db:"id_sport_category"`
	Username          string `db:"username"`
	Password          string `db:"password_"`
}

type Climb struct {
	ID          int64     `db:"id"`
	ID_group    int64     `db:"id_groups"`
	ID_mountain int64     `db:"id_mountain"`
	ID_category int64     `db:"id_category"`
	Start_date  time.Time `db:"start_date_"`
	End_date    time.Time `db:"end_date_"`
	Total       string    `db:"total"`
	Photo_Url   string    `db:"photo_url"`
}

type Equipment struct {
	ID                int    `db:"id"`
	Title             string `db:"title"`
	QuantityAvailable int    `db:"quantity_available"`
	ImageUrl          string `db:"image_url"`
	Description       string `db:"description"`
}
