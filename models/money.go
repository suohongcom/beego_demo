package models

import "time"

type Money struct {
	Id int
	Username string
	Address string
	Money float64
	Status int8
	Created time.Time
	Updated time.Time
}

func (m *Money) TableName() string {
	return TableName("money")
}