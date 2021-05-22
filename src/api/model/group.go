// Package model
// @Title  group.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:23
package model

import (
	"framework/db"
	"framework/logger"
)

type Group struct {
	RoomID     string `json:"room_id"`
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int64  `json:"create_time"`
}

func NewGroup() *Group {
	return &Group{}
}

func InsertGroup(group *Group) error {
	g := NewGroup()
	r := NewGroupRoom()
	g.RoomID = r.RoomID
	g.GroupID = r.RoomID
	// First try to insert the room
	if err := InsertRoom(r); nil != err {
		logger.Error("InsertGroupRoom err: %v", err)
		return err
	}
	// Second try to insert group
	if err := insertGroup(g); nil != err {
		logger.Error("InsertGroup err: %v", err)
		return err
	}
	return nil
}

func insertGroup(group *Group) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Group", group)
	if err != nil {
		logger.Error("mongo insert Group err: %v", err)
		return err
	}
	logger.Info("Mongo insert Group success, id: %v", res.InsertedID)
	return nil
}
