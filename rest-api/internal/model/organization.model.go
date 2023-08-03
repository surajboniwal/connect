package model

import "time"

type Organization struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"-"`
}
