package test

import (
	"log"
	"testing"
	"xorm.io/xorm"
)

type Model struct {
	Id      int
	Name    string
	Content string
}

func TestCon(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/latex?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("xorm.NewEngine err:%v", err)

	}

	md := new(Model)
	_, err = engine.Where("name=?", "njupt").Get(md)
	if err != nil {
		log.Fatal(err)
	}

}
