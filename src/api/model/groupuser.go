// Package model
// @Title  groupuser.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:40
package model

import (
	"framework/db"
	"framework/logger"
	"framework/tool"
	"go.mongodb.org/mongo-driver/bson"
)

type GroupUser struct {
	GroupID    string `json:"groupID" bson:"groupID"`
	UID        string `json:"uid" bson:"uid"`
	CreateTime string `json:"-" bson:"createTime"`
}

func NewGroupUser() *GroupUser {
	return &GroupUser{
		CreateTime: tool.GetNowUnixMilliSecond(),
	}
}

func insertGroupUser(groupUser *GroupUser) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne(MongoCollectionGroupUser, groupUser)
	if err != nil {
		logger.Error("mongo insert GroupUser err: %v", err)
		return err
	}
	logger.Info("Mongo insert GroupUser success, id: %v", res.InsertedID)
	return nil
}

func CreateGroupUser(group, user string) error {
	gp := NewGroupUser()
	gp.GroupID = group
	gp.UID = user
	if err := insertGroupUser(gp); nil != err {
		logger.Error("mongo insert GroupUser err: %v", err)
		return err
	}
	return nil
}

func DeleteGroupUser(group, user string) error {
	mongo := db.GetLastMongoClient()
	filter := bson.M{"groupID": group, "uid": user}
	if _, err := mongo.DeleteOne(MongoCollectionGroupUser, filter); nil != err {
		return err
	}
	return nil
}
