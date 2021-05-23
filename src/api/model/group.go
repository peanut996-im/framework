// Package model
// @Title  group.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:23
package model

import (
	"framework/db"
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
		return err
	}
	// Second try to insert group
	_, err := mongo.InsertOne("Group", group)
	if err != nil {
		return err
	}
	return nil
}

func CreateGroup(name, admin string) error {
	g := NewGroup()
	g.GroupAdmin = admin
	g.GroupName = name
	if err := insertGroup(g); nil != err {
		return err
	}
	// create admin
	if err := CreateGroupUser(g.GroupID, admin); nil != err {
		return err
	}
	return nil
}
