package controllers

import (
	"github.com/revel/revel"
)

type Patient struct {
	*revel.Controller
}

func (c Patient) Index() revel.Result {
	return c.Render()
}
