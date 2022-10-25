package service

import (
	"encoding/json"

	"server.pluto.com/log"
	"server.pluto.com/mysql"
)

type AllowList struct {
	Id   int
	Name string
}

type User struct{}

func (user *User) QueryUserList() (userList string) {
	// sql := "select * from allow_list"
	// r, _ := mysql.Engine.QueryResult(sql).List()
	// s := fmt.Sprint(r)
	everyone := []AllowList{}
	err := mysql.Engine.Find(&everyone)
	if err != nil {
		log.Error("数据库查询错误", err)
	}
	b, err2 := json.Marshal(everyone)
	if err2 != nil {
		log.Error("查询结果转json错误", err)
	}
	result := string(b)
	return result
}

func (user *User) InsertUser(username string) (result string) {
	allowList := new(AllowList)
	allowList.Name = username
	_, err := mysql.Engine.Insert(allowList)
	if err != nil {
		log.Error("用户新增错误", err)
	}

	return "ok"
}

func (user *User) DelUser(username string) (result string) {
	// sql := "delete from 'allow_list' where name = '" + username + "'"
	// _, err := mysql.Engine.Exec(sql, 1)
	// if err != nil {
	// 	log.Error("用户删除错误", err)
	// 	return "删除用户失败"
	// }
	// return "ok"

	allowList := new(AllowList)
	_, err := mysql.Engine.ID(1).Delete(allowList)
	if err != nil {
		log.Error("用户删除错误", err)
		return "删除用户失败"
	}
	return "ok"

}
