package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	// "gorm.io/gorm" // 3.24
)

func GetUserInfo(UserId string) (UserInfor, error) {
	var user UserInfor
	if err := DB.Table("users").Where("student_id=?", UserId).Find(&user).Error; err != nil {
		return UserInfor{}, err
	}

	return user, nil
}

//修改名片信息
func ChangeCardInfo(card Card) error {
	fmt.Println(card.UserId)
	// if err := DB.Table("cards").Where("users_id=?", card.UserId).Updates(map[string]interface{}{"avatar": card.Avatar, "nickname": card.NickName, "declaration": card.Declaration, "infor": card.Infor, "tag1": card.Tag1, "tag2": card.Tag2, "tag3": card.Tag3, "tag4": card.Tag4, "tag5": card.Tag5}).Error; err != nil {
	// 删除了更新头像(写在单独的api里面)，增加了年级，学院和专业
	if err := DB.Table("cards").Where("users_id=?", card.UserId).Updates(map[string]interface{}{"nickname": card.NickName, "declaration": card.Declaration, "infor": card.Infor, "tag1": card.Tag1, "tag2": card.Tag2, "tag3": card.Tag3, "tag4": card.Tag4, "tag5": card.Tag5, "grade": card.Grade, "college": card.College, "major": card.Major}).Error; err != nil {
		return err //应该是数据库的
	}
	return nil
}

//查询对应学号的名片
func GetCardInfo(UserId string) (Card, error) {
	var card Card
	if err := DB.Table("cards").Where("users_id=?", UserId).Find(&card).Error; err != nil {
		return Card{}, err
	}
	return card, nil
}

//查询全部名片记录
func GetCardsInfo() ([]Card, error) {
	var cards []Card
	if err := DB.Table("cards").Where("status=?", "0").Find(&cards).Error; err != nil {
		return []Card{}, err
	}
	return cards, nil
}

//搜索tag查询名片
func GetCardByTag(Tag string) ([]Card, error) {
	var cards []Card
	//if err := DB.Table("cards").Where("tag1 in(?) or tag2 in(?) or tag3 in(?) or tag4 in(?) or tag5 in(?)", Tag, Tag, Tag, Tag, Tag).Find(&cards).Error; err != nil {
	if err := DB.Table("cards").Where("tag1 LIKE(?) or tag2 LIKE(?) or tag3 LIKE(?) or tag4 LIKE(?) or tag5 LIKE(?) and status=(?)", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%", "0").Find(&cards).Error; err != nil {
		return []Card{}, err
	}
	//"tag1 in(?) or tag2 in(?) or tag3 in(?) or tag4 in(?) or tag5 in(?)"
	return cards, nil
}

//查看申请对象此时是否有同桌
func GetUserStatus(UserId string) (string, error) {
	var status string
	if err := DB.Table("cards").Where("users_id=?", UserId).Select("status").Find(&status).Error; err != nil {
		return "查询申请对象是否有同桌失败", err
	}
	return status, nil
}

//查询自己收到的申请
func GetMyApply(UserId string) ([]Applycation, error) {
	var apply []Applycation
	if err := DB.Table("applycations").Where("users_id2=? and result=?", UserId, " ").Find(&apply).Error; err != nil { //查询记录时忽略已经有结果的即result不为空的记录
		return []Applycation{}, err
	}
	return apply, nil
}

func GetMyApplyUserId1(UserId string) ([]string, error) {
	var apply []string
	if err := DB.Table("applycations").Where("users_id2=? and result=?", UserId, " ").Select("users_id1").Find(&apply).Error; err != nil { //查询记录时忽略已经有结果的即result不为空的记录
		return nil, err
	}
	return apply, nil
}

//确认同意收到的申请后将该申请的结果(result)设为同意(1)
func ConfirmApplication(UserId string, RespondentId string) error {
	if err := DB.Table("applycations").Where("users_id1 in(?) and users_id2 in (?)", RespondentId, UserId).Update("result", "1").Error; err != nil {
		return err
	}
	return nil
}

//确认同意收到的申请后将该申请的结果(result)设为拒绝(0)
func RefuseApplication(UserId string, RespondentId string) error {
	if err := DB.Table("applycations").Where("users_id1 in(?) and users_id2 in (?)", RespondentId, UserId).Update("result", "0").Error; err != nil {
		return err
	}
	return nil
}

