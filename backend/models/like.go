package models

import (
	"time"
)

type Like struct {
	user      string
	post      string
	createdAt time.Time
	updatedAt time.Time
	quantity  int
}
