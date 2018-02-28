package model

import (
	"time"
)

type Patient struct {
	id         string
	name       string
	phone      string
	birthday   string
	created_at time.Time
	updated_at time.Time
}
