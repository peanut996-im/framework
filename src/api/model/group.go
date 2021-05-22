// Package model
// @Title  group.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:23
package model

import (
	"framework/api"
	"framework/db"
	"framework/logger"
	"time"
)

type Group struct {
	RoomID     string `json:"room_id" bson:"roomID"`
	GroupID    string `json:"group_id" bson:"groupID"`
	GroupName  string `json:"group_name" bson:"groupName"`
	GroupAdmin string `json:"group_admin" bson:"groupAdmin"`
	CreateTime int64  `json:"create_time" bson:"createTime"`
}

func NewGroup() *Group {
	return &Group{
		GroupID:    api.NewSnowFlakeID(),
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
