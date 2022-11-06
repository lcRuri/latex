package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
	"xorm.io/xorm"
)

type UserBasic struct {
	Id        int
	Identity  string
	Name      string
	Password  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func TestMysql(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/latex?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("xorm.NewEngine err:%v", err)

	}

	u := new(UserBasic)
	u.Name = "latex"
	u.Password = "123456"
	_, err = engine.Insert(u)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	user := &UserBasic{
		Identity: "123456",
		Name:     "admin",
		Password: "123456",
	}

	engine, _ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/latex?charset=utf8mb4&parseTime=True&loc=Local")

	insert, err := engine.Insert(user)
	fmt.Println(insert, err)
}
