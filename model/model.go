package model //模型表，用来将结构体与数据库的表进行一个映射

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql users表
type User struct {
	Id        int    `json:"id" gorm:"column:id"`
	StudentID string `json:"student_id" gorm:"column:student_id"`
	PassWord  string `json:"password" gorm:"column:password"`
	Name      string `json:"name" gorm:"column:name"`
	College   string `json:"college" gorm:"column:college"`
	Grade     string `json:"grade" gorm:"column:grade"`
	//Major     string `json:"major" gorm:"column:major"`
}

type Card struct {
	Id          int    `json:"id" gorm:"column:id"`
	UserId      int    `json:"users_id" gorm:"column:users_id"`
	Avatar      string `json:"avatar" gorm:"column:avatar"`
	NickName    string `json:"nickname" gorm:"column:nickname"`
	Declaration string `json:"declaration" gorm:"column:declaration"`
	Infor       string `json:"infor" gorm:"column:infor"`
}

type Sign struct {
	Id    int    `json:"id" gorm:"column:id"`
	Time  int    `json:"time" gorm:"column:time"`
	Daily string `json:"daily" gorm:"column:daily"`
}

type UserAndSign struct {
	Id      int `json:"id" gorm:"column:id"`
	UserId1 int `json:"users_id1" gorm:"column:users_id1"`
	UserId2 int `json:"users_id2" gorm:"column:users_id2"`
	SignID  int `json:"signs_id" gorm:"column:signs_id"`
}

type Apply struct {
	Id      int `json:"id" gorm:"column:id"`
	UserId1 int `json:"users_id1" gorm:"column:users_id1"`
	UserId2 int `json:"users_id2" gorm:"column:users_id2"`
}

type Tag struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"tags_name" gorm:"column:tags_name"`
}

type CardAndTag struct {
	Id     int `json:"id" gorm:"column:id"`
	CardId int `json:"cards_id" gorm:"column:cards_id"`
	TagsId int `json:"tags_id" gorm:"column:tags_id"`
}
