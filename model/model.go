package model //模型表，用来将结构体与数据库的表进行一个映射

// "time"

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
	UserId      string `json:"users_id" gorm:"column:users_id"`
	Avatar      string `json:"avatar" gorm:"column:avatar"`
	Sha         string `json:"sha" gorm:"column:sha"`
	Path        string `json:"path" gorm:"column:path"`
	NickName    string `json:"nickname" gorm:"column:nickname"`
	Declaration string `json:"declaration" gorm:"column:declaration"`
	Infor       string `json:"infor" gorm:"column:infor"`
	Status      string `json:"status" gorm:"column:status"` //记录是否有同桌
	Tag1        string `json:"tag1" gorm:"column:tag1"`
	Tag2        string `json:"tag2" gorm:"column:tag2"`
	Tag3        string `json:"tag3" gorm:"column:tag3"`
	Tag4        string `json:"tag4" gorm:"column:tag4"`
	Tag5        string `json:"tag5" gorm:"column:tag5"`
	College     string `json:"college" gorm:"column:college"` //学院
	Major       string `json:"major" gorm:"column:major"`     //专业
	Grade       string `json:"grade" gorm:"column:grade"`     //年级
}

type Dailyrecords struct {
	Id      int    `json:"id" gorm:"column:id"`
	UserId1 string `json:"users_id1" gorm:"column:users_id1"`
	UserId2 string `json:"users_id2" gorm:"column:users_id2"`
	Time    int    `json:"time" gorm:"column:time"`
	// Daily   int    `json:"daily" gorm:"column:daily"` 这里因为我把消息重新弄了个表，忘记删了 // 3.22
	Status string `json:"status" gorm:"column:status"` //记录此次同桌是否结束
}

type Message struct {
	Id             int    `json:"id" gorm:"column:id"`
	DailyrecordsId int    `json:"dailyrecords_id" gorm:"column:dailyrecords_id"`
	UserId         string `json:"user_id" gorm:"column:user_id"`
	Time           string `json:"time" gorm:"column:time"`
	Information    string `json:"information" gorm:"column:information"`
}

type ReturnMessage struct {
	Message
	Name string `json:"name"`
}
type Applycation struct {
	Id      int    `json:"id" gorm:"column:id"`
	UserId1 string `json:"users_id1" gorm:"column:users_id1"`
	UserId2 string `json:"users_id2" gorm:"column:users_id2"`
	Result  string `json:"result" gorm:"column:result"`
}

type Update struct {
	Id             int `json:"id" gorm:"column:id"`
	DailyrecordsId int `json:"dailyrecords_id" gorm:"column:dailyrecords_id"`
	Time           string `json:"time" gorm:"column:time"`
}

/* type UserAndSign struct {
	Id      int    `json:"id" gorm:"column:id"`
	UserId1 string `json:"users_id1" gorm:"column:users_id1"`
	UserId2 string `json:"users_id2" gorm:"column:users_id2"`
	SignID  int    `json:"signs_id" gorm:"column:signs_id"`
} */

/*type Tag struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"tags_name" gorm:"column:tags_name"`
}*/

/* type CardAndTag struct {
	Id     int `json:"id" gorm:"column:id"`
	CardId int `json:"cards_id" gorm:"column:cards_id"`
	TagsId int `json:"tags_id" gorm:"column:tags_id"`
} */
