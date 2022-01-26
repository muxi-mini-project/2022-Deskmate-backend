package model

import (
	"fmt"
)

func GetUserInfo(UserId string) (User, error) {
	var user User
	if err := DB.Table("users").Where("student_id=?", UserId).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//修改名片信息
func ChangeCardInfo(card Card) error {
	fmt.Println(card.UserId)
	if err := DB.Table("cards").Where("users_id=?", card.UserId).Updates(map[string]interface{}{"avatar": card.Avatar, "nickname": card.NickName, "declaration": card.Declaration, "infor": card.Infor, "tag1": card.Tag1, "tag2": card.Tag2, "tag3": card.Tag3, "tag4": card.Tag4, "tag5": card.Tag5}).Error; err != nil {
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
	if err := DB.Table("cards").Where("status=?", "无").Find(&cards).Error; err != nil {
		return []Card{}, err
	}
	return cards, nil
}

//搜索tag查询名片
func GetCardByTag(Tag string) ([]Card, error) {
	var cards []Card
	if err := DB.Table("cards").Where("tag1 in(?) or tag2 in(?) or tag3 in(?) or tag4 in(?) or tag5 in(?)", Tag, Tag, Tag, Tag, Tag).Find(&cards).Error; err != nil {
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
