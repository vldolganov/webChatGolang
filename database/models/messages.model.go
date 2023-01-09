package models

import "time"

type Messages struct {
	FromUser  float64   `json:"from_user"`
	ToUser    uint      `json:"to_user"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
