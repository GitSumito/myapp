package controllers

import (
    "github.com/revel/revel"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "log"
    "fmt"
//    "github.com/myaccount/myapp/app/model"
)

type Login struct {
	*revel.Controller
}

type User struct {
    Id   int    `db:"id"`
    Name string `db:"name"`
}

func (c Login) Index(Id string, Password int) revel.Result {

        fmt.Printf("%#v\n%#v\n", Id, Password)
        d := UserLogin(Id, Password) 
        
	return c.Render(d)
}


func UserLogin(Id string, Password int) []User {

    db, err := connect()
    if err != nil {
        log.Fatalln(err)
    }

    people := []User{}
    err = db.Select(&people, "SELECT id, name From patient where id = ?;", Id)
    if err != nil {
        fmt.Println(err)
    }

    for i :=0; i <len(people); i++ {
        fmt.Println(people[i], " ", people[i].Id, " ", people[i].Name, " ")
    }

    return people
}

func connect() (*sqlx.DB, error) {
    return sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
}
