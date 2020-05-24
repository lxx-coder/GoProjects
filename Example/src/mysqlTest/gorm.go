package mysqlTest

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "github.com/spaolacci/murmur3"
    "strings"
    "time"
)

type Like struct {
    Id             int     `gorm:"primary_key"`
    Ip             string  `gorm:"type:varchar(20);not null;index:ip_idx"`
    Ua             string  `gorm:"type:varchar(256);not null"`
    Title          string  `gorm:"type:varchar(128);not null;index:title_idx"`
    Hash           uint64  `gorm:"unique_index:hash_idx"`
    CreateAt       time.Time
}

var db *gorm.DB

func connect() {
    var err error
    db, err = gorm.Open("mysql","root:xx2019yy@tcp(localhost:3306)/test?" +
        "charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db.SingularTable(true)
}

func Create() {
    if db != nil && !db.HasTable(&Like{}) {
        if err := db.Set(
            "gorm:table_options",
            "ENGINE=InnoDB DEFAULT CHARSET=utf8",
            ).CreateTable(&Like{}).Error;err != nil {
            panic(err)
        }
    }
}

func Insert() {
    ip := "127.0.0.1"
    ua := "abcd1234"
    title := "example"
    like := &Like{
        Ip: ip,
        Ua: ua,
        Title: title,
        Hash: murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
        CreateAt: time.Now(),
    }

    if err := db.Create(like).Error; err != nil {
        panic(err)
    }
}


