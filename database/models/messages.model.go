package models

import "time"

type Messages struct {
	UserID    uint      `json:"userId"`
	Text      string    `json:"text"`
	Visible   bool      `json:"visible"`
	CreatedAt time.Time `json:"createdAt"`
}
