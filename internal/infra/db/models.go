// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)

type Balance struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	Amount    float64
}
