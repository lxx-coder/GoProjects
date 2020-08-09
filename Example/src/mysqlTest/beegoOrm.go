package mysqlTest

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

type User struct {
    Id       int
    Name     string     `orm:"sizes(100)"`
}

func ConnectTable() {
    
    orm.RegisterDataBase("default", "mysql","root:xx921226@tcp(localhost:3306)/test?charset=utf8",30)

    // register model
    orm.RegisterModel(new(customer))

    // create table
    orm.RunSyncdb("default",false, true)

}

func TakeOrm() {
    o := orm.NewOrm()

    user := User{Name: "slene"}
    //insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, Err: %v\n", id, err)

    //update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM:%d, Err: %v\n", num, err)

    // read one
    u := User{Id:user.Id}
    err = o.Read(u)
    fmt.Printf("ERR: %v\n", err)
    //delete
    num, err = o.Delete(&u)
    fmt.Printf("NUM: %d, ERR: %d\n", num, err)

}
