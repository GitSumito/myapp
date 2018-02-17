package controllers

import (
    "github.com/revel/revel"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "log"
    "fmt"
)

type TsukadaApp struct {
	*revel.Controller
}

type UserInfomation struct {
    Id   int    `db:"id"`
    Name string `db:"name"`
}

func (c TsukadaApp) Index() revel.Result {
	return c.Render()
}


func (c TsukadaApp) Apple() revel.Result {

    db, err := sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
    if err != nil {
        log.Fatalln(err)
    }

    people := []UserInfomation{}
    err = db.Select(&people, "SELECT id, name From patient;")
    if err != nil {
        fmt.Println(err)
    }

    jason, john := people[0], people[1]

    fmt.Printf("%#v\n%#v\n", jason, john)


    for i :=0; i <len(people); i++ {
        fmt.Println(people[i], " ", people[i].Id, " ", people[i].Name, " ")
    }


    return c.Render(people)
}


