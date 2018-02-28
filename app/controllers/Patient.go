package controllers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"log"
	"time"
)

type PatientStruct struct {
	Id         string `db:"id"`
	Name       string `db:"name"`
	Phone      string `db:"phone"`
	Birthday   string `db:"birthday"`
	Created_at string `db:"created_at"`
	Updated_at string `db:"updated_at"`
}

type Reservation struct {
	Datetime      string `db:"datetime"`
	Room          string `db:"room"`
	Treat_time    string `db:"treat_time"`
	Dr_no         string `db:"dr_no"`
	Assist_no     string `db:"assist_no"`
	Dr_weight     string `db:"dr_weight"`
	Assist_weight string `db:"assist_weight"`
	Patient_id    string `db:"patient_id"`
	Created_at    string `db:"created_at"`
	Updated_at    string `db:"updated_at"`
}

type Patient struct {
	*revel.Controller
}

func (c Patient) Index() revel.Result {
	return c.Render()
}

// ログイン処理を行う
// クッキーとセッションを保存する
// 開き時間を返す
func Choose(Id string) []Reservation {
	// ユーザーを追加する処理
	fmt.Print("*** Choose: " + Id)

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	bookableSlot := []Reservation{}

	// #現在時間+1時間
	// TODO ユーザーごとの処理を追加する
	ltime := time.Now()
	const layout = "2006-01-02 15:04:05"
	fmt.Println(ltime.Format(layout))

	datetime := ltime.Add(1 * time.Minute)
	iikanji := datetime.Format(layout)

	SQL := "select * from reservation where datetime >= " + "\"" + iikanji + "\"" + " and patient_id = 0 order by datetime limit 3;"
	fmt.Println(SQL)
	err = db.Select(&bookableSlot, SQL)

	for i := 0; i < len(bookableSlot); i++ {
		fmt.Println(bookableSlot[i], " ", bookableSlot[i].Datetime, " ", bookableSlot[i].Room, " ")
	}

	return bookableSlot
}

// ID/PW を元にログインする
// TODO セッションを登録する Cookieを返す
func (c Patient) Login(id string, password string, remember bool) revel.Result {

	db, err := Connect()

	if err != nil {
		c.Flash.Out["id"] = id
		c.Flash.Error("Connection failed")

		log.Fatalln(err)
		return c.Redirect("/")
	}

	people := []PatientStruct{}

	SQL := "SELECT * From patient where id = " + id + ";"

	fmt.Println("** " + SQL + " **")
	err = db.Select(&people, SQL)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i], " ", people[i].Id, " ", people[i].Name, " ")
	}

	if len(people) == 0 {
		fmt.Println("data nothing")
		c.Flash.Error("Login failed")
		return c.Redirect("/")
	} else {

		fmt.Println("OK")
	}

	vacantTime := Choose(id)
	fmt.Println(len(vacantTime))

	return c.Render(vacantTime)
}