//成为同桌后将两人的状态改为有同桌
func ChangeStatus(UserId string) error {
	if err := DB.Table("cards").Where("users_id=?", UserId).Update("status", "1").Error; err != nil {
		return err
	}
	return nil
}

//成为同桌后将两人的状态改为无同桌
func ChangeStatusAgain(UserId string) error {
	if err := DB.Table("cards").Where("users_id=?", UserId).Update("status", "0").Error; err != nil {
		return err
	}
	return nil
}

//通过自己的id来查询当前正在进行中的同桌关系的id
func GetDeskmateId(id string) (int, error) {
	var DeskmateID int
	if err := DB.Table("dailyrecords").Where("users_id1 in (?) or users_id2 in (?) and status in (?) ", id, id, "进行中").Select("id").Find(&DeskmateID).Error; err != nil {
		return 0, err
	}
	return DeskmateID, nil
}

//通过同桌关系记录表的id来查询相关的信息
func GetMessage(id int) ([]Message, error) {
	var message []Message
	if err := DB.Table("messages").Where("dailyrecords_id = ?", id).Find(&message).Error; err != nil {
		return []Message{}, err
	}
	return message, nil
}

//添加了Name识别对方与自己
func GetReturnMessage(id int) ([]ReturnMessage, error) {
	var message []ReturnMessage
	if err := DB.Table("messages").Where("dailyrecords_id = ?", id).Find(&message).Error; err != nil {
		return []ReturnMessage{}, err
	}
	return message, nil
}

//这个函数返回一条消息
func GetLastMessgae(id int) (Message, error) {
	var message Message
	if err := DB.Table("messages").Where("dailyrecords_id = ?", id).Last(&message).Error; err != nil {
		return message, err
	}
	return message, nil
}

//解除关系后对应的同桌记录状态变为“已结束”
func ChangeDeskmateStatus(id int) error {
	if err := DB.Table("dailyrecords").Where("id=?", id).Update("status", "已结束").Error; err != nil {
		return err
	}
	return nil
}

//通过同桌关系记录表的id来查询两位同学的id用于后续操作
func GetPartnerId(id int) (string, string, error) {
	var sid1, sid2 string
	if err := DB.Table("dailyrecords").Where("id=?", id).Select("users_id1", "users_id2").Find(&sid1, &sid2).Error; err != nil {
		return sid1, sid2, err
	}
	return sid1, sid2, nil
}

//通过同桌关系记录表的id来查询两位同学的id用于后续操作
func GetPartner(DeskmateId int, id string) (string, error) {
	var sid string
	if err := DB.Table("dailyrecords").Where("id=?", DeskmateId).Select("users_id1").Find(&sid).Error; err != nil {
		return strconv.Itoa(DeskmateId), err
	}
	if sid == id {
		if err := DB.Table("dailyrecords").Where("id=?", DeskmateId).Select("users_id2").Find(&sid).Error; err != nil {
			return strconv.Itoa(DeskmateId), err
		}
		return sid, nil
	}
	return sid, nil
}

//判断时间差，根据两人是否每天都发送一条信息来判读同桌关系是否继续，如果昨天两人都发表了一条消息，那么天数加1
func GudgeTime(DeskmateId int) string {
	//找寻上一条消息距离现在的时间差,如果超过一天以上，返回false,否则返回ture
	var result string
	var message Message
	if err := DB.Table("messages").Where("dailyrecords_id=?", DeskmateId).Last(&message).Error; err != nil {
		return "查询上一条记录失败"
	}
	//timestr := "2020-05-01 12:12:12"
	/* timestr := message.Time
	time1 := timestr.Format("2006-01-02 15:04:05") //go语言的时间起点，用于转化时间成字符串
	t1 := strings.Split(time1, " ")                //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24 */

	time1 := message.Time
	t1 := strings.Split(time1, " ") //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24
	if dt > 1 {
		result = "false"
	} else {
		result = "ture"
	}
	return result
}

