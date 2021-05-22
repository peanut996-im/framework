// Package model
// @Title  group.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:23
package model

import (
	"framework/db"
	"framework/logger"
	"framework/tool"
	"time"
)

type Group struct {
	RoomID     string `json:"roomID" bson:"roomID"`
	GroupID    string `json:"groupID" bson:"groupID"`
	GroupName  string `json:"groupName" bson:"groupName"`
	GroupAdmin string `json:"groupAdmin" bson:"groupAdmin"`
	CreateTime int64  `json:"createTime" bson:"createTime"`
}

func NewGroup() *Group {
	return &Group{
		GroupID:    tool.NewSnowFlakeID(),
		CreateTime: time.Now().Unix(),
	}
}

func insertGroup(group *Group) error {
	mongo := db.GetLastMongoClient()
	g := NewGroup()
	r := NewGroupRoom()
	g.RoomID = r.RoomID
	g.GroupID = r.RoomID
	// First try to insert the room
	if err := insertRoom(r); nil != err {
		logger.Error("InsertGroupRoom err: %v", err)
		return err
	}
	// Second try to insert group
	res, err := mongo.InsertOne("Group", group)
	if err != nil {
		logger.Error("mongo insert Group err: %v", err)
		return err
	}
	logger.Info("Mongo insert Group success, id: %v", res.InsertedID)
	return nil
}

func CreateGroup(name, admin string) error {
	g := NewGroup()
	g.GroupAdmin = admin
	g.GroupName = name
	if err := insertGroup(g); nil != err {
		logger.Error("mongo insert Group err: %v", err)
		return err
	}
	return nil
}
