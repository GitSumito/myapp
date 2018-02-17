package controllers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/revel/revel"
	"log"
	//    "github.com/myaccount/myapp/app/model"
)

type Login struct {
	*revel.Controller
}

type Adduser struct {
	Id string `db:"id"`
	Name string  `db:"name"`
	Phone string `db:"phone"`
	Birthday string `db:"birthday"`
        Created_at string `db:"created_at"`
        Updated_at string `db:"updated_at"`
}

// 戻り値削除
func (c Login) Index(Id string, Name string, Phone String,Birthday string) {


	c.Validation.Required(Id).Message("診察券番号を入力してください")
	c.Validation.Required(Password).Message( "患者名を入力してください")
	c.Validation.Required(Phone).Message( "電話番号を入力してください")
	c.Validation.Required(Birthday).Message( "誕生日を入力してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	d := Useradd(Id, Name, Phone, Birthday)

	return c.Render(d)
}

func Useradd(Id string, Name string, Phone string, Birthday string) []User {

	db, err := connect()
	if err != nil {
		log.Fatalln(err)
	}

	people := []User{}
	err = db.Select(&people, "insert into  patient (id, name, phone, birthday, created_at, updated_at) VALUES (?, ?, ?, ?, now(), now();", Id, Name, Phone,Birthday)
	if err != nil {
		fmt.Println(err)
	}

}

func connect() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
}

