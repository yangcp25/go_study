package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Comments struct {
	Id         int64  `json:"id" orm:"column(id);auto;pk"`
	Content    string `json:"content" orm:"column(content)"`
	UserId     int    `orm:"column(userid)"`
	CreateTime string `orm:"column(create_time)"`
}

func init() {
	orm.RegisterModel(new(Comments))
}

func SaveComments(content string, userid int) (int64, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var comments Comments
	comments.Content = content
	comments.UserId = userid
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	comments.CreateTime = tm.Format("2006-01-02 03:04:05")

	id, err := o.Insert(&comments)
	return id, err
}

func EditComments(id int64, content string) (int64, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var comments Comments
	comments.Id = id

	if o.Read(&comments) == nil {
		if content != "" {
			comments.Content = content
			to, err := o.Begin()
			if err != nil {
				return 0, err
			}
			if res, err := o.Update(&comments, "content"); err != nil {
				to.Rollback()
				return res, errors.New("修改失败")
			} else {
				to.Commit()
				return res, nil
			}
		}
	}

	return 0, nil
}

func DeleteComments(content string) (int64, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var comments Comments
	comments.Content = content

	err := o.QueryTable(comments).Filter("content", content).One(&comments)

	if err == nil {
		if res, err := o.QueryTable(comments).Filter("content", content).Delete(); err != nil {
			return res, errors.New("删除失败")
		} else {
			return res, nil
		}
	}

	return 0, nil
}
