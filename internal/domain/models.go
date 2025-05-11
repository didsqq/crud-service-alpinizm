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
	ID          int64       `db:"id"`
	IdMountain  int64       `db:"id_mountain"`
	IdCategory  int64       `db:"id_category"`
	Title       string      `db:"title"`
	Season      string      `db:"season"`
	Duration    string      `db:"duration"`
	Distance    string      `db:"distance"`
	Elevation   string      `db:"elevation"`
	MapUrl      string      `db:"map_url"`
	Rating      float64     `db:"rating"`
	Description string      `db:"description"`
	StartDate   time.Time   `db:"start_date"`
	EndDate     time.Time   `db:"end_date"`
	Total       string      `db:"total"`
	PhotoUrl    string      `db:"photo_url"`
	TeamLeaders []Team      `db:"team_leaders"`
	Equipments  []Equipment `db:"equipments"`
	Images      []Image     `db:"images"`
	Mountain    Mountain    `db:"mountain"`
	Category    string      `db:"category"`
}

type Team struct {
	ID          int64  `db:"id"`
	SurnameName string `db:"surname_name"`
	Experience  string `db:"experience"`
	DateOfBirth string `db:"date_of_birth"`
	Address     string `db:"address_"`
	IdPosition  string `db:"id_position"`
	Phone       string `db:"phone"`
	Login       string `db:"login_"`
	Password    string `db:"password_"`
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

type Image struct {
	ID       int    `db:"id"`
	ClimbID  int    `db:"climb_id"`
	ImageUrl string `db:"url"`
}
