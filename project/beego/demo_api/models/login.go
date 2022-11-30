package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Users struct {
	Id              int64     `json:"id" orm:"column(id);auto;pk"`
	Name            string    `json:"name" orm:"column(name)"`
	Email           string    `json:"email" orm:"column(email)"`
	EmailVerifiedAt time.Time `orm:"column(email_verified_at);auto_now;type(datetime)"`
	Password        string    `json:"password" orm:"column(password)"`
	RememberToken   string    `json:"remember_token" orm:"column(remember_token)"`
	CreatedAt       time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt       time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func (u *Users) Insert(userObj interface{}) error {
	var o orm.Ormer
	o = orm.NewOrm()
	_, err := o.Insert(userObj)
	return err
}

func (u *Users) Find() bool {
	var o orm.Ormer
	o = orm.NewOrm()
	err := o.Read(u)
	if err != orm.ErrNoRows {
		return false
	}
	return true
}
