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

func FindUser(email string, password string) (int64, string) {
	var o orm.Ormer
	o = orm.NewOrm()

	var user Users
	err := o.QueryTable("users").Filter("email", email).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}
func GetUserList(limit int, offset int, remember_token int) (int64, []orm.Params, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var users []orm.Params

	query := o.QueryTable("users")
	query = query.Filter("remember_token", remember_token)
	query = query.OrderBy("-created_at")

	count, _ := query.Values(&users)

	query = query.Limit(limit, offset)
	_, err := query.Values(&users)
	return count, users, err
}
