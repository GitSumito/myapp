package model

import (
		"time"
	//	"github.com/revel/revel"

	"github.com/revel/modules/orm/gorp/app/controllers"
)

type Patient struct {
	id         string
	name       string
		phone      string
		birthday   string
		created_at time.Time
		updated_at time.Time
	gorpController.Controller
}
