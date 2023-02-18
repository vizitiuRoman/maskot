package model

import "time"

type Balance struct {
	ID         int
	PlayerName string
	Balance    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
