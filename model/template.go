package model

//用来规范与前端对接时发送的信息模板

type Log struct {
	StudentID string `json:"student_id" gorm:"column:student_id"`
	PassWord  string `json:"password" gorm:"column:password"`
}

// 这个结构体用于规范前端发送的结构体
type Cardinfor struct{
	NickName    string `json:"nickname" gorm:"column:nickname"`
	Declaration string `json:"declaration" gorm:"column:declaration"`
	Infor       string `json:"infor" gorm:"column:infor"`
	Tag1        string `json:"tag1" gorm:"column:tag1"`
	Tag2        string `json:"tag2" gorm:"column:tag2"`
	Tag3        string `json:"tag3" gorm:"column:tag3"`
	Tag4        string `json:"tag4" gorm:"column:tag4"`
	Tag5        string `json:"tag5" gorm:"column:tag5"`
	Collge      string `json:"college" gorm:"column:college"`
	Major       string `json:"major" gorm:"column:major"`
	Grade  		string `json:"grade" gorm:"column:grade"`
}

type ApplicantCard struct{
	NickName    string `json:"nickname" gorm:"column:nickname"`
	Declaration string `json:"declaration" gorm:"column:declaration"`
	Infor       string `json:"infor" gorm:"column:infor"`
	Tag1        string `json:"tag1" gorm:"column:tag1"`
	Tag2        string `json:"tag2" gorm:"column:tag2"`
	Tag3        string `json:"tag3" gorm:"column:tag3"`
	Tag4        string `json:"tag4" gorm:"column:tag4"`
	Tag5        string `json:"tag5" gorm:"column:tag5"`
	Avatar      string `json:"avatar" gorm:"column:avatar"`
}


type Search struct {
	Tag string `json:"tag"`
} 

type Tag struct{
	Tag        string `json:"tag" gorm:"column:tag"`
}

// 这个结构体是每日打卡的内容
type Record struct { 
	Information    string    `json:"information" gorm:"column:information"`
}

type Id struct{
	Id string `json:"users_id"`
}

/* type Response struct{
	Id string `json:"id"`
} */

type Respondent struct {
	StudentID string `json:"respondent_id"`
}

type UserInfor struct {
	Id        int    `json:"id" gorm:"column:id"`
	StudentID string `json:"student_id" gorm:"column:student_id"`
	Name      string `json:"name" gorm:"column:name"`
	College   string `json:"college" gorm:"column:college"`
	Grade     string `json:"grade" gorm:"column:grade"`
	//Major     string `json:"major" gorm:"column:major"`
}