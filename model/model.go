package model //模型表，用来将结构体与数据库的表进行一个映射

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql users表
type User struct {
	Id    int    `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`
	Collge string `json:"collge" gorm:"collge"`
	Major string `json:"major" gorm:"major"`
	Grade string `json:"grade" gorm:"gorm"`
}

type Card struct {
	Id          int    `json:"id" gorm:"id"`
	UserId      int    `json:"users_id" gorm:"users_id"`
	Avatar      string `json:"avatar" gorm:"avatar"`
	NickName    string `json:"nickname" gorm:"nickname"`
	Declaration string `json:"declaration" gorm:"declaration"`
	Infor       string `json:"infor" gorm:"infor"`
}

type Sign struct {
	Id    int    `json:"id" gorm:"id"`
	Time  int    `json:"time" gorm:"time"`
	Daily string `json:"daily" gorm:"daily"`
}

type UserAndSign struct {
	Id      int `json:"id" gorm:"id"`
	UserId1 int `json:"users_id1" gorm:"users_id1"`
	UserId2 int `json:"users_id2" gorm:"users_id2"`
	SignID  int `json:"signs_id" gorm:"signs_id"`
}

type Apply struct {
	Id      int `json:"id" gorm:"id"`
	UserId1 int `json:"users_id1" gorm:"users_id1"`
	UserId2 int `json:"users_id2" gorm:"users_id2"`
}

type Tag struct {
	Id   int    `json:"id" gorm:"id"`
	Name string `json:"tags_name" gorm:"tags_name"`
}

type CardAndTag struct {
	Id int `json:"id" gorm:"id"`
	CardId int `json:"cards_id" gorm:"cards_id"`
	TagsId int `json:"tags_id" gorm:"tags_id"`
}
