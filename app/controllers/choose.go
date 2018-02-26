package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"log"
)

type Choose struct {
	*revel.Controller
}

func (c Choose) Choose(Cookie string) revel.Result {

	c.Validation.Required(Cookie).Message("Cookie is wrong")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		//return c.Redirect(Choose.Adduser)
		return c.Redirect(Choose.Menu)
	}


        // ユーザーを追加する処理
	FindSlots(Cookie)

	return c.Render()
}

func FindSlots(Cookie string) []reservation{

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	bookableSlot := []reservation{}
	
	datetime= #現在時間+1時間


	err = db.Select(&people, "select * from reservation where datetime >= ? and patient_id == null;", 日にち時間 )

	return bookableSlot
    }


}
