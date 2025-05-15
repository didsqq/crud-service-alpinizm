package domain

import "time"

type ApiResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type AlpinistEquipment struct {
	ID                int       `db:"id"`
	Title             string    `db:"title"`
	QuantityAvailable int       `db:"quantity_available"`
	ImageUrl          string    `db:"image_url"`
	Description       string    `db:"description"`
	DateOfIssue       time.Time `db:"date_of_issue"`
	DateOfReturn      time.Time `db:"date_of_return"`
	Status            string    `db:"status"`
}

type AlpinistsEquipments struct {
	ID                int       `db:"id"`
	AlpinistID        int       `db:"alpinist_id"`
	EquipmentID       int       `db:"equipment_id"`
	Title             string    `db:"title"`
	QuantityAvailable int       `db:"quantity_available"`
	ImageUrl          string    `db:"image_url"`
	Description       string    `db:"description"`
	DateOfIssue       time.Time `db:"date_of_issue"`
	DateOfReturn      time.Time `db:"date_of_return"`
	Status            string    `db:"status"`
}
