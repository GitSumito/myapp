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

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}

func (c TsukadaApp) Index() revel.Result {
	return c.Render()
}


func (c TsukadaApp) Apple() revel.Result {

    db, err := sqlx.Connect("mysql", "root:passwordPASSWORDpasswordPASSWORD@(localhost:3306)/booking")
    if err != nil {
        log.Fatalln(err)
    }

    people := []User{}
    err = db.Select(&people, "SELECT id, name From patient;")
    if err != nil {
        fmt.Println(err)
    }

    jason, john := people[0], people[1]

    fmt.Printf("%#v\n%#v\n", jason, john)


//    defer rows.Close()
    for people.Next() {
        user := User{}
//        err = rows.StructScan(&user)


        log.Println(user)
        log.Println(err)
    }
    return c.Render(jason)
}


