package training

import "time"

type Training struct {
	Id       uint      `json:"id"`
	UserId   string    `json:"userId" db:"user_id" validate:"required,gte=3,lte=255"`
	Title    string    `json:"title" validate:"required,gte=3,lte=255"`
	Type     string    `json:"type" validate:"required,oneof=run ride"`
	Distance uint      `json:"distance" validate:"required,gt=0,lt=1000000"`
	Time     uint      `json:"time" validate:"required,gt=0,lte=864000"`
	Date     time.Time `json:"date" validate:"required"`
}
