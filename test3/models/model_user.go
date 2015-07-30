package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int64
	Username string    `xorm:"not null unique`
	Password string    `xorm:"not null`
	Created  time.Time `xorm:"not null created"`
}

func (t *User) Info() string {
	return fmt.Sprintf("%#v", t)
}

func GetUserById(id int64) *User {
	user := new(User)
	Engine.Where("id = ?", id).Get(user)
	return user
}

func GetUserByName(name string) *User {
	user := new(User)
	Engine.Where("username = ?", name).Get(user)
	return user
}

func GetUsers() (users []User) {
	Engine.Asc("id").Find(&users)
	return users
}

func SetUser(a *User) bool {

	user := new(User)
	Engine.Where("username =  ?", a.Username).Get(user)
	if user.Id != 0 {
		return false
	}

	_, err := Engine.Insert(a)
	if err == nil {
		return true
	}
	return false
}

func UpdUser(id int64, user map[string]interface{}) bool {
	_, err := Engine.Table(new(User)).Id(id).Update(user)
	if err == nil {
		return true
	}
	return false
}

func DelUserById(id int64) bool {
	user := new(User)
	_, err := Engine.Id(id).Delete(user)
	if err == nil {
		return true
	}
	return false
}
