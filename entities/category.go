package entities

import "time"

type Category struct {
	Id        uint
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
