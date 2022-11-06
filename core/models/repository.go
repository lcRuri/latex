package models

import "time"

type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Name      string
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository"
}