//判断你今天是否已经发送过信息,如果已经发送过则返回false,没有则返回true
func GudgeSendRepeat(UserId string, DeskmateId int) (string, error) {
	//找寻该用户上一条消息距离现在的时间差,如果超过一天以上，返回false,否则返回ture
	var result string
	var message Message
	// if err := DB.Table("messages").Where("dailyrecords_id=? and user_id1 =? or users_id2 = ?", DeskmateId, UserId, UserId).Find(&message).Error; err != nil {
	err := DB.Table("messages").Where("dailyrecords_id=? and user_id =? ", DeskmateId, UserId).Last(&message).Error
	if err != nil { // 3.24 把Find改为Last，获取最新的一条数据
		/* if gorm.ErrRecordNotFound == err {
			return "true",nil
		}
		return "查询上一条记录失败", err */
		err = nil
	}
	//timestr := "2020-05-01 12:12:12"
	/* timestr := message.Time
	time1 := timestr.Format("2006-01-02 15:04:05") //go语言的时间起点，用于转化时间成字符串
	t1 := strings.Split(time1, " ")                //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24 */

	time1 := message.Time
	t1 := strings.Split(time1, " ") //Split函数是处理解析字符串的一个函数，将time1以空格分割成字符串数组
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24                                                        //func (t Time) Sub(u Time) Duration 返回一个时间段t-u;Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年
	if dt == 0 {
		result = "false"
	} else {
		result = "ture"
	}
	return result, nil
}

//判断你昨天天是否已经发送过信息,如果已经发送过则返回true,没有则返回false
func GudgeYesterday(UserId string, DeskmateId int) (string, error) {
	//找寻该用户上一条消息距离现在的时间差,如果超过一天以上，返回false,否则返回ture
	var result string
	var message Message
	
	if err := DB.Table("messages").Where("dailyrecords_id=? and user_id =? ", DeskmateId, UserId).Last(&message).Error; err != nil {
		result = "ture"
		return result, nil
		// return "查询上一条记录失败", err
	}
	//timestr := "2020-05-01 12:12:12"
	/* timestr := message.Time
	time1 := timestr.Format("2006-01-02 15:04:05") //go语言的时间起点，用于转化时间成字符串
	t1 := strings.Split(time1, " ")                //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24 */

	time1 := message.Time
	t1 := strings.Split(time1, " ") //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24
	if dt <= 1  {
		result = "ture"
	} else {
		result = "false"
	}
	return result, nil
}

//查询打卡天数
func GetDays(DeskmateId int) (int, error) {
	var day int
	if err := DB.Table("dailyrecords").Where("id=?", DeskmateId).Select("time").Find(&day).Error; err != nil {
		return -1, err
	}
	return day, nil
}

//更新打卡天数
func UpdateDays(DeskmateId int, Days int) error {
	if err := DB.Table("dailyrecords").Where("id=?", DeskmateId).Update("time", Days).Error; err != nil {
		return err
	}
	return nil
}


//修改用户头像
/* func UpdateAvator(avatar Card) error {
	if err := DB.Table("cards").Where("users_id= ?", avatar.UserId).Updates(map[string]interface{}{"avatar": avatar.Avatar, "sha": avatar.Sha, "path": avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}  */

func UpdateAvator(avatar Card) error {
	if err := DB.Table("cards").Where("users_id = ?", avatar.UserId).Updates(Card{Avatar: avatar.Avatar, Sha: avatar.Sha, Path: avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}

//判断今天是否已经更新了打卡
func GudgeUpdateRepeat(DeskmateId int) (string, error) {
	//找寻上一条更新距离现在的时间差,如果今天已经更新，则不再更新
	var result string
	var update Update
	err := DB.Table("updates").Where("dailyrecords_id=? ", DeskmateId).Last(&update).Error
	if err != nil { // 3.24 把Find改为Last，获取最新的一条数据
		
		return "更新打卡失败", err 

	}
	//timestr := "2020-05-01 12:12:12"
	/* timestr := message.Time
	time1 := timestr.Format("2006-01-02 15:04:05") //go语言的时间起点，用于转化时间成字符串
	t1 := strings.Split(time1, " ")                //Split函数是处理解析字符串的一个函数，
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24 */

	time1 := update.Time
	t1 := strings.Split(time1, " ") //Split函数是处理解析字符串的一个函数，将time1以空格分割成字符串数组
	s1 := t1[0] + " 00:00:00"
	t2, _ := time.Parse("2006-01-02 15:04:05", s1)
	t3, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00")) //当前时间
	dt := t3.Sub(t2).Hours() / 24                                                        //func (t Time) Sub(u Time) Duration 返回一个时间段t-u;Duration类型代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年
	if dt == 0 {
		result = "0"//返回0说明今天已经更新过
	} else {
		result = "1"//返回1说明今天没有更新过
	}
	return result, nil
}