// Package model
// @Title  groupuser.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:40
package model

import (
	"framework/db"
	"framework/logger"
)

type GroupUser struct {
	GroupId string `json:"group_id"`
	UID     string `json:"user_id"`
}

func NewGroupUser() *GroupUser {
	return &GroupUser{}
}

func InsertGroupUser(groupUser *GroupUser) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Group", groupUser)
	if err != nil {
		logger.Error("mongo insert GroupUser err: %v", err)
		return err
	}
	logger.Info("Mongo insert GroupUser success, id: %v", res.InsertedID)
	return nil
}
