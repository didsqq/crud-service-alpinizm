package domain

import "time"

type User struct {
	ID              int64  `db:"id"`
	Surname         string `db:"surname"`
	Name            string `db:"name_"`
	Address         string `db:"address_"`
	Phone           string `db:"phone"`
	Sex             string `db:"sex"`
	IdSportCategory int    `db:"id_sport_category"`
	Username        string `db:"username"`
	Password        string `db:"password_"`
}

type Climb struct {
	ID         int64     `db:"id"`
	IdGroup    int64     `db:"id_groups"`
	IdMountain int64     `db:"id_mountain"`
	IdCategory int64     `db:"id_category"`
	StartDate  time.Time `db:"start_date"`
	EndDate    time.Time `db:"end_date"`
	Total      string    `db:"total"`
	PhotoUrl   string    `db:"photo_url"`
}

type Equipment struct {
	ID                int    `db:"id"`
	Title             string `db:"title"`
	QuantityAvailable int    `db:"quantity_available"`
	ImageUrl          string `db:"image_url"`
	Description       string `db:"description"`
}

type Mountain struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	Height        int    `db:"height"`
	MountainRange string `db:"mountain_range"`
}

type SportCategory struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
