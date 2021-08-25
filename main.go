package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	DB = Open()
}

func Open() *gorm.DB {

	host := "localhost"
	port := "5432"
	user := "postgres"
	//pass := ""
	dbname := "test01"
	sslmode := "disable"

	dsn := "host=" + host + " user=" + user + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
	fmt.Println("dsn string setting: " + dsn)

	mode := "GIN_MODE"
	config := &gorm.Config{}
	if mode != "release" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// type auth_ab struct {
// 	A int
// 	B string
// }

// func (auth_ab) TableName() string {
// 	return "auth_ab"
// }

// type Customer struct {
// 	base.BaseModel
// 	Name          string  json:"name"
// 	ContactPerson string  json:"contact_person"
// 	PhoneNumber   string  json:"phone_number"
// 	Address       string  json:"address"
// 	Square        float32 json:"square"
// 	PeopleAmount  int     json:"people_amount"
// 	Budget        int     json:"budget"
// 	BuildingType  int     json:"building_type"
// 	CreatAT       string  json:"creat_at"
//    }
//    type AddCustomer struct {
// 	Name          string  json:"name"
// 	ContactPerson string  json:"contact_person"
// 	PhoneNumber   string  json:"phone_number"
// 	Address       string  json:"address"
// 	Square        float32 json:"square"
// 	PeopleAmount  int     json:"people_amount"
// 	Budget        int     json:"budget"
// 	BuildingType  int     json:"building_type"
// 	CreatAT       string  json:"creat_at"
//    }
//    type UpdateCustomer struct {
// 	Name          string  json:"name"
// 	ContactPerson string  json:"contact_person"
// 	PhoneNumber   string  json:"phone_number"
// 	Address       string  json:"address"
// 	Square        float32 json:"square"
// 	PeopleAmount  int     json:"people_amount"
// 	Budget        int     json:"budget"
// 	BuildingType  int     json:"building_type"
// 	CreatAT       string  json:"creat_at"
//    }

type MisProperty struct {
	// BHMWbTx9JmAe3
	PID       string         `json:"coco" binding:"required"`
	JsonbT01  datatypes.JSON `json:"jsonb_t01" binding:"required"` //undeclared name: datatypes
	CreatedAt *time.Time     `json:"created_at"`
}

const (
	a = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
)

// 創造一個自己定義的結構，可能要跟表一樣名字
// 開頭要大寫，json後的才跟表一樣
// accounts資料庫的表
type account struct {
	ID              int64          `json:"id"`
	Pid             string         `json:"pid"`
	Info            string         `json:"info"`
	Cname           string         `json:"cname"`
	Jsonb_01        string         `json:"jsonb_01"`
	Account         string         `json:"account"`
	Passwd          string         `json:"passwd"`
	Updated_at      time.Time      `json:"updated_at"`
	Deleted_at      gorm.DeletedAt `json:"deleted_at"`
	Last_login_time time.Time      `json:"last_login_time"`
}

func main() {

	// Find：取得全部筆數
	// First：根據pk升冪取得第一筆
	// Last：根據pk降冪取得第一筆
	// Take：取得一筆資料

	// 註冊 finish
	// DB.Create(&account{
	// 	Pid:             "39",
	// 	Account:         "hhh",
	// 	Passwd:          "456",
	// 	Updated_at:      time.Now(),
	// 	Last_login_time: time.Now(),
	// })

	// 修改資料 update
	// 改密碼和更新時間 finish
	// DB.Model(&account{}).Where("pid = ? AND account = ? And passwd = ?", "40", "hhh", "567").Updates(account{Passwd: "456", Updated_at: time.Now()})
	// 改用map來做更新 non
	// DB.Model(&account{}).Where("pid = ? AND account = ? And passwd = ?", "40", "hhh", "456").Updates(map[string]string{"passwd": "567", "updated_at": time.Now().Format("2006-1-2 3:4:5")})

	// 插入資料 insert into
	// DB.Select("cname", "updated_at").Create(&account{Cname: "jjj", Updated_at: time.Now()})
	//INSERT INTO "accounts" ("cname","updated_at") VALUES ('jjj','2021-08-25 14:48:14.752') RETURNING "id"

	// 增刪
	// DB.Delete(&account{}).Where(account{Pid: "40"})

}
