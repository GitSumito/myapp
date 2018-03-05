package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"log"
)

type DaySchedule struct {
	Datetime      string `db:"datetime"`
	Room          string `db:"room"`
	Treat_time    string `db:"treat_time"`
	Dr_no         string `db:"dr_no"`
	Assist_no     string `db:"assist_no"`
	Dr_weight     string `db:"dr_weight"`
	Assist_weight string `db:"assist_weight"`
	Patient_id    string `db:"patient_id"`
	Dr_name       string `db:"dr_name"`
	Patient_name  string `db:"patient_name"`
	Assist_name   string `db:"assist_name"`
}

//type ReservationAdmin struct {
//	Datetime      string `db:"datetime"`
//	Day           string `db:"day"`
//	Time          string `db:"time"`
//	Viewtime      string `db:"viewtime"`
//	Room          string `db:"room"`
//	Treat_time    string `db:"treat_time"`
//	Dr_no         string `db:"dr_no"`
//	Assist_no     string `db:"assist_no"`
//	Dr_weight     string `db:"dr_weight"`
//	Assist_weight string `db:"assist_weight"`
//	Patient_id    string `db:"patient_id"`
//	Created_at    string `db:"created_at"`
//	Updated_at    string `db:"updated_at"`
//}

type Admin struct {
	*revel.Controller
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) Menu() revel.Result {
	return c.Render()
}

func (c Admin) Adduser(Id string, Name string, Phone string, Birthday string) revel.Result {

	c.Validation.Required(Id).Message("診察券番号を入力してください")
	c.Validation.Required(Name).Message("患者名を入力してください")
	c.Validation.Required(Phone).Message("電話番号を入力してください")
	c.Validation.Required(Birthday).Message("誕生日を入力してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		//return c.Redirect(Admin.Adduser)
		return c.Redirect(Admin.Menu)
	}

	// userが既に登録されていた場合の処理
	//    people := []UserInfomation{}
	//    err = db.Select(&people, "SELECT id, name From patient;")
	//    if err != nil {
	//        fmt.Println(err)
	//    }
	//
	//    return c.Render(people)

	// ユーザーを追加する処理
	Useradd(Id, Name, Phone, Birthday)

	return c.Render()
}

func Useradd(Id string, Name string, Phone string, Birthday string) {

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Id, Name, Phone, Birthday)
	tx := db.MustBegin()

	query := `insert into patient (id, name, phone, birthday, created_at, updated_at) VALUES (?, ?, ?, ?, now(), now())`
	fmt.Println(query)

	db.MustExec(query, Id, Name, Phone, Birthday)
	tx.Commit()
}

func (c Admin) List(day string) revel.Result {

	fmt.Println("Admin.Date")
	dshedule := []DaySchedule{}

	yyyy := day[:4]
	mm := day[4:6]
	dd := day[6:]

	ymd := yyyy + ":" + mm + ":" + dd

	separator := " "

	starttime := "00:00:00"
	endtime := "23:59:59"

	start := "\"" + ymd + separator + starttime + "\""
	fmt.Println(start)

	end := "\"" + ymd + separator + endtime + "\""
	fmt.Println(end)

	SQL := "select " +
		"res.datetime, " +
		"res.room, " +
		"res.treat_time, " +
		"res.dr_no, " +
		"res.assist_no, " +
		"res.dr_weight, " +
		"res.assist_weight, " +
		"res.patient_id, " +
		"s.name as dr_name, " +
		"p.name as patient_name, " +
		"a.name as assist_name " +
		"from reservation as res " +
		"left outer join staff as s on res.dr_no = s.id " +
		"left outer join patient as p on res.patient_id = p.id " +
		"left outer join assist as a on res.assist_no = a.id " +
		" where datetime > " + start +
		" and  datetime < " +
		end + " order by datetime;"
	fmt.Println(SQL)

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Select(&dshedule, SQL)

	fmt.Println(len(dshedule))

	for i := 0; i < len(dshedule); i++ {
		fmt.Println(dshedule[i], " ", dshedule[i].Datetime, " ", dshedule[i].Room, " ")
	}

	return c.Render(dshedule)
}
