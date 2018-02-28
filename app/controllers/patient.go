package controllers

import (
	"fmt"
	"github.com/myaccount/myapp/app/model"
	"github.com/revel/revel"
	"log"
)

//type Patient struct {
//	*revel.Controller
//}

func (c Patient) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.Patient.Choose())
	}
	return c.Render()
}

// ログイン処理を行う
// クッキーとセッションを保存する
// 開き時間を返す
func (c Patient) Choose() revel.Result {
	// ユーザーを追加する処理
	FindSlots(Cookie)
	return c.Render()
}

func FindSlots(Cookie string) {

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
}

// ID/PW を元にログインする
// TODO セッションを登録する Cookieを返す
func Login(username, password string, remember bool) revel.Result {

	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Hotels.Index())
		}
	}

	db, err := Connect()

	people := []UserInfomation{}
	err = db.Select(&people, "SELECT id, name From patient;")

	if err != nil {
		c.Flash.Out["username"] = username
		c.Flash.Error("Login failed")
		return c.Redirect(routes.Application.Index())
	}

	if err != nil {
		fmt.Println(err)
	}

}
