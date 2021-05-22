// Package model
// @Title  groupuser.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:40
package model

import (
	"framework/db"
	"framework/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type GroupUser struct {
	GroupId string `json:"group_id" bson:"groupID"`
	UID     string `json:"user_id" bson:"UID"`
}

func NewGroupUser() *GroupUser {
	return &GroupUser{}
}

func insertGroupUser(groupUser *GroupUser) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Group", groupUser)
	if err != nil {
		logger.Error("mongo insert GroupUser err: %v", err)
		return err
	}
	logger.Info("Mongo insert GroupUser success, id: %v", res.InsertedID)
	return nil
}

func CreateGroupUser(group, user string) error {
	gp := NewGroupUser()
	gp.GroupId = group
	gp.UID = user
	if err := insertGroupUser(gp); nil != err {
		logger.Error("mongo insert GroupUser err: %v", err)
		return err
	}
	return nil
}

func DeleteGroupUser(group, user string) error {
	mongo := db.GetLastMongoClient()
	filter := bson.M{"groupID": group, "UID": user}
	if _, err := mongo.DeleteOne("GroupUser", filter); nil != err {
		return err
	}
	return nil
}
