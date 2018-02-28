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

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (c Login) Index(Id int, Password int) revel.Result {

	fmt.Printf("%#v\n%#v\n", Id, Password)

	c.Validation.Required(Id).Message("IDを入力してください")
	// c.Validation.MinSize(Id,3).Message("IDが正しいか確認してください")

	c.Validation.Required(Password).Message("電話番号を入力してください")
	// c.Validation.MinSize(Password, 3).Message( "電話番号が正しいか確認してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		// リダイレクト先を変えたい。とりあえずこのまま
		//	return c.Redirect(Patient.Index)
		return c.Render()
	}

	d := UserLogin(Id, Password)

	return c.Render()
}

func UserLogin(Id int, Password int) []User {

	db, err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	people := []User{}
	err = db.Select(&people, "SELECT id, name From patient where id = ?;", Id)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i], " ", people[i].Id, " ", people[i].Name, " ")
	}

	return people
}

func Connect() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
}
