package controllers

import (
	"fmt"
	"github.com/myaccount/myapp/app/model/"
	"github.com/revel/revel"
	"log"
)

type Choose struct {
	*revel.Controller
}

// ログイン処理を行う
// クッキーとセッションを保存する
// 開き時間を返す
func (c Choose) Choose(Id string, Pw string ,Cookie string) revel.Result {

	c.Validation.Required(Id).Message("Id is wrong")
	c.Validation.Required(Pw).Message("Pw is wrong")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Render()
	}

// ログイン処理を行う
login(Id,Pw)
	// ユーザーを追加する処理
	FindSlots(Cookie)

	return c.Render()
}

func FindSlots(Cookie string) []reservation {

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	bookableSlot := []Reservation{}

	// #現在時間+1時間
	// TODO ユーザーごとの処理を追加する
	loc, _ := time.LoadLocation("Asia/Tokyo")
	datetime := loc.Add(1 * time.Minute)

	err = db.Select(&bookableSlot, "select * from reservation where datetime >= ? and patient_id == null;", datetime)

	return bookableSlot

}


// ID/PW を元にログインする
// TODO セッションを登録する Cookieを返す
func Login(Id string, Pw string) {

    db, err := Connect()
    if err != nil {
        log.Fatalln(err)
    }

    people := []UserInfomation{}
    err = db.Select(&people, "SELECT id, name From patient;")
    if err != nil {
        fmt.Println(err)
    }

}
