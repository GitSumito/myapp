package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"log"
)

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
	c.Validation.Required(Name).Message( "患者名を入力してください")
	c.Validation.Required(Phone).Message( "電話番号を入力してください")
	c.Validation.Required(Birthday).Message( "誕生日を入力してください")

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

