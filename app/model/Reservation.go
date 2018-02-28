package model

import (
	"time"
)

type Reservation struct {
	datetime      time.Time
	room          int
	treat_time    int
	dr_no         string
	assist_no     string
	dr_weight     int
	assist_weight int
	patient_id    string
	created_at    time.Time
	updated_at    time.Time
}
